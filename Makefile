url:
	go run utils/register.go

clean:
	rm -rf ./build/

lint:
	golangci-lint run