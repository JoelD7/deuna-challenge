start-bank:
	 BANK_HOST=localhost:8081 BANK_DB=deuna-bank-db.sqlt APP_HOST=localhost:8080 APP_DB=deuna-db.sqlt go run bank/main.go
start-app:
	  BANK_HOST=localhost:8081 BANK_DB=deuna-bank-db.sqlt APP_HOST=localhost:8080 APP_DB=deuna-db.sqlt go run app/main.go