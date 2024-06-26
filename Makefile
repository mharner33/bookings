build:
	@go build -o bin/bookings cmd/web/*.go

run: build
	@./bin/bookings

test:
	@go test -v ./...

clean:
	@rm bin/*