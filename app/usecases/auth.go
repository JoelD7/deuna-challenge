package usecases

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/JoelD7/deuna-challenge/app/models"
	"github.com/gbrlsnchs/jwt/v3"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

const (
	passwordCost         = bcrypt.DefaultCost
	accessTokenDuration  = 300
	refreshTokenDuration = 86400
)

var (
	accessTokenIssuer = os.Getenv("APP_HOST")
	privateKey        = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAogelHdI7G4dlunQ0WvDdLU5ZjHoptN9i8BYfqe9VieDUxOUz
jXzD7OvLc3ARgqel5nu8LB0jcYPSeR+yTbL9BN2kvBOGxPf1mlALdPRDLlTnKe6Y
o+K7+Xv9tkEbMn7zl3//emPS4m3kRc47XHtJSpuHgR4tEyxxZ/IA+XdcBcnbwmz9
6tBofE7iQWc3DCedf79+L3JPeJnzkH0Lg5dDU8e6/fRBwKSfAU//JGvjHKXaUoGx
JPBU0PnjxsOuIoN7b6ZkK+llba6K8wrn/LhqKog1/S097v9tlPM7gr8m8XbfgNzH
UXCBcLhep3534d+UEnn4ouslaZ9XGP6e5yyX7wIDAQABAoIBAADgr7lIoT9V7Wwk
IwB3G4uaSAvlwYIUT7HjPMqr3DfB+wUSBMR4b4tB/7khW0bs544nD27hvYZo42P6
kvmuxYYYOUM3i9xXR4JNerJofFCs7w+gFj2VBdWlIUuycJZGb8VbUSP1lHfbhogG
RPYMSOpZi1NcuXvIGtkoS28OgXYajpvzEG/obyUUkadUngIXa/TDghKandaU6js+
mlbc5d7SBnoeBiAh6AND6H1iQHAqTjSPvU75fCidB8Ab3ezurzgK2GxC5gaFE96T
q6BM4MPPRkvjg6xD1tEhQXG8IuHU2BLkpZZr+oTCMVKM9SiFFzhjxBEWG9NRgCLQ
QUnE/nkCgYEA01HVkx4B0qq0/47jff0a6QJ1/Bl9QfrEtT/ATFQSBFR52OjxRgSV
s1+hW/NtG9tFPYgrt7lay0G557gojn2wxC9yk71JvbgNL38HZsObgcbjllRSS/Zn
pWlFS2v/XunkcFqfcPZUK8iXU+J8TXdJ48modVdPibHqfzyilYyDMHUCgYEAxEni
21DdsNm9aZ7ZSu+7LvJ+/VFQUSJtiTKRx4DWNGIsfy9VBEZ93ocOuVX1iDd3a9mE
8PkS6tc59AY6d8Qc4f6IklaYPK5YPiLbskfknOVq22U0rlgiLd8wDDg5NPFqHpKc
LexeuKQCskGPt2mvHYydQzi7ZsRznTkaVzfYGlMCgYAf1bc4F5AsvXzQ9yS8aTHx
omZF2U0ucGnL6FO+6/de4Z8Nl2Ipqy0mPaTgZlasmKbgsy/q2Kid8EPibbLmbHcB
xygaq6x9QUnzOs7Ro2w868qDbiaLvQ42NBq1Vwq8sL2yU2SrruBVTD3H7FnPjcX6
4/lV3BZmZwAttOPFZcqptQKBgA7FkLD4kPZyLHL6ZVfiWq/Zx/zAVc8FTED68UWW
SIiAquCXa0p7E5XfjBgeg+/QXMhdAkgwNmA9+jqHDXdd5t6LDTQWGDbY2AM1FFuC
VY4JJdWE9EX6k/fnx/HjeUqmsFnEpsQ9+ZLjpOBNVsdyyJ7sqhkY9+Fv/1NhrL3L
khPJAoGAJeUe8/JBnSG6+pcKMCS/7BKrDwjnx32Av9JgKMR662M8VzYuQm5dn/r7
zrArJVptyTX9EYwOkOA3e6aupi06iDJ7ym2y6+1eOokcHkySqfOSxvLEhCi3MnHL
6J9fSRUtc/JyYMClgSaMOABaQvlv02bFpe0qnVCKUujUrDnv5xo=
-----END RSA PRIVATE KEY-----`)
	publicKey = os.Getenv("RSA_PUBLIC_KEY")
)

type UserManager interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, email string) (*models.User, error)
}

func NewUserCreator(userManager UserManager) func(ctx context.Context, user *models.User) error {
	return func(ctx context.Context, user *models.User) error {
		unhashedPassword := *user.Password

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(unhashedPassword), passwordCost)
		if err != nil {
			return err
		}

		password := string(hashedPassword)

		user.Password = &password
		now := time.Now()
		user.CreatedDate = &now

		err = userManager.CreateUser(ctx, *user)
		if err != nil {
			return err
		}

		return nil
	}
}

func NewUserAuthenticator(userManager UserManager) func(ctx context.Context, email, password string) (*models.User, error) {
	return func(ctx context.Context, email, password string) (*models.User, error) {
		user, err := userManager.GetUser(ctx, email)
		if err != nil {
			return nil, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
		if err != nil {
			return nil, models.ErrWrongCredentials
		}

		return user, nil
	}
}

func NewUserTokenGenerator() func(ctx context.Context, user *models.User) (*models.AuthToken, *models.AuthToken, error) {
	return func(ctx context.Context, user *models.User) (*models.AuthToken, *models.AuthToken, error) {
		now := time.Now()

		accessTokenExpiry := jwt.NumericDate(now.Add(time.Duration(accessTokenDuration) * time.Second))

		accessTokenPayload := &jwt.Payload{
			Issuer:         accessTokenIssuer,
			Subject:        *user.Email,
			ExpirationTime: accessTokenExpiry,
			IssuedAt:       jwt.NumericDate(now),
		}

		accessToken, err := generateJWT(accessTokenPayload)
		if err != nil {
			return nil, nil, err
		}

		refreshTokenExpiry := jwt.NumericDate(now.Add(time.Duration(refreshTokenDuration) * time.Second))

		refreshTokenPayload := &jwt.Payload{
			Subject:        *user.Email,
			ExpirationTime: refreshTokenExpiry,
		}

		refreshToken, err := generateJWT(refreshTokenPayload)
		if err != nil {
			return nil, nil, err
		}

		access := &models.AuthToken{
			Value:      accessToken,
			Expiration: accessTokenExpiry.Time,
		}

		refresh := &models.AuthToken{
			Value:      refreshToken,
			Expiration: refreshTokenExpiry.Time,
		}

		return access, refresh, nil
	}
}

func generateJWT(payload *jwt.Payload) (string, error) {
	priv, err := getPrivateKey()
	if err != nil {
		return "", fmt.Errorf("private key fetching failed: %w", err)
	}

	var signingHash = jwt.NewRS256(jwt.RSAPrivateKey(priv))

	token, err := jwt.Sign(payload, signingHash)
	if err != nil {
		return "", fmt.Errorf("jwt signing failed: %w", err)
	}

	return string(token), nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	privatePemBlock, _ := pem.Decode(privateKey)
	if privatePemBlock == nil || !strings.Contains(privatePemBlock.Type, "PRIVATE KEY") {
		return nil, fmt.Errorf("failed to decode PEM private block containing private key")
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(privatePemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return rsaPrivateKey, nil
}
