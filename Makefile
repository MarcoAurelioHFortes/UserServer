build:
	@go build -o bin/APISTUDY cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/APISTUDY