---
id: production-ready-checklist
sidebar_label: Backend Production Checklist
title: Backend Production Ready Checklist
---

# Production Ready Checklist <!-- omit in toc -->

This document considers the production readiness of a solution from an infrastructure and applicaton perspective.

- [Backups and Restoring](#backups-and-restoring)
- [Code Efficiency and Performance](#code-efficiency-and-performance)
- [Code Readability](#code-readability)
- [Continuous Delivery](#continuous-delivery)
- [Continuous Integration](#continuous-integration)
- [Cost](#cost)
- [Documentation](#documentation)
- [High Availability](#high-availability)
- [Logging](#logging)
- [Maintenance](#maintenance)
- [Monitoring](#monitoring)
- [Networking](#networking)
- [Performance](#performance)
- [Provisioning and Deployment](#provisioning-and-deployment)
- [Scalability](#scalability)
- [Security](#security)
- [Testing](#testing)

## Backups and Restoring

- [ ] All code is in shared repository
- [ ] Configure backup
- [ ] Configure cross-account backup
- [ ] Event Storage can be backed up
- [ ] Event Storage is routinely backed up
- [ ] Event Storage restore has been tested
- [ ] Solution can be easily deployed via pipeline
- [ ] Test your backups

## Code Efficiency and Performance

- [ ] Code performance has been measured and tested

## Code Readability

- [ ] Code follows standard naming style conventions
- [ ] Code has sufficient documentation for increased readability
- [ ] Code is being formatted automatically
- [ ] Code is being linted automatically

## Continuous Delivery

- [ ] Automate your deployments
- [ ] Create deployment environments
- [ ] Define your infrastructure as code
- [ ] Do zero-downtime deployments
- [ ] Promote artifacts
- [ ] Roll back in case of failure
- [ ] Set up immutable infrastructure
- [ ] Set up per-environment configuration
- [ ] Test your infrastructure code
- [ ] Use canary deployments
- [ ] Use feature toggles

## Continuous Integration

- [ ] Configure a build system
- [ ] Configure static analysis
- [ ] Do code reviews
- [ ] Pick a Version Control System
- [ ] Publish versioned artifacts
- [ ] Set up a build server
- [ ] Set up automated tests
- [ ] Set up automatic code formatting
- [ ] Use dependency management

## Cost

- [ ] Clean up old data with S3 Lifecycle settings
- [ ] Clean up unused resources
- [ ] Cost analysis has been done against all resources used by Solution
- [ ] Create billing alarms
- [ ] Learn to analyze your AWS bill
- [ ] Pick proper EC2 Instance types and sizes
- [ ] Resources are on a pay per use model
- [ ] Shut down EC2 Instances and RDS DBs when not using them
- [ ] Use Auto Scaling
- [ ] Use Docker when possible
- [ ] Use Lambda when possible
- [ ] Use Reserved EC2 Instances for dedicated work
- [ ] Use Spot EC2 Instances for background jobs

## Documentation

- [ ] An architecture diagram is in place
- [ ] Code has sufficient comments and docstrings
- [ ] Documentation is kept within the repository
- [ ] Operational Runbooks are included
- [ ] Updated and accurate README is included in repository

## High Availability

- [ ] Configure Auto Recovery
- [ ] Configure graceful degradation
- [ ] Perform load tests and use chaos engineering
- [ ] Solution has load balancing for all requests
- [ ] Solution is run concurrently in multiple regions
- [ ] Solution requires no downtime for maintenance

## Logging

- [ ] Solution has logging in place through Cloudwatch Logs

## Maintenance

- [ ] Configure disk space
- [ ] Set up cron jobs
- [ ] Solution has sufficient documentation for code maintenance to take place
- [ ] Solution has sufficient high availability for maintenance to take place
- [ ] Solution has sufficient monitoring for maintenance to take place

## Monitoring

- [ ] Configure services for observability
- [ ] Set up alerts
- [ ] Solution has alerting in place
- [ ] Solution has desired monitoring defined
- [ ] Solution has monitoring in place
- [ ] Store logs
- [ ] Track application metrics
- [ ] Track availability metrics
- [ ] Track business metrics
- [ ] Track server metrics

## Networking

- [ ] Configure DNS using Route 53
- [ ] Configure Network ACLs
- [ ] Configure Security Groups
- [ ] Configure Static IPs
- [ ] Scalability in place for sudden increase in incoming network traffic
- [ ] Set up subnets
- [ ] Set up VPCs
- [ ] Solution has load balancing in front of exposed HTTP endpoints
- [ ] Stress test has been done on volume of webhook ingestion

## Performance

- [ ] Degraded performance is handled gracefully
- [ ] Expected performance in ideal conditions is defined
- [ ] Stress test has been done on solution performance

## Provisioning and Deployment

- [ ] Build AMIs
- [ ] Build Docker images
- [ ] Configure CPU, memory and GC settings
- [ ] Configure hard drives
- [ ] Deploy AMIs using Auto Scaling Groups
- [ ] Deploy Docker images using ECS, EKS or Fargate
- [ ] Deploy serverless apps using Lambda and API Gateway
- [ ] Solution can be deployed automatically via a pipeline
- [ ] Solution has release tags
- [ ] Solution is defined in Infrastructure as code
- [ ] Solution pipeline has multiple environments for dev and prod

## Scalability

- [ ] Set up load balancing
- [ ] Solution resources can auto scale to handle workload
- [ ] Stress test has been done on autoscaling
- [ ] Use multiple Availability Zones
- [ ] Use multiple Instances

## Security

- [ ] API usage requires credentials or authentication
- [ ] AWS IAM Permissions are minimal, users only have permissions they need
- [ ] Configure encryption at rest
- [ ] Create a password policy and enforce MFA
- [ ] Create cross-account IAM Roles
- [ ] Create IAM Groups
- [ ] Create IAM Roles
- [ ] Create IAM Users
- [ ] Data is encrypted at REST
- [ ] Deploy a Bastion Host
- [ ] Deploy a VPN Server
- [ ] Go through a security audit
- [ ] Go through the OWASP Top 10
- [ ] Network traffic is over HTTPS
- [ ] No secrets exposed in repository
- [ ] Record audit Logs
- [ ] Set up a secrets management solution
- [ ] Set up SSH access
- [ ] Sign up for security advisories
- [ ] Use server hardening practices

## Testing

- [ ] Automated coverage report is being carried out
- [ ] Automated testing is being utilised in the pipeline
- [ ] Testing safety net is in place
