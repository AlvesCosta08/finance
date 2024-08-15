DB_CONTAINER := finance
DB_USER := postgres
DB_NAME := go_finance

migrate-up:
	docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -f /migration/000001_initial_tables.up.sql

migrate-down:
	docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -f /migration/000001_initial_tables.down.sql
