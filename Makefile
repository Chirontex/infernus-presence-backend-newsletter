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
	migrate -path migrations -database "mysql://newsletter_user:newsletter_password@tcp(localhost:3306)/newsletter" up

migrate-down:
	migrate -path migrations -database "mysql://newsletter_user:newsletter_password@tcp(localhost:3306)/newsletter" down

.DEFAULT_GOAL := build
