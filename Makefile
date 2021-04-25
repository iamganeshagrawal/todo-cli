.PHONY: client

client:
	@echo "Building the client binary"
	go build -o bin/todo.exe main.go