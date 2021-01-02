# authentication_handler_service
[![Go Reference](https://pkg.go.dev/badge/BlackspaceInc/BlackspacePlatform.svg)](https://pkg.go.dev/BlackspaceInc/BlackspacePlatform)
[![Go Report Card](https://goreportcard.com/badge/github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service)](https://goreportcard.com/report/github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service)
[![Docker Pulls](https://img.shields.io/docker/pulls/blackspaceinc/shopper_service)](https://hub.docker.com/r/blackspaceinc/authentication_handler_service)
[![Issues](https://img.shields.io/github/issues/BlackspaceInc/BlackspacePlatform)](https://img.shields.io/github/issues/BlackspaceInc/BlackspacePlatform)
[![Forks](https://img.shields.io/github/forks/BlackspaceInc/BlackspacePlatform)](https://img.shields.io/github/forks/BlackspaceInc/BlackspacePlatform)
[![Stars](	https://img.shields.io/github/stars/BlackspaceInc/BlackspacePlatform)](https://img.shields.io/github/stars/BlackspaceInc/BlackspacePlatform)
[![License](https://img.shields.io/github/license/BlackspaceInc/BlackspacePlatform)](https://img.shields.io/github/license/BlackspaceInc/BlackspacePlatform)


authentication_handler_service is built on the podinfo open source golang microservice template which showcases the best
 practices of running microservices in Kubernetes.

 Please reference [SLA Details](./docs/authentication.md) for further information specific to the various SLAs (Service Level Agreements)

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
* `GET /version` prints authentication_handler_service version and git commit hash
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
* `POST /token` issues a JWT token valid for one minute `JWT=$(curl -sd 'anon' authentication_handler_service:9898/token | jq -r .token)`
* `GET /token/validate` validates the JWT token `curl -H "Authorization: Bearer $JWT" authentication_handler_service:9898/token/validate`
* `GET /configs` returns a JSON with configmaps and/or secrets mounted in the `config` volume
* `POST/PUT /cache/{key}` saves the posted content to Redis
* `GET /cache/{key}` returns the content from Redis if the key exists
* `DELETE /cache/{key}` deletes the key from Redis if exists
* `POST /store` writes the posted content to disk at /data/hash and returns the SHA1 hash of the content
* `GET /store/{hash}` returns the content of the file /data/hash if exists
* `GET /ws/echo` echos content via websockets `podcli ws ws://localhost:9898/ws/echo`
* `GET /chunked/{seconds}` uses `transfer-encoding` type `chunked` to give a partial response and then waits for the specified period
* `GET /swagger.json` returns the API Swagger docs, used for Linkerd service profiling and Gloo routes discovery
* `POST /v1/account/create` creates an account record from the context of the authentication service through a distributed transaction. Request
 body must be a json string comprised of the following `{"email": sample@gmail.com, "password": sample_password}`
 * `DELETE /v1/account/delete/{id}` deletes a user account record from the context of the authentication service through a distributed transaction.
 * `GET /v1/account/{id}` gets an account from the context of the authentication service by ID.
 * `POST /v1/account/lock/{id}` locks an account from the context of the authentication service by ID.
 * `POST /v1/account/login` logs in a user into the system and returns a jwt token which must be used to authenticate all requests. Request body
  must be a json string comprised of the following `{"email": sample@gmail.com, "password": sample_password}`
 * `POST /v1/account/logout/{id{}` logs out a user account from the system
 * `POST /v1/account/unlock/{id}` unlocks an account from the context of the authentication service by ID.
 * `POST /v1/account/update/{id}` updates a user account's email address from the context of the authentication service.

gRPC API:
* `/grpc.health.v1.Health/Check` health checking

Web UI:

![authentication_handler_service-ui](https://raw.githubusercontent.com/github.com/blackspaceInc/BlackspacePlatform/authentication_handler_service/gh-pages/screens/authentication_handler_service-ui-v3.png)

To access the Swagger UI open `<localhost:port>/swagger/index.html` in a browser.

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
helm repo add authentication_handler_service https://github.com/blackspaceInc/BlackspacePlatform/authentication_handler_service

helm upgrade --install --wait frontend \
--namespace test \
--set replicaCount=2 \
--set backend=http://backend-authentication_handler_service:9898/echo \
authentication_handler_service/authentication_handler_service

# Test pods have hook-delete-policy: hook-succeeded
helm test frontend

helm upgrade --install --wait backend \
--namespace test \
--set hpa.enabled=true \
authentication_handler_service/authentication_handler_service
```

Kustomize:

```bash
kubectl apply -k github.com/blackspaceInc/BlackspacePlatform/src/services/authentication_handler_service//kustomize
```

Docker:

```bash
docker run -dp 9898:9898 github.com/blackspaceInc/BlackspacePlatform/authentication_handler_service
```

### Starting Locally
To start the service and its dependencies locally, in the command line, run `make start-e2e-dependencies`. This will spin up a set of docker containers connected
 to the same docker network.
