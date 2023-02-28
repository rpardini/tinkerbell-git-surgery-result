OSFLAG := $(shell go env GOHOSTOS)
BINARY := ipxe
IPXE_BUILD_SCRIPT := binary/script/build_ipxe.sh
IPXE_FETCH_SCRIPT := binary/script/fetch_and_extract_ipxe.sh
IPXE_NIX_SHELL := binary/script/shell.nix

help: ## show this help message
	@grep -E '^[a-zA-Z_-]+.*:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

include lint.mk

.PHONY: binary
binary: binary/ipxe.efi binary/snp.efi binary/undionly.kpxe binary/ipxe.iso ## build all upstream ipxe binaries

# ipxe_sha_or_tag := v1.21.1 # could not get this tag to build ipxe.efi
# https://github.com/ipxe/ipxe/tree/2265a65191d76ce367913a61c97752ab88ab1a59
ipxe_sha_or_tag := $(shell cat binary/script/ipxe.commit)
ipxe_readme := upstream-$(ipxe_sha_or_tag)/README

# building iPXE on a Mac is troublesome and difficult to get working. For that reason, on a Mac, we build the iPXE binary using Docker.
ipxe_build_in_docker := $(shell if [ $(OSFLAG) = "darwin" ]; then echo true; else echo false; fi)

.PHONY: extract-ipxe
extract-ipxe: $(ipxe_readme) ## Fetch and extract ipxe source
$(ipxe_readme): binary/script/ipxe.commit
	${IPXE_FETCH_SCRIPT} "$(ipxe_sha_or_tag)"
	touch "$@"

binary/ipxe.efi: $(ipxe_readme) ## build ipxe.efi
	+${IPXE_BUILD_SCRIPT} bin-x86_64-efi/ipxe.efi "$(ipxe_sha_or_tag)" $(ipxe_build_in_docker) $@ "${IPXE_NIX_SHELL}"

binary/undionly.kpxe: $(ipxe_readme) ## build undionly.kpxe
	+${IPXE_BUILD_SCRIPT} bin/undionly.kpxe "$(ipxe_sha_or_tag)" $(ipxe_build_in_docker) $@ "${IPXE_NIX_SHELL}"

binary/snp.efi: $(ipxe_readme) ## build snp.efi
	+${IPXE_BUILD_SCRIPT} bin-arm64-efi/snp.efi "$(ipxe_sha_or_tag)" $(ipxe_build_in_docker) $@  "${IPXE_NIX_SHELL}" "CROSS_COMPILE=aarch64-unknown-linux-gnu-"

binary/ipxe.iso: $(ipxe_readme) ## build ipxe.iso
	+${IPXE_BUILD_SCRIPT} bin-x86_64-efi/ipxe.iso "$(ipxe_sha_or_tag)" $(ipxe_build_in_docker) $@  "${IPXE_NIX_SHELL}"

.PHONY: binary/clean
binary/clean: ## clean ipxe binaries, upstream ipxe source code directory, and ipxe source tarball
	rm -rf binary/ipxe.efi binary/snp.efi binary/undionly.kpxe binary/ipxe.iso
	rm -rf upstream-*
	rm -rf ipxe-*

.PHONY: test
test: ## run unit tests
	go test -v -covermode=count ./...

.PHONY: cover
cover: ## Run unit tests with coverage report
	go test -coverprofile=coverage.out ./... || true
	go tool cover -func=coverage.out

.PHONY: build-linux
build-linux: ## Compile for linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -extldflags "-static"' -o bin/${BINARY}-linux cmd/main.go

.PHONY: build-darwin
build-darwin: ## Compile for darwin
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -extldflags '-static'" -o bin/${BINARY}-darwin cmd/main.go

.PHONY: build
build: ## Compile the binary for the native OS
ifeq (${OSFLAG},linux)
	@$(MAKE) build-linux
else
	@$(MAKE) build-darwin
endif
