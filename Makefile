.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/*

.PHONY: run
run: 
	go run cmd/server/main.go

.PHONY: test
test: 
	go test ./...

.PHONY: benchmark
benchmark: 
	go test -bench=. -benchmem ./...

.PHONY: image
image: goapp
	docker build -t goapp -f build/package/Dockerfile .

.PHONY: clean-image
clean-image: goapp
	docker build --no-cache -t goapp -f build/package/Dockerfile .

.PHONY: run-container
 run-container: image
	docker run --rm -it -p 8080:8080 goapp