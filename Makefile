default: help

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

clean: ## Clean all the build artefacts up.
	./scripts/pipeline.sh clean
.PHONY: cleanup

serve: ## Run a local webserver for the web app.
	./scripts/pipeline.sh serve
.PHONY: serve

unittest: ## Test the go packages.
	./scripts/pipeline.sh unit
.PHONY: unittest

build: unittest frontend backend ## Build the artefacts for the wasm webapp and the go services.

frontend: web/dist

web/dist:
	./scripts/pipeline.sh build_frontend

backend: cmd/createquestion/handler.zip cmd/readquestion/handler.zip cmd/answerquestion/handler.zip
	
cmd/createquestion/handler.zip cmd/readquestion/handler.zip cmd/answerquestion/handler.zip:
	./scripts/pipeline.sh build_backend

plan-deploy: ## Plan a deployment of artefacts and infrastructure.
	./scripts/pipeline.sh plan
.PHONY: plan-deploy

deploy: ## Deploy the artefacts and infrastructure.
	./scripts/pipeline.sh deploy
.PHONY: deploy

destroy: ## Destroy infrastructure.
	./scripts/pipeline.sh destroy
.PHONY: deploy

smoketest: ## Test the deployed backend go services.
	./scripts/pipeline.sh smoke
.PHONY: smoketest