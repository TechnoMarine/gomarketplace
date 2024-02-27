NAME=go-gin-boilerplate
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test -v -cover ./...

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


dbup:
	@docker-compose up db -d

dbstop:
	@docker-compose stop db

migrateup:
	@migrate -path ./db/migration/ -database "postgresql://postgres:secret@localhost:12001/gomarketplacedb?sslmode=disable" -verbose up

migratedown:
	@migrate -path ./db/migration/ -database "postgresql://postgres:secret@localhost:12001/gomarketplacedb?sslmode=disable" -verbose down

sqlc:
	@sqlc generate
