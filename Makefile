GENERATED_TARGET_DIR ?= $(CURDIR)/generated
BUILD_TARGET_DIR ?= $(GENERATED_TARGET_DIR)/bin
SUDO_COMMAND := sudo

.PHONY: setup-linux-install
setup-linux-install:
	$(SUDO_COMMAND) apt-get update

.PHONY: setup-docker-go
setup-docker-go: ## Setup of Docker go build container
setup-docker-go: SUDO_COMMAND :=
setup-docker-go: setup-linux-install

.PHONY: generate
generate:
	# TODO: Add code generation here

build-go-%: generate
	CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $(BUILD_TARGET_DIR)/cmd/$(*F)/$(*F) $(CURDIR)/cmd/$(*F)/
