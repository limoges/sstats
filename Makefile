NAME=sstats
OUTPUT=cmd/$(NAME)

build:
	go build -o $(OUTPUT) cmd/main.go

test:
	go test ./...

install: build
	cp $(OUTPUT) $(GOPATH)/bin/$(NAME)

clean:
	rm -fd $(OUTPUT)
