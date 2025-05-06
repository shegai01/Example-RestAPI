PORT="8080"
.phony: run
run:
	go run ./cmd/api/main.go
build:
	go build -v ./cmd/api
stop:
	@fuser -k ${PORT}/tcp