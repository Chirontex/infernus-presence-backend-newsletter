.PHONY: build run test clean migrate-up migrate-down

build:
	go build -o bin/main cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

migrate-up:
	@set -a; . ../../compose/.env; set +a; \
	migrate -path migrations -database "mysql://$${NEWSLETTER_DB_USER}:$${NEWSLETTER_DB_PASSWORD}@tcp($${NEWSLETTER_DB_HOST}:$${NEWSLETTER_DB_PORT})/$${NEWSLETTER_DB_NAME}" up

migrate-down:
	@set -a; . ../../compose/.env; set +a; \
	migrate -path migrations -database "mysql://$${NEWSLETTER_DB_USER}:$${NEWSLETTER_DB_PASSWORD}@tcp($${NEWSLETTER_DB_HOST}:$${NEWSLETTER_DB_PORT})/$${NEWSLETTER_DB_NAME}" down

.DEFAULT_GOAL := build
