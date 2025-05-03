APP_NAME := task-manager
CMD_DIR := cmd/${APP_NAME}
MIGRATE=migrate
MIGRATIONS_DIR=./db/migrations
DB_URL=mysql://root:password@tcp(localhost:3306)/taskdb
.PHONY: migrate-up migrate-down migrate-force migrate-create
# default target
all: build
run:
	go run ${CMD_DIR}/main.go


# DB
migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migrate-drop:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" drop -f

migrate-force:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force

migrate-version:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name