# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.9. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for cobra-cli variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(COBRA_CLI)
#	@echo "Running cobra-cli"
#	@$(COBRA_CLI) <flags/args..>
#
COBRA_CLI := $(GOBIN)/cobra-cli-v1.3.0
$(COBRA_CLI): $(BINGO_DIR)/cobra-cli.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/cobra-cli-v1.3.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=cobra-cli.mod -o=$(GOBIN)/cobra-cli-v1.3.0 "github.com/spf13/cobra-cli"

OAPI_CODEGEN := $(GOBIN)/oapi-codegen-v2.1.0
$(OAPI_CODEGEN): $(BINGO_DIR)/oapi-codegen.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/oapi-codegen-v2.1.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=oapi-codegen.mod -o=$(GOBIN)/oapi-codegen-v2.1.0 "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"

TASK := $(GOBIN)/task-v3.35.1
$(TASK): $(BINGO_DIR)/task.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/task-v3.35.1"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=task.mod -o=$(GOBIN)/task-v3.35.1 "github.com/go-task/task/v3/cmd/task"

YQ := $(GOBIN)/yq-v4.43.1
$(YQ): $(BINGO_DIR)/yq.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/yq-v4.43.1"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=yq.mod -o=$(GOBIN)/yq-v4.43.1 "github.com/mikefarah/yq/v4"

