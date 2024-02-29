start:
	BANK_URL=localhost:8081 BANK_DB=deuna-bank-db.sqlt go run bank/main.go
	APP_URL=localhost:8080 APP_DB=deuna-db.sqlt go run app/main.go