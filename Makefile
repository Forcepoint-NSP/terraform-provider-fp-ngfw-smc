PROVIDER_NAME=fp-ngfw-smc
PROVIDER_FULLNAME=terraform-provider-fp-ngfw-smc
PROVIDER_VERSION=0.0.1
PROVIDER_ORGANIZATION=Forcepoint
PLUGIN_DIR=plugins/registry.terraform.io/${PROVIDER_ORGANIZATION}/${PROVIDER_NAME}/${PROVIDER_VERSION}/linux_amd64
RUN=./scripts/run_go

all: build

# Show available targets and file management commands
.PHONY: help
help:
	@echo "SMC Terraform Provider Build Commands:"
	@echo "  make build              - Build the binary"
	@echo "  make release            - Prepare the provider release (build + goreleaser)"
	@echo "  make release-snapshot- Build snapshot without publishing (no validation, dirty edit allowed, tag must be set on HEAD)"
	@echo "  make clean              - Clean all build artifacts"
	@echo "  make help               - Show this help message"


.PHONY: build
build:
	mkdir $(PLUGIN_DIR) 2>/dev/null; true
	go build  -o $(PLUGIN_DIR)/${PROVIDER_FULLNAME} .

.PHONY: release-snapshot
go-release-snapshot: docker-build
	goreleaser build --snapshot --clean --skip=validate

.PHONY: release
release: build
	./scripts/init_pgp_and_make_release.sh

.PHONY: clean
clean:
	chmod -R u+w .cache/gomod || true
	rm -fr .cache $(PROVIDER_FULLNAME) dist/ || true
