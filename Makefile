# Print warnings if the .env does not exist.
ifeq (,$(wildcard .env))
$(info WARN: .env file not found!)
$(info WARN: some environment variables may need to be set!)
endif

-include .env


###############################################################################
# CODE GENERATION
###############################################################################

# Generate code from .sql files for data layers.
sqlc:
	sqlc generate
.PHONY: sqlc

###############################################################################
# MIGRATION
###############################################################################

# Create new migration files.
# Example: make add-migration name=create_table_abc
add-migration:
	migrate create -ext sql -seq -dir $(POSTGRES_MIGRATION) $(name)
.PHONY: add-migration


# Run database migration up.
up-migration:
	migrate -path $(POSTGRES_MIGRATION) -database $(POSTGRES_URI) up
.PHONY: up-migration


# Run database migration down.
down-migration:
	migrate -path $(POSTGRES_MIGRATION) -database $(POSTGRES_URI) down -all
.PHONY: down-migration


###############################################################################
# DOCKER
###############################################################################

compose-up:
	docker compose up $(args)
.PHONY: compose-up

compose-down:
	docker compose down
.PHONY: compose-down

# Rebuild application image for docker compose.
compose-build:
	docker build . -t appointment
.PHONY: compose-build


###############################################################################
# OTHERS
###############################################################################

# Start a Postgres database in Docker.
start-db:
	docker run --rm -d --name appointment-db -p 5432:5432 \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
 		postgres:16-alpine
.PHONY: start-db


# Stop the Postgres database in Docker.
stop-db:
	docker stop appointment-db
.PHONY: stop-db
