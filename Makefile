AirVersion := $(shell air -v 2>/dev/null)

run:
ifdef AirVersion
	@air server
else
	@echo "please run the command 'go install github.com/cosmtrek/air@latest' to install Air live reload"
endif

migration-status:
	@go run cmd/db-migration.go status

migration-up:
	@go run cmd/db-migration.go up

migration-down:
	@go run cmd/db-migration.go down