APP_BIN = app/build/app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/app/main.go

.PHONY: clean
clean:
	rm -rf ./app/build || true