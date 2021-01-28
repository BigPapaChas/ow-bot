# .PHONY: 

clean:
	rm -rf ./build/

lint:
	golangci-lint run

build:./build/runserver
	go build -o ./build/runserver ./cmd/bot

# runserver: build
