hello:
	echo "Hello"

build:
	go build -o tmp/main cmd/main.go

run:
	go run cmd/main.go

dev:
	go build -o tmp/main cmd/main.go && /tmp/main
