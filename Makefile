# Go(Golang) Options
GOCMD=go
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOTEST=$(GOCMD) test
GOFLAGS := -v 
LDFLAGS := -s -w

# Golangci Options
GOLANGCILINTCMD=golangci-lint
GOLANGCILINTRUN=$(GOLANGCILINTCMD) run

ifneq ($(shell go env GOOS),darwin)
LDFLAGS := -extldflags "-static"
endif

.PHONY: tidy
tidy:
	$(GOMOD) tidy

.PHONY: format
format:
	$(GOFMT) ./...

.PHONY: lint
lint:
	$(GOLANGCILINTRUN) ./...

.PHONY: lint-fix
lint-fix:
	$(GOLANGCILINTRUN) --fix ./...

.PHONY: test
test:
	$(GOTEST) $(GOFLAGS) ./...