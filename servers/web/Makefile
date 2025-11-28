

ENV_FILE ?= ../../envs/migrator.env
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

export PGPASSWORD=$(DB_PASS)
DB_DSN = postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&search_path=${DB_SCHEMA}
sqlc:
	## clean ./db
	rm -rf ./db
	mkdir -p ./db

	pg_dump -U $(DB_USER) \
		-h $(DB_HOST) \
		-p $(DB_PORT) \
		-s -d $(DB_NAME) \
		--schema=$(DB_SCHEMA) \
		--no-owner \
		--no-privileges \
		--no-comments \
		--no-security-labels \
		--no-tablespaces \
		> schema.sql	
	sqlc generate

