include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	@psql ${TODOS_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}


## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	@migrate -path=./migrations -database=${TODOS_DB_DSN} up

## db/migrations/down version=$1: apply down migrations
.PHONY: db/migrations/down
db/migrations/down:
	@if [ -n '${version}' ]; then \
		echo 'Running down to migration ${version}...'; \
	else \
		echo 'Running down migrations...'; \
	fi
	@migrate -path=./migrations -database=${TODOS_DB_DSN} down ${version}

## db/migrations/force version=$1: force a migration version
.PHONY: db/migrations/force
db/migrations/force:
	@migrate -path=./migrations -database=${TODOS_DB_DSN} force ${version}


## run/api: run the api
.PHONY: run/api
run/api:
	go run ./cmd/api

## audit: tidy, verify, format, vet and test code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	@go mod tidy
	@go mod verify
	@echo 'Formatting code...'
	@go fmt ./...
	@echo 'Vetting code...'
	@go vet ./...
	@echo 'Running tests...'
	@go test -race -vet=off ./...
