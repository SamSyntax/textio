hello:
	echo "Hello"

build:
	go build -o tmp/main cmd/main.go

run:
	tmp/main

dev:
	go build -o tmp/main cmd/main.go && /tmp/main
