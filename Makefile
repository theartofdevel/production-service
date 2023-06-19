.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf ./app/build || true

.PHONY: up-local-env
up-local-env: down-local-env
	@docker-compose -f docker-compose.local.yml up -d

.PHONY: down-local-env
down-local-env:
	@docker-compose -f docker-compose.local.yml stop