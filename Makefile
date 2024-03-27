LOCAL_BIN:=$(CURDIR)/bin

install-swag:
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@v1.16.3

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose@v2.7.0