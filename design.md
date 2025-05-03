logger --> present at all level (this should be singleton)
server --> start server may be intilize other service as well like kafka in later section
main --> run server and initilaize the logger

--------Db
db instance only availabe here to make query


#======= create migration rule 
brew install golang-migrate
migrate create -ext sql -dir db/migrations -seq create_tasks_table

db/migrations/
├── 000001_create_tasks_table.up.sql
└── 000001_create_tasks_table.down.sql


MIGRATE=migrate
MIGRATIONS_DIR=./db/migrations
DB_URL=mysql://root:password@tcp(localhost:3306)/taskdb

.PHONY: migrate-up migrate-down migrate-force migrate-create

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


----- model task

taskID -> auto_increment, primary , int 
title string
descriptions strings
status (in_progress, on hold, done)
remarks string
createdat time.Time
updatedat time.Time