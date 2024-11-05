.PHONY: build

build:
	go build -o task-cli ./cmd
	chmod 755 ./task-cli