include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations
DB_URL = postgres://user:adminpassword@localhost:5432/social?sslmode=disable

.PHONY: migrate-create migrate-up migrate-down migrate-force

migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(NAME)

migrate-up:
	@migrate -source file://$(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	@migrate -source file://$(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate-force:
	@migrate -source file://$(MIGRATIONS_PATH) -database "$(DB_URL)" force $(VERSION) 