build:
	go build -o cmd/sstats cmd/main.go

test:
	go test ./...
