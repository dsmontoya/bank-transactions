GENERATED_TARGET_DIR ?= $(CURDIR)/generated
BUILD_TARGET_DIR ?= $(GENERATED_TARGET_DIR)/bin
SUDO_COMMAND := sudo

DOCKER_COMPOSE_DEV_DB_PROJECT := transactions-dev-db

INSTALL_BIN_DIR := $(CURDIR)/bin
export PATH := $(INSTALL_BIN_DIR):$(PATH)
export GOBIN := $(INSTALL_BIN_DIR)

GOOSE_VERSION := v3.22.1
SQLC_VERSION := v1.27.0

.PHONY: setup-linux-install
setup-linux-install:
	$(SUDO_COMMAND) apt-get update

.PHONY: setup-common-gen
setup-common-gen:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION)
	go install -tags='no_mysql no_sqlite3' github.com/pressly/goose/v3/cmd/goose@$(GOOSE_VERSION)

.PHONY: setup-docker-go
setup-docker-go: ## Setup of Docker go build container
setup-docker-go: SUDO_COMMAND :=
setup-docker-go: setup-linux-install setup-common-gen

.PHONY: generate
generate: gen-sql

build-go-%: generate
	CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $(BUILD_TARGET_DIR)/cmd/$(*F)/$(*F) $(CURDIR)/cmd/$(*F)/

.PHONY: gen-sql
gen-sql:
	find sql -mindepth 2 -maxdepth 2 -name sqlc.yaml | xargs -n 1 sqlc generate -f


db-create-migration-%:
	mkdir -p "$(CURDIR)/sql/$(*F)/migrations"
	"$(INSTALL_BIN_DIR)/goose" -dir "$(CURDIR)/sql/$(*F)/migrations" -table schema_migrations postgres "UNNEEDED_DATABASE_URL" create "$(name)" sql

run-dev-db-%:
	docker-compose -p $(DOCKER_COMPOSE_DEV_DB_PROJECT) -f $(CURDIR)/deployments/docker-compose.yaml up $(*F)db --detach

run:
	docker-compose -p $(DOCKER_COMPOSE_DEV_DB_PROJECT) -f $(CURDIR)/deployments/docker-compose.yaml up
