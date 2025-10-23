.PHONY: build run test clean docker-build docker-up docker-down migrate-up migrate-down

build:
	go build -o bin/main cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

migrate-up:
	migrate -path migrations -database "mysql://newsletter_user:newsletter_password@tcp(localhost:3306)/newsletter" up

migrate-down:
	migrate -path migrations -database "mysql://newsletter_user:newsletter_password@tcp(localhost:3306)/newsletter" down

.DEFAULT_GOAL := build
