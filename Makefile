export BANK_HOST := localhost:8081
export BANK_DB := deuna-bank-db.sqlt
export APP_HOST := localhost:8080
export APP_DB := deuna-db.sqlt

start-bank:
	  go run bank/main.go
start-app:
	   go run app/main.go