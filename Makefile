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

.PHONY: swagger
swagger:
	swag init -g ./app/cmd/app/main.go -o ./app/docs

.PHONY: migrate
migrate:
	$(APP_BIN) migrate -version $(version)

.PHONY: migrate.down
migrate.down:
	$(APP_BIN) migrate -seq down

.PHONY: migrate.up
migrate.up:
	$(APP_BIN) migrate -seq up

.PHONY: proto
proto: clean format gen lint


# ===================== BUF =====================

BUF_VERSION=v1.6.0

.PHONY: buf-install
buf-install:
ifeq ($(shell uname -s), Darwin)
	@[ ! -f $(GOPATH)/bin/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Darwin-arm64 && \
	mv buf $(GOPATH)/bin/buf && \
	chmod +x $(GOPATH)/bin/buf
else
	@[ ! -f $(GOPATH)/bin/buf ] || exit 0 && \
	tmp=$$(mktemp -d) && cd $$tmp && \
	curl -L -o buf \
		https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-Linux-x86_64 && \
	mv buf $(GOPATH)/bin/buf && \
	chmod +x $(GOPATH)/bin/buf
endif

.PHONY: gen
gen: buf-install
	@$(GOPATH)/bin/buf generate
	@for dir in $(CURDIR)/gen/go/*/; do \
	  cd $$dir && \
	  go mod init && go mod tidy; \
  	done

.PHONY: lint
lint: buf-install
	@$(GOPATH)/bin/buf lint

.PHONY: format
format: buf-install
	@$(GOPATH)/bin/buf format


.PHONY: clean
clean:
	@rm -rf ./gen || true