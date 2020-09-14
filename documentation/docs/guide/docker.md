---
title: How to use Blackspace with Docker
sidebar_label: Docker
description:
  Tutorial showing how to use Blackspace with Docker.
---

Docker is great to get started in minutes with just a few commands. Follow this
guide to set up and start Blackspace locally. By the end, you will be able to send and
query data to backend services using the REST API.
## Install Docker

Before we start, you will need to install Docker. You can find guides for your
platform [on the official documentation](https://docs.docker.com/get-docker/).

## Blackspace Backend Service Images

With Docker installed, you will need to pull All Blackspace Backend Service Images and create a
container. You can do both in one command using `docker run`:

```shell
git clone BlackspaceInc/BlackspacePlatform
make clean
make up
```

## Container status

You can check the status of your container with **docker ps**. It also lists the
ports we published:

```shell
docker ps
```

```shell title="Result"
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                NAMES
dd363939f261        blackspaceInc/frontend-service     "/app/bin/go -m io…"   3 seconds ago       Up 2 seconds        8812/tcp, 9000/tcp   frosty_gauss
dd363939f262        blackspaceInc/company-service      "/app/bin/go -m io…"   3 seconds ago       Up 2 seconds        8813/tcp, 9001/tcp   frosty_gauss
dd363939f263        blackspaceInc/user-service         "/app/bin/go -m io…"   3 seconds ago       Up 2 seconds        8813/tcp, 9002/tcp   frosty_gauss
dd363939f264        blackspaceInc/auth-service         "/app/bin/go -m io…"   3 seconds ago       Up 2 seconds        8813/tcp, 9003/tcp   frosty_gauss

```

## Importing data and sending queries

🎉 Congratulations, you have a running QuestDB server. You can now start to
interact with it:

## Data persistence

### Restart an existing container

When you stop the container, it will not be removed by Docker. This means that
you can restart it anytime and your data will be accessible:

```shell title="Start container from the  ID obtained with 'docker ps'"
docker start dd363939f261
```

### Re-run `docker run`

If you re-run the command:

```shell
docker run -p 9000:9000 -p 8812:8812 blackspaceInc/frontend-service
```

A new container will be created for the microservice image. This means that the
container will be fresh, any data you may have created previously won't be
accessible.
