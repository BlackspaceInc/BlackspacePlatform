signoff: ## Signsoff all previous commits since branch creation
	scripts/signoff.sh

release: # Invokes a script to automate the creation of a release

spin-up-kube: # Spins up local mini kube cluster