# Go(Golang) Options
GOCMD=go
GOMOD=$(GOCMD) mod
GOGET=$(GOCMD) get
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

.PHONY: update-deps
update-deps:
	$(GOGET) -f -t -u ./...
	$(GOGET) -f -u ./...

.PHONY: format
format:
	$(GOFMT) ./...

.PHONY: lint
lint:
	$(GOLANGCILINTRUN) ./... --fix

.PHONY: test
test:
	$(GOTEST) $(GOFLAGS) ./...