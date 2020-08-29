
signoff: ## Signsoff all previous commits since branch creation
	scripts/signoff.sh

release-docker: ## Release to Docker Hub
	@scripts/release-docker.sh

release-github: ## Release to Github
	@scripts/release-github.rb

release-homebrew: ## Release to timberio Homebrew tap
	@scripts/release-homebrew.sh

release-prepare: ## Prepares the release with metadata and highlights
	@scripts/release-prepare.rb

release-push: ## Push new Blackspace version
	@scripts/release-push.sh

release-rollback: ## Rollback pending release changes
	@scripts/release-rollback.rb

release-s3: ## Release artifacts to S3
	@scripts/release-s3.sh

release-helm: ## Package and release Helm Chart
	@scripts/release-helm.sh

check-meta: ## Check that all /.meta file are valid
	./scripts/check-meta.sh