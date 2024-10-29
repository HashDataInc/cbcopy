.DEFAULT_GOAL := all
.PHONY: integration

CBCOPY=cbcopy
COPYHELPER=cbcopy_helper

GINKGO_FLAGS := -r --keep-going --no-color
GIT_VERSION := $(shell git describe --tags | perl -pe 's/(.*)-([0-9]*)-(g[0-9a-f]*)/\1+dev.\2.\3/')
VERSION_STR="-X github.com/cloudberrydb/cbcopy/utils.Version=$(GIT_VERSION)"

SUBDIRS_HAS_UNIT= meta/builtin/ testutils/ utils/
GINKGO=$(shell go env GOPATH)/bin/ginkgo
GOIMPORTS=$(shell go env GOPATH)/bin/goimports

all: build

depend:
	$(GO_ENV) go mod download

$(GINKGO):
	$(GO_ENV) go install github.com/onsi/ginkgo/v2/ginkgo@v2.8.4

$(GOIMPORTS):
	$(GO_ENV) go install golang.org/x/tools/cmd/goimports

format: $(GOIMPORTS)
	@goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

unit: $(GINKGO)
	@echo "Running unit tests..."
	$(GINKGO) $(GINKGO_FLAGS) $(SUBDIRS_HAS_UNIT) 2>&1

integration: $(GINKGO)
	@echo "Running integration tests..."
	$(GINKGO) $(GINKGO_FLAGS) integration 2>&1

test: unit

end_to_end: $(GINKGO)
	@echo "Running end to end tests..."
	$(GINKGO) $(GINKGO_FLAGS) end_to_end 2>&1

build:
	$(GO_ENV) go build -tags '$(CBCOPY)' $(GOFLAGS) -o $(CBCOPY) -ldflags $(VERSION_STR)
	$(GO_ENV) go build -tags '$(COPYHELPER)' $(GOFLAGS) -o $(COPYHELPER) -ldflags $(VERSION_STR)

install:
	cp $(CBCOPY) $(GPHOME)/bin
	cp $(COPYHELPER) $(GPHOME)/bin

clean:
	rm -f $(CBCOPY)
	rm -f $(COPYHELPER)
