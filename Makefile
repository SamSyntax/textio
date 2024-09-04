.PHONY: hello build run dev down up

hello:
	echo "Hello"

build:
	go build -o tmp/main cmd/main.go

run:
	go run cmd/main.go

dev:
	go build -o tmp/main cmd/main.go && /tmp/main

down:
	@bash -c "goose -dir sql/schema postgres 'postgres://postgres:1234@localhost:5433/textio?sslmode=disable' down"

up:
	@bash -c "goose -dir sql/schema postgres 'postgres://postgres:1234@localhost:5433/textio?sslmode=disable' up"
