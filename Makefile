# Makefile

# Comando para criar o banco de dados
createdb:
	createdb --username=postgres --owner=postgres go_finance

# Comando para iniciar o container PostgreSQL
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

# Comando para executar migrações para cima
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

# Comando para reverter migrações
migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose down

# Comando para rodar testes
test:
	go test -v -cover ./...

# Comando para iniciar o servidor
server:
	go run main.go

# Comando para gerar código com sqlc
sqlc-gen:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

# Define os targets que não são arquivos
.PHONY: createdb postgres migrateup migrationdrop test server sqlc-gen
