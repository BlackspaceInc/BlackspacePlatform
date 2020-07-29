# Makefile for releasing podinfo
#
# The release version is controlled from pkg/version
TAG?=latest
NAME:=blackspacebackend
TARGET:=$(NAME)
DOCKER_REPOSITORY:=blackspaceinc
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')
SHELL := /bin/bash
REV := $(shell git rev-parse HEAD)
CHANGES := $(shell test -n "$$(git status --porcelain)" && echo '+CHANGES' || true)
VERSION := $(shell cat VERSION)
OS := darwin freebsd linux openbsd
ARCH := 386 amd64
LDFLAGS := -X github.com/kwilczynski/$(NAME)/$(NAME).Revision=$(REV)$(CHANGES)
GPG_SIGNING_KEY :=
DOCKER_NETWORK=backend
SERVICE_NAME_NAME=user-management-service

build:
	GO111MODULE=on GIT_COMMIT=$$(git rev-list -1 HEAD) && GO111MODULE=on CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/BlackspaceInc/Backend/user-management-service/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podinfo ./cmd/podinfo/*
	GO111MODULE=on GIT_COMMIT=$$(git rev-list -1 HEAD) && GO111MODULE=on CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/BlackspaceInc/Backend/user-management-service/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/podcli ./cmd/podcli/*

# Addional makefile utilities

.PHONY: \
	help \
	default \
	clean \
	clean-artifacts \
	clean-releases \
	clean-vendor \
	tools \
	deps \
	test \
	coverage \
	vet \
	errors \
	lint \
	imports \
	fmt \
	env \
	build \
	build-all \
	doc \
	release \
	package-release \
	sign-release \
	check \
	vendor \
	version \
	proto	\
	swagger

all: imports fmt lint vet errors build

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    run                Runs the service'
	@echo '    help               Show this help screen.'
	@echo '    clean              Remove binaries, artifacts and releases.'
	@echo '    clean-artifacts    Remove build artifacts only.'
	@echo '    clean-releases     Remove releases only.'
	@echo '    clean-vendor       Remove content of the vendor directory.'
	@echo '    tools              Install tools needed by the project.'
	@echo '    deps               Download and install build time dependencies.'
	@echo '    test               Run unit tests.'
	@echo '    coverage           Report code tests coverage.'
	@echo '    vet                Run go vet.'
	@echo '    errors             Run errcheck.'
	@echo '    lint               Run golint.'
	@echo '    imports            Run goimports.'
	@echo '    fmt                Run go fmt.'
	@echo '    env                Display Go environment.'
	@echo '    build              Build project for current platform.'
	@echo '    build-all          Build project for all supported platforms.'
	@echo '    doc                Start Go documentation server on port 8080.'
	@echo '    release            Package and sing project for release.'
	@echo '    package-release    Package release and compress artifacts.'
	@echo '    sign-release       Sign release and generate checksums.'
	@echo '    check              Verify compiled binary.'
	@echo '    vendor             Update and save project build time dependencies.'
	@echo '    version            Display Go version.'
	@echo '    proto              Update tables via proto definitions.'
	@echo '    swagger            Creates an updated swagger file.'
	@echo '    auth-service       Starts up authentication service.'
	@echo '    zipkin             Starts up zipkin tracing container.'
	@echo '    queue-up           Starts up rabbitmq instance.'
	@echo '    client             Generates service client based on swagger specs.'
	@echo '    start              Starts all dependent services via docker-compose.'
	@echo '    dev                Starts a development friendly auto reload server and allnecessary service dependencies.'
	@echo ''
	@echo 'Targets run by default are: imports, fmt, lint, vet, errors and build.'
	@echo ''


build-charts:
	helm lint charts/*
	helm package charts/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) -f ./Dockerfile.prod ..

build-dev-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(TAG) -f ./Dockerfile.dev ..
	docker run -p 9898:9898 --name=user-service $(DOCKER_IMAGE_NAME):$(TAG)

test-dev-container:
	@docker run -dp 9898:9898 --name=user-service $(DOCKER_IMAGE_NAME):$(TAG)

test-container:
	@docker rm -f user-manegement-service || true
	@docker run -dp 9898:9898 --name=user-service $(DOCKER_IMAGE_NAME):$(VERSION)
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
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/podinfo/Chart.yaml && \
	sed -i '' "s/podinfo:$$current/podinfo:$$next/g" kustomize/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

print-%:
	@echo $* = $($*)

run:
	GO111MODULE=on go run -ldflags "-s -w -X github.com/BlackspaceInc/Backend/user-management-service/pkg/version.REVISION=$(GIT_COMMIT)" cmd/podinfo/* \
	--level=debug --grpc-port=9999 --backend-url=https://httpbin.org/status/401 --backend-url=https://httpbin.org/status/500 \
	--ui-logo=https://media.giphy.com/media/l0ExbNdlJFGRphMR2/giphy.gif

clean: clean-artifacts clean-releases
	go clean -i ./...
	rm -vf \
	  $(CURDIR)/coverage.* \
	  $(CURDIR)/packer-provisioner-itamae-local_*

clean-artifacts:
	rm -Rf artifacts/*

clean-releases:
	rm -Rf releases/*

clean-vendor:
	find $(CURDIR)/vendor -type d -print0 2>/dev/null | xargs -0 rm -Rf

clean-containers:
	@echo 'running cleanup scrip to remove all docker containers, volumes, networks, cached values'
	./scripts/cleanup.sh
	docker network rm $(DOCKER_NETWORK)

clean-all: clean clean-artifacts clean-vendor clean-containers

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/golang/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox

deps:
	godep restore

test: # init deps
	GO111MODULE=on go test -v -race ./...

coverage: deps
	gocov test ./... > $(CURDIR)/coverage.out 2>/dev/null
	gocov report $(CURDIR)/coverage.out
	if test -z "$$CI"; then \
	  gocov-html $(CURDIR)/coverage.out > $(CURDIR)/coverage.html; \
	  if which open &>/dev/null; then \
	    open $(CURDIR)/coverage.html; \
	  fi; \
	fi

vet:
	go vet -v ./...

errors:
	errcheck -ignoretests -blank ./...

lint:
	golint ./...

imports:
	goimports -l -w .

fmt:
	go fmt ./...

env:
	@go env

build: deps
	go build -v \
	   -ldflags "$(LDFLAGS)" \
	   -o "$(TARGET)" .

build-all: deps
	mkdir -v -p $(CURDIR)/artifacts/$(VERSION)
	gox -verbose \
	    -os "$(OS)" -arch "$(ARCH)" \
	    -ldflags "$(LDFLAGS)" \
	    -output "$(CURDIR)/artifacts/$(VERSION)/{{.OS}}_{{.Arch}}/$(TARGET)" .
	cp -v -f \
	   $(CURDIR)/artifacts/$(VERSION)/$$(go env GOOS)_$$(go env GOARCH)/$(TARGET) .

doc:
	godoc -http=:8080 -index

release: package-release sign-release

package-release:
	@test -x $(CURDIR)/artifacts/$(VERSION) || exit 1
	mkdir -v -p $(CURDIR)/releases/$(VERSION)
	for release in $$(find $(CURDIR)/artifacts/$(VERSION) -mindepth 1 -maxdepth 1 -type d 2>/dev/null); do \
	  platform=$$(basename $$release); \
	  pushd $$release &>/dev/null; \
	  zip $(CURDIR)/releases/$(VERSION)/$(TARGET)_$${platform}.zip $(TARGET); \
	  popd &>/dev/null; \
	done

sign-release:
	@test -x $(CURDIR)/releases/$(VERSION) || exit 1
	pushd $(CURDIR)/releases/$(VERSION) &>/dev/null; \
	shasum -a 256 -b $(TARGET)_* > SHA256SUMS; \
	if test -n "$(GPG_SIGNING_KEY)"; then \
	  gpg --default-key $(GPG_SIGNING_KEY) -a \
	      -o SHA256SUMS.sign -b SHA256SUMS; \
	fi; \
	popd &>/dev/null

check:
	@test -x $(CURDIR)/$(TARGET) || exit 1
	if $(CURDIR)/$(TARGET) --version | grep -qF '$(VERSION)'; then \
	  echo "$(CURDIR)/$(TARGET): OK"; \
	else \
	  exit 1; \
	fi

vendor: deps
	godep save

version:
	@go version

proto:
	@echo "setting up service schema definition via protobuf"
	protoc -I. \
			-I$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
            --govalidators_out=./pkg/ \
			--go_out=plugins=grpc:./pkg/ --gorm_out="engine=postgres:./pkg/" ./models/*.proto


check-install-swagger:
	@echo "installing required libraries for swagger doc. generation"
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check-install-swagger
	@echo "generating swagger documentation"
	GO111MODULE=off swagger generate spec -o ./pkg/api/docs/swagger.yaml --scan-models
	GO111MODULE=off swagger generate spec -o ./pkg/api/docs/swagger.json --scan-models

network:
	@echo "create docker network"
	docker network create $(DOCKER_NETWORK)

auth-service: network zipkin queue-up
	@echo "spinning up authentication service via docker command"
	docker run -it -d --rm \
      --publish 8404:3000 \
      --network $(DOCKER_NETWORK) \
      -e AUTHN_URL=localhost:8404 \
      -e APP_DOMAINS=localhost \
      -e DATABASE_URL=sqlite3://:memory:?mode=memory\&cache=shared \
      -e SECRET_KEY_BASE='my-authn-test-secret' \
      -e HTTP_AUTH_USERNAME=blackspaceinc \
      -e HTTP_AUTH_PASSWORD=blackspaceinc \
      --name auth \
      keratin/authn-server:latest \
      sh -c "./authn migrate && ./authn server"

queue-up:
	@echo "spinning up rabbitMQ instance"
	docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 --network $(DOCKER_NETWORK) rabbitmq:3-management

zipkin:
	@echo 'zipkin server up'
	GO111MODULE=off docker run -it -d --rm -p 9411:9411 --network $(DOCKER_NETWORK) openzipkin/zipkin

init: network zipkin auth-service queue-up
	@echo 'started dependent service containers on the same docker network'

dev: init
	@echo "development server up with configured auto-reload"
	GO111MODULE=on modd

start:
	@echo "spinning up all required containers"
	./scripts/composeup.sh

client: swagger 
	@echo "starting service client generation"
	GO111MODULE=off cd sdk && swagger generate client -f ../pkg/api/docs/swagger.json -a operationa -m models, -s server -c client -A user-management-client
