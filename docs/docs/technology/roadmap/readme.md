---
id: blackspace-roadmap
sidebar_label: General Roadmap
title: Blackspace General Product Roadmap
---

# Project Plan
<!--{h1:.massive-header.-with-tagline}-->

## Contents

- [Project Plan](#project-plan)
  - [Contents](#contents)
  - [Backend Development Deliverables](#backend-development-deliverables)
    - [Refactoring](#refactoring)
    - [Testing](#testing)
    - [Documentation](#documentation)
    - [Enhancements](#enhancements)
    - [Instrumentation](#instrumentation)
    - [Service Mesh](#service-mesh)
    - [Deployment & Clusters](#deployment--clusters)
    - [Microservice Enhancements](#microservice-enhancements)
    - [Security](#security)
  - [System Architecture Deliverables](#system-architecture-deliverables)
  - [Frontend Development Deliverables](#frontend-development-deliverables)

## Backend Development Deliverables
### Refactoring
* Refactor backend service code & add concise comments (function, interfaces, ....)
### Testing
* Unit test all packages (85% coverage minimum achieved per project)
* Provide E2E api tests at the service livereload
* Load test all backend service api
### Documentation
* Document service interaction formally
### Enhancements
* Assert service reads from env variables
  * Define .env file for all environments (dev, staging, prod)
* Enforece security for all backend microservice (follow best-practices)
* Apply more robust http middlewares
### Instrumentation
* Emit/implement prometheus metric emitting logic 
* Implement distributed tracing for all services
* Implement log aggregation logic via webhooks with `ELK` stack
* Provision helm chart template for each service to expedite deployment
  * Define helm chart at the environment level (namespace=prod, dev, staging)
* Implement circuit breaker logic for all services with smart backoff techniques in face of failure
* Provide GraphQl api for all services
* Provide GRPC endpoints for all Services
  * Provision grpc middlewares
* Consolidate service level contracts via protobuff
### Service Mesh 
* Add Istio manifests
  * Configure istio for all environments (feature flag)
* Add Linkerd manifests
  * Configure linkerd for all environments (feature flag)
### Deployment & Clusters
* Define redis cluster for kubernetes
* Define message bus cluster for kubernetes
* Define monitoring cluster for kubernetes
### Microservice Enhancements
* Enhance email service functionality
  * Emit metrics & Logs
  * Provide distributed tracing func.
  * Provide more authentication functionality through service
    * password reset
    * email resets
    * change password
    * change email
    * invitation only-signup
    * email confirmation
### Security
* Secure all backend Services
* Secure database clusters
* Enforce kubernetes security best practices
* Enhance reverse proxy (follow security best practices)

## System Architecture Deliverables
* ReArchitect services/backend 
* Implement an api gateway 
  * Choose between in-house & open-source api gateway
    * Ensure Api gateway has graphql schema stitching functionality
* Enhance nginx reverse proxy
* Experiment with traefik vs nginx
  
## Frontend Development Deliverables
* Enhance user experience
  * Use templates & themes
* Convert project from vanilla javascript to typescript
* Provide unit testing functionality for all react components
* Enhance test coverage to 85%

**BlackSpacePlatform** © 2020
