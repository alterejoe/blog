docker-up:
	docker compose up -d

docker-down:
	docker compose down


plan:
	cd ./terraform/platform/local-infra/ && make plan
	
apply:
	cd ./terraform/platform/local-infra/ && make apply

init:
	cd ./terraform/platform/local-infra/ && make init

MIG_ENV_FILE ?= ./envs/migrator.env
include $(MIG_ENV_FILE)
export $(shell sed 's/=.*//' $(MIG_ENV_FILE))

DATABASE_URL="postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
DATABASE_URL_SEARCH="$(DATABASE_URL)&search_path=$(DB_SCHEMA)"
MIGRATIONS_PATH=./migrations

up:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DATABASE_URL_SEARCH) up
down:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DATABASE_URL_SEARCH) down 1
clean:
	psql $(DATABASE_URL) -c "UPDATE "$(DB_SCHEMA)".schema_migrations SET dirty = false;"
