---
id: services
sidebar_label: Service Architecture
title: Backend Service Architecture
---

:::tip
If you want to use a different folder, replace */mnt/qaboard* with your path in `docker-compose.yml` and *services/nginx/conf.d/qaboard.conf*.
:::

:::danger
If you want to use a different folder, replace */mnt/qaboard* with your path in `docker-compose.yml` and *services/nginx/conf.d/qaboard.conf*.
:::

:::info
If you want to use a different folder, replace */mnt/qaboard* with your path in `docker-compose.yml` and *services/nginx/conf.d/qaboard.conf*.
:::

:::caution
If you want to use a different folder, replace */mnt/qaboard* with your path in `docker-compose.yml` and *services/nginx/conf.d/qaboard.conf*.
:::

:::important
If you want to use a different folder, replace */mnt/qaboard* with your path in `docker-compose.yml` and *services/nginx/conf.d/qaboard.conf*.
:::

:::note Shared Storage?
Later, read how to setup [**NFS**](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nfs-mount-on-ubuntu-18-04) or [**Samba**](https://www.digitalocean.com/community/tutorials/how-to-set-up-a-samba-share-for-a-small-organization-on-ubuntu-16-04). If you need fine-tuning read about [options for NFS volumes](https://docs.docker.com/compose/compose-file/#volume-configuration-reference) in *docker-compose.yml*.
:::

:::note Working in the cloud?
Use file-base storage like [AWS EFS](https://aws.amazon.com/en/efs/) or [GCP Filestore](https://cloud.google.com/filestore).

We plan on supporting blob-stores like AWS **S3**. <a href="mailto:arthur.flam@gmail.com">Contact us</a> or [create an issue](https://github.com/samsung/qaboard/issues) if it would help.
:::
