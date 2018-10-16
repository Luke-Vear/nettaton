default: help

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

cleanup: ## Clean all the build artefacts up.
	@find . -name "*".zip -exec rm {} \+
.PHONY: cleanup

backend-unit: ## Test the source for the backend go services.
	./scripts/backend.bash unit
.PHONY: backend-unit

backend-build: backend-unit ## Build the artefacts for the backend go services.
	./scripts/backend.bash build
.PHONY: backend-build

backend-plan: backend-build ## Plan a deployment of artefacts and infrastructure for the backend go services.
	./scripts/backend.bash plan
.PHONY: backend-plan

backend-deploy: backend-build ## Deploy the artefacts and infrastructure for the backend go services.
	./scripts/backend.bash deploy
	$(MAKE) backend-smoketest
	$(MAKE) cleanup
.PHONY: backend-deploy

backend-smoketest: ## Test the deployed backend go services.
	./scripts/backend.bash smoketest ${ENV}
.PHONY: backend-smoketest
