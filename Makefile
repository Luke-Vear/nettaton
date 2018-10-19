default: help

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

cleanup: ## Clean all the build artefacts up.
	@find . -name "*".zip -exec rm {} \+
.PHONY: cleanup

serve: ## Run a local webserver for the web app.
	go run ./tools/server/server.go --dir ./web
.PHONY: serve

unittest: ## Test the go packages.
	./scripts/pipeline.sh unit
.PHONY: unittest

build: unittest ## Build the artefacts for the wasm webapp and the go services.
	./scripts/pipeline.sh build
.PHONY: build

plan-deploy: build ## Plan a deployment of artefacts and infrastructure.
	./scripts/pipeline.sh plan
.PHONY: plan-deploy

deploy: build ## Deploy the artefacts and infrastructure.
	./scripts/pipeline.sh deploy
	$(MAKE) smoketest
	$(MAKE) cleanup
.PHONY: deploy

smoketest: ## Test the deployed backend go services.
	./scripts/pipeline.sh smoke
.PHONY: smoketest


