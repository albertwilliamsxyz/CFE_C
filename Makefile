.PHONY: proto
proto: ## Generate source code from protos
	@bash ./scripts/code_generation/generate_code_from_proto.sh
