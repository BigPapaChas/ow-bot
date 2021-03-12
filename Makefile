.PHONY: clean lint runserver build

clean:
	rm -rf ./build/
	mkdir -p build

lint:
	golangci-lint run

build:
	go build -o ./build/runserver ./cmd/bot


BOT_TOKEN=<<unset>>
PERMISSIONS=<<unset>>
runserver:
	./build/runserver -t ${BOT_TOKEN} -p ${PERMISSIONS}