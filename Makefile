APP_NAME=galileo
PROJECT=github.com/panace9i/${APP_NAME}
HAS_DEP=$(shell command -v dep;)
GO_LIST_FILES=$(shell go list ${PROJECT}/... | grep -v vendor)

all: help
help:
	@echo
	@echo "------------------HELP-------------------------"
	@echo "build"
	@echo "Desc: build application"
	@echo
	@echo "test"
	@echo "Desc: run tests"
	@echo "-----------------------------------------------"
	@echo

.PHONY: build
build:  vendor dumps test
		go install ${PROJECT}

.PHONY: dumps
dumps:
ifeq ($(wildcard /tmp/${APP_NAME}_dumps),)
	mkdir /tmp/${APP_NAME}_dumps
endif

.PHONY: vendor
vendor: bootstrap
		dep ensure

.PHONY: bootstrap
bootstrap:
ifndef HAS_DEP
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: test
test:
	  @go test -race ${GO_LIST_FILES}
