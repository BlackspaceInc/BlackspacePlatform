---
title: Docker
description:
  Instructions explaining how use Blackspace with Docker on Linux/macOS and
  Windows.
---

Blackspace serbives has images for both Linux/macOS and Windows on
[Docker Hub]({@dockerUrl@}).

## Install Docker

Please follow the [official documentation](https://docs.docker.com/get-docker/).

## Using the image

If you never fetched a Blackspace microservice image, you can run the following example to setup the frontend for example:

```shell
docker run -p 9000:9000 -p 8812:8812 blackspaceInc/frontend-service
```

If you want to make sure that you are running the latest version:

```shell
docker run -p 9000:9000 -p 8812:8812 blackspaceInc/frontend-service:latest
```

### Options

| Argument | Description                 |
| -------- | --------------------------- |
| `-p`     | Port to publish to the host |
| `-v`     | To bind mount a volume      |

#### -p ports

- `-p 9000:9000` for the REST API. The API is
  available on http://localhost:9000
- `-p 8812:8812` for the GRPC wire protocol
- `-p 9009:9009` for the health checkpoint

#### -v volumes

The QuestDB
[root_directory](/docs/reference/configuration/root-directory-structure/) will
be in the following location:

import Tabs from "@theme/Tabs"
import TabItem from "@theme/TabItem"

<Tabs defaultValue="nix" values={[
  { label: "Linux & macOS", value: "nix" },
  { label: "Windows", value: "windows" },
]}>


<TabItem value="nix">


```shell
/root/.frontend-service/frontend-service
```

</TabItem>


<TabItem value="windows">


```shell
C:\frontend-service\frontend-service
```

</TabItem>


</Tabs>
