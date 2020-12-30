# business_account_service

[![e2e](https://github.com/stefanprodan/business_account_service/workflows/e2e/badge.svg)](https://github.com/stefanprodan/business_account_service/blob/master/.github/workflows/e2e.yml)
[![test](https://github.com/stefanprodan/business_account_service/workflows/test/badge.svg)](https://github.com/stefanprodan/business_account_service/blob/master/.github/workflows/test.yml)
[![cve-scan](https://github.com/stefanprodan/business_account_service/workflows/cve-scan/badge.svg)](https://github.com/stefanprodan/business_account_service/blob/master/.github/workflows/cve-scan.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stefanprodan/business_account_service)](https://goreportcard.com/report/github.com/stefanprodan/business_account_service)
[![Docker Pulls](https://img.shields.io/docker/pulls/stefanprodan/business_account_service)](https://hub.docker.com/r/stefanprodan/business_account_service)

business_account_service is a tiny web application made with Go that showcases best practices of running microservices in Kubernetes.

Specifications:

* Health checks (readiness and liveness)
* Graceful shutdown on interrupt signals
* File watcher for secrets and configmaps
* Instrumented with Prometheus
* Tracing with Istio and Jaeger
* Linkerd service profile
* Structured logging with zap
* 12-factor app with viper
* Fault injection (random errors and latency)
* Swagger docs
* Helm and Kustomize installers
* End-to-End testing with Kubernetes Kind and Helm
* Kustomize testing with GitHub Actions and Open Policy Agent
* Multi-arch container image with Docker buildx and Github Actions
* CVE scanning with trivy

Web API:

* `GET /` prints runtime information
* `GET /version` prints business_account_service version and git commit hash
* `GET /metrics` return HTTP requests duration and Go runtime metrics
* `GET /healthz` used by Kubernetes liveness probe
* `GET /readyz` used by Kubernetes readiness probe
* `POST /readyz/enable` signals the Kubernetes LB that this instance is ready to receive traffic
* `POST /readyz/disable` signals the Kubernetes LB to stop sending requests to this instance
* `GET /status/{code}` returns the status code
* `GET /panic` crashes the process with exit code 255
* `POST /echo` forwards the call to the backend service and echos the posted content
* `GET /env` returns the environment variables as a JSON array
* `GET /headers` returns a JSON with the request HTTP headers
* `GET /delay/{seconds}` waits for the specified period
* `POST /token` issues a JWT token valid for one minute `JWT=$(curl -sd 'anon' business_account_service:9898/token | jq -r .token)`
* `GET /token/validate` validates the JWT token `curl -H "Authorization: Bearer $JWT" business_account_service:9898/token/validate`
* `GET /configs` returns a JSON with configmaps and/or secrets mounted in the `config` volume
* `POST/PUT /cache/{key}` saves the posted content to Redis
* `GET /cache/{key}` returns the content from Redis if the key exists
* `DELETE /cache/{key}` deletes the key from Redis if exists
* `POST /store` writes the posted content to disk at /data/hash and returns the SHA1 hash of the content
* `GET /store/{hash}` returns the content of the file /data/hash if exists
* `GET /ws/echo` echos content via websockets `podcli ws ws://localhost:9898/ws/echo`
* `GET /chunked/{seconds}` uses `transfer-encoding` type `chunked` to give a partial response and then waits for the specified period
* `GET /swagger.json` returns the API Swagger docs, used for Linkerd service profiling and Gloo routes discovery

gRPC API:

* `/grpc.health.v1.Health/Check` health checking

Web UI:

![business_account_service-ui](https://raw.githubusercontent.com/stefanprodan/business_account_service/gh-pages/screens/business_account_service-ui-v3.png)

To access the Swagger UI open `<business_account_service-host>/swagger/index.html` in a browser.

### Guides

* [GitOps Progressive Deliver with Flagger, Helm v3 and Linkerd](https://helm.workshop.flagger.dev/intro/)
* [GitOps Progressive Deliver on EKS with Flagger and AppMesh](https://eks.handson.flagger.dev/prerequisites/)
* [Automated canary deployments with Flagger and Istio](https://medium.com/google-cloud/automated-canary-deployments-with-flagger-and-istio-ac747827f9d1)
* [Kubernetes autoscaling with Istio metrics](https://medium.com/google-cloud/kubernetes-autoscaling-with-istio-metrics-76442253a45a)
* [Autoscaling EKS on Fargate with custom metrics](https://aws.amazon.com/blogs/containers/autoscaling-eks-on-fargate-with-custom-metrics/)
* [Managing Helm releases the GitOps way](https://medium.com/google-cloud/managing-helm-releases-the-gitops-way-207a6ac6ff0e)
* [Securing EKS Ingress With Contour And Let’s Encrypt The GitOps Way](https://aws.amazon.com/blogs/containers/securing-eks-ingress-contour-lets-encrypt-gitops/)

### Install

Helm:

```bash
helm repo add business_account_service https://github.com/BlackspaceInc/BlackspacePlatform/business_account_service

helm upgrade --install --wait frontend \
--namespace test \
--set replicaCount=2 \
--set backend=http://backend-business_account_service:9898/echo \
business_account_service/business_account_service

# Test pods have hook-delete-policy: hook-succeeded
helm test frontend

helm upgrade --install --wait backend \
--namespace test \
--set hpa.enabled=true \
business_account_service/business_account_service
```

Kustomize:

```bash
kubectl apply -k github.com/stefanprodan/business_account_service//kustomize
```

Docker:

```bash
docker run -dp 9898:9898 stefanprodan/business_account_service
```

## Service Level Interationc
---
To read more about service level interactions check out [docs](./docs/readme.md)

## TODO
---
- add support for other models in gqlgen.yml
- emit metrics
- setup kubernetes local environment for end to end testing
- set up end to end testing flow using .github actions and kubernetes
- setup automatic docker container deployment as a github action with all platform support
- learn about k8, kustomize, helm and linkerd deployment options & ensure CI flow to run all tests (unit and end to end tests)

### API Gateway Service
- build api-gateway with schema stitching functionality
- coordinate interactions via the api gateway

### Frontend
- build sign up flow in typescript & react