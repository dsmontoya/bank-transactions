GENERATED_TARGET_DIR ?= $(CURDIR)/generated
BUILD_TARGET_DIR ?= $(GENERATED_TARGET_DIR)/bin

.PHONY: generate
generate:
	# TODO: Add code generation here

build-go-%: generate
	CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $(BUILD_TARGET_DIR)/cmd/$(*F)/$(*F) $(CURDIR)/cmd/$(*F)/
