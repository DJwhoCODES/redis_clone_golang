run:
	go run ./cmd/server

build:
	go build -o bin/redis-server ./cmd/server

test:
	go test ./...

clean:
	rm -rf bin
