OUTPUT=cmd/sstats

build:
	go build -o $(OUTPUT) cmd/main.go

test:
	go test ./...

clean:
	rm -fd $(OUTPUT)
