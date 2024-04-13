build:
	@go build -o bin/verse-cli cmd/cli/cli.go
	@go build -o bin/verse-server cmd/server/server.go

run-cli: build
	bin/verse-cli

run-server: build
	bin/verse-server

clean:
	rm -rf bin
