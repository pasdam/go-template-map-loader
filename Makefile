BUILD_DIR ?= .build
PROJECT_NAME ?= "app"

include scripts/makefiles/third_party/pasdam/makefiles/go.mk
include scripts/makefiles/third_party/pasdam/makefiles/go.mod.mk
include scripts/makefiles/third_party/pasdam/makefiles/help.mk

.DEFAULT_GOAL := help

## clean: Remove all artifacts
.PHONY: clean
clean: go-clean
