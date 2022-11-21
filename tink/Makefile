help: ## Print this help
	@grep --no-filename -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed 's/:.*##/·/' | sort | column -ts '·' -c 120

all: cli server worker ## Build all binaries for host OS and CPU

-include rules.mk
-include lint.mk
-include kube.mk

crosscompile: $(crossbinaries) ## Build all binaries for Linux and all supported CPU arches
images: tink-cli-image tink-server-image tink-worker-image virtual-worker-image ## Build all docker images
run: crosscompile run-stack ## Builds and runs the Tink stack (tink, db, cli) via docker-compose

test: e2etest-setup ## Run tests
	source <(setup-envtest use  -p env) && go test -coverprofile=coverage.txt ./...

verify: lint check-generated ## Verify code style, is lint free, freshness ...
	$(GOFUMPT) -s -d .

generated: pbfiles protomocks generate-manifests ## Generate dynamically created files
check-generated: check-pbfiles check-protomocks ## Check if generated files are up to date

tools: ${toolsBins} ## Build Go based build tools

.PHONY: all check-generated crosscompile generated help images run test tools verify
