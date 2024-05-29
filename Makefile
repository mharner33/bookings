build:
	@go build -o bin/ cmd/web/*.go

run: build
	@go run cmd/web/*.go

test:
	@go test -v ./...

clean:
	@rm bin/*