AirVersion := $(shell air -v 2>/dev/null)

run:
ifdef AirVersion
	@air server
else
	@echo "please run the command 'go install github.com/cosmtrek/air@latest' to install Air live reload"
endif

migrate:
	@go run cmd/db_migration/main.go up

test-coverage:
	@go test -v -coverprofile cover.out ./...
	@go tool cover -html=cover.out -o cover.html
	@open cover.html

generate-mock:
	@echo "Generating..."
	@mockgen -destination=package/settings/settings_test.go -package=settings -source=package/settings/settings.go
	@echo "Done!"

build:
	@echo "build ghcr.io/ruriazz/gopen-api:latest"
	@container/backend/build.sh
	@echo "ghcr.io/ruriazz/gopen-api:latest build complete\n\n"

	@echo "build ghcr.io/ruriazz/gopen-api-commands:latest"
	@container/commands/build.sh
	@echo "ghcr.io/ruriazz/gopen-api-commands:latest build complete"