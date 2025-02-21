.phony: build run
build :
	go run  ./api/main.go
run: build
	./main