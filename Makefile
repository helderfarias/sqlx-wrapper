PKG_LIST_ALL_TESTS := $(shell go list ./... | grep -v /vendor)

test:
	@echo 'Unit Tests'
	@go test $(PKG_LIST_ALL_TESTS)
