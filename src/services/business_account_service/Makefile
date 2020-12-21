# Makefile for releasing business_account_service
#
# The release version is controlled from pkg/version

TAG?=latest
NAME:=podinfo
DOCKER_REPOSITORY:=github.com/BlackspaceInc/BlackspacePlatform
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')
EXTRA_RUN_ARGS?=

run:
	go run -ldflags "-s -w -X github.com/stefanprodan/podinfo/pkg/version.REVISION=$(GIT_COMMIT)" cmd/podinfo/* \
	--level=debug --grpc-port=9999 --backend-url=https://httpbin.org/status/401 --backend-url=https://httpbin.org/status/500 \
	--ui-logo=https://raw.githubusercontent.com/stefanprodan/podinfo/gh-pages/cuddle_clap.gif $(EXTRA_RUN_ARGS)

test:
	go test -v -race ./...

build:
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/stefanprodan/podinfo/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podinfo ./cmd/podinfo/*
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/stefanprodan/podinfo/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podcli ./cmd/podcli/*

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

build-charts:
	helm lint charts/*
	helm package charts/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

build-base:
	docker build -f Dockerfile.base -t $(DOCKER_REPOSITORY)/podinfo-base:latest .

push-base: build-base
	docker push $(DOCKER_REPOSITORY)/podinfo-base:latest

test-container:
	@docker rm -f podinfo || true
	@docker run -dp 9898:9898 --name=podinfo $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@TOKEN=$$(curl -sd 'test' localhost:9898/token | jq -r .token) && \
	curl -sH "Authorization: Bearer $${TOKEN}" localhost:9898/token/validate | grep test

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) quay.io/$(DOCKER_IMAGE_NAME):latest
	docker push quay.io/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker push quay.io/$(DOCKER_IMAGE_NAME):latest

version-set:
	@next="$(TAG)" && \
	current="$(VERSION)" && \
	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/podinfo/values.yaml && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/podinfo/values-prod.yaml && \
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" kustomize/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/webapp/frontend/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/webapp/backend/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/bases/frontend/deployment.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" deploy/bases/backend/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

download_proto_dependencies:
	@echo "downloading protoc-gen tool"
	go install github.com/danielvladco/go-proto-gql/protoc-gen-gql && go install github.com/danielvladco/go-proto-gql/protoc-gen-gogql && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate_grpc_definitions: download_proto_dependencies
	@echo "setting up grpc service schema definition via protobuf"
	protoc -I. \
			-I$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
			--govalidators_out=./pkg/grpc/ \
			--go_out=./pkg/grpc --go-grpc_out=./pkg/grpc/ --gorm_out="engine=postgres:./pkg/grpc/" ./proto/*.proto

generate_graphql_schema_from_proto_definitions: generate_grpc_definitions
	proto2graphql --input ./proto/schema.proto --output ./pkg/graphql/schema/schema.graphql

generate_graphql_api: generate_graphql_schema_from_proto_definitions
	# go:generate go run github.com/99designs/gqlgen
	cd pkg/graphql && go run github.com/99designs/gqlgen schema/*.graphql

down:
	docker-compose down
	./scripts/clean.sh

up:
	docker-compose -f docker-compose.dev.yml -f docker-compose.monitoring.dev.yml config
	docker-compose -f docker-compose.dev.yml -f docker-compose.monitoring.dev.yml up --remove-orphans

kube:
	kompose convert -f docker-compose.prod.yml -o ./deployments/kubemanifests.yml
	kubectl apply -f kubemanifests.yml

deploy-service:
	kubectl apply -f ./deployments/kubemanifests.yml
	kubectl apply -f ./deployments/service.yml

deploy-dev:
	@echo "deploying service to development environment First we attach the linkerd proxy sidecar to the service then we deploy"
	kustomize build ./deploy/overlays/dev | kubectl apply -f-

deploy-staging:
	@echo "deploying service to staging environment. First we attach the linkerd proxy sidecar to the service then we deploy"
	kustomize build ./deploy/overlays/staging | kubectl apply -f-

deploy-production:
	@echo "deploying service to production environment First we attach the linkerd proxy sidecar to the service then we deploy"
	kustomize build ./overlays/production | kubectl apply -f-

deploy: deploy-dev deploy-staging
	@echo "deploying"

start:
	# starts the necessary services in order to permit e2e testing
	@echo starts the required support container services to allow e2e testing
	docker-compose -f docker-compose.jaeger.yaml \
				   -f docker-compose.postgres.yaml \
				   -f docker-compose.authentication_handler_service.yaml \
				   -f docker-compose.dev.yaml up --remove-orphans

stop:
	@echo stopping the container services used for testing
	docker-compose -f docker-compose.jaeger.yaml \
				   -f docker-compose.postgres.yaml \
 				   -f docker-compose.authentication_handler_service.yaml \
 				   -f docker-compose.dev.yaml down
