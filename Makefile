# Define local binary directory
LOCALBIN = $(shell pwd)/.bin

# Define the linter binary path
GOLANGCI_LINT = $(LOCALBIN)/golangci-lint

.PHONY: all lint lint-fix clean

all: lint

# Run the linter
lint: $(GOLANGCI_LINT)
	@echo "==> Running golangci-lint..."
	$(GOLANGCI_LINT) run ./...

# Run the linter and automatically fix safe issues (like formatting)
lint-fix: $(GOLANGCI_LINT)
	@echo "==> Running golangci-lint and fixing issues..."
	$(GOLANGCI_LINT) run --fix ./...

# Rule to build golangci-lint locally if it doesn't exist
$(GOLANGCI_LINT):
	@echo "==> Installing golangci-lint locally to $(LOCALBIN)..."
	@mkdir -p $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint

# Clean up local binaries
clean:
	rm -rf $(LOCALBIN)