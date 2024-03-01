export BANK_HOST := localhost:8081
export BANK_DB := deuna-bank-db.sqlt
export APP_HOST := localhost:8080
export APP_DB := deuna-db.sqlt

start-bank:
	  go run bank/main.go
start-app:
	   go run app/main.go
setup-db:
	sqlite3 deuna-db.sqlt < app/db/migrations/delete_database_tables.sql && sqlite3 deuna-bank-db.sqlt < bank/db/migrations/delete_database_tables.sql
	sqlite3 deuna-db.sqlt < app/db/migrations/create_database_tables.sql && sqlite3 deuna-bank-db.sqlt < bank/db/migrations/create_database_tables.sql
install:
	go get ./...

prepare: install setup-db