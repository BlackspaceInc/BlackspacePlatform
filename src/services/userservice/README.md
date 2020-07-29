# BlackSpace (User Management Service)

[![CircleCI](https://circleci.com/gh/stefanprodan/podinfo.svg?style=svg)](https://circleci.com/gh/stefanprodan/podinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/LensPlatform/BlackSpace)](https://goreportcard.com/report/github.com/LensPlatform/BlackSpace)
[![Docker Pulls](https://img.shields.io/docker/pulls/stefanprodan/podinfo)](https://hub.docker.com/r/stefanprodan/podinfo)

BlackSpace is a platform providing black entrepreneurs with economic independence

The user management services provides basic user access capabilities. It interfaces with the authentication
service in order to provide the system with authentication capabilities. 

The authentication service witholds its own credentials store and from the application context, it is the singular source of truth. 
Prior to performing any updates to the data store of the user management service, we ensure that the authentication service is aware
of the account and the account is not restricted. Due to this authentication service being central to authentication
the data model of the user-management service holds a reference id present in the authentication service data store. 

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

Web API:

* `GET /` prints runtime information
* `GET /version` prints podinfo version and git commit hash 
* `GET /metrics` return HTTP requests duration and Go runtime metrics
* `GET /healthz` used by Kubernetes liveness probe
* `GET /readyz` used by Kubernetes readiness probe
* `POST /readyz/enable` signals the Kubernetes LB that this instance is ready to receive traffic
* `POST /readyz/disable` signals the Kubernetes LB to stop sending requests to this instance
* `GET /env` returns the environment variables as a JSON array
* `GET /headers` returns a JSON with the request HTTP headers
* `POST /token` issues a JWT token valid for one minute `JWT=$(curl -sd 'anon' podinfo:9898/token | jq -r .token)`
* `GET /token/validate` validates the JWT token `curl -H "Authorization: Bearer $JWT" podinfo:9898/token/validate`
* `GET /configs` returns a JSON with configmaps and/or secrets mounted in the `config` volume
* `GET /ws/echo` echos content via websockets `podcli ws ws://localhost:9898/ws/echo`
* `GET /swagger.json` returns the API Swagger docs, used for Linkerd service profiling and Gloo routes discovery
* `POST /v1/user/login` returns a JWT token valid for 24 hours if the user of interest exists in our backend database
* `DELETE v1/user/logout` logs out an existing user from our backend systems
* `POST /v1/user/signup` attempts to a signup a user into BlackSpace 
* `PUT /v1/user/{id}` attempts to update a user in our backend systems
* `GET /v1/user/{id}` attempts to get a user from our backend systems by id
* `DELETE /v1/user/{id}` attempts to delete a user from our backend systems by id
* `/docs` presents swagger api definitions for the service

gRPC API:

* `/grpc.health.v1.Health/Check` health checking

Web UI:

![podinfo-ui](https://raw.githubusercontent.com/stefanprodan/podinfo/gh-pages/screens/podinfo-ui.png)

To access the Swagger UI open `<podinfo-host>/swagger/index.html` in a browser.

### Guides

* [Automated canary deployments with Flagger and Istio](https://medium.com/google-cloud/automated-canary-deployments-with-flagger-and-istio-ac747827f9d1)
* [Kubernetes autoscaling with Istio metrics](https://medium.com/google-cloud/kubernetes-autoscaling-with-istio-metrics-76442253a45a)
* [Managing Helm releases the GitOps way](https://medium.com/google-cloud/managing-helm-releases-the-gitops-way-207a6ac6ff0e)
* [Expose Kubernetes services over HTTPS with Ngrok](https://stefanprodan.com/2018/expose-kubernetes-services-over-http-with-ngrok/)

### Install

Helm:

```bash
helm repo add sp https://stefanprodan.github.io/podinfo

helm upgrade --install --wait frontend \
--namespace test \
--set replicaCount=2 \
--set backend=http://backend-podinfo:9898/echo \
sp/podinfo

helm test frontend --cleanup

helm upgrade --install --wait backend \
--namespace test \
--set hpa.enabled=true \
sp/podinfo
```

Kustomize:

```bash
kubectl apply -k github.com/LensPlatform/BlackSpace/kustomize
```

Docker:

```bash
make help
make start
```