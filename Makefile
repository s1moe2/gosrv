OS_NAME := $(shell uname -s | tr A-Z a-z)

migrate-up:
	./migrations/migrate.$(OS_NAME)-amd64 \
		-source file://migrations \
		-database "$(DB_URI)" up

migrate-down:
	./migrations/migrate.$(OS_NAME)-amd64 \
		-source file://migrations \
		-database "$(DB_URI)" down

migrate-create:
	./migrations/migrate.$(OS_NAME)-amd64 create \
		-dir ./migrations \
		-ext "sql" \
		-seq \
		"$(MIGRATION_NAME)"

test:
	go test ./... -cover