# mock-server

> **mock-server** is a tiny (<10 MB) Web Server that allows to emulate RESTful and WebSocket.
> It could be useful for a quick mocking Back-End endpoints during UI prototype implementation, integration or system tests.
> It could be configured by config file (YAML or JSON) and/or with API.

## Table of Contents

!TOC

## Prerequisites

### Install

- [Docker](https://docs.docker.com/install/)
- _Optional_ [go](https://golang.org/doc/install)

### Example files

> It is possible just to clone current repositiory with examples instead of creating files manually:
>
> ```bash
> git clone https://github.com/abproject/mock-server.git
> cd mock-server
> ```
>
> and then run docker commands with examples from repository (no path changes needed).

## Hello World

!INCLUDE "docs/src/examples/hello-world/hello-world.md", 2

## Examples

### [Files](docs/FILES_EXAMPLE.md)

### [CRUD](docs/CRUD_EXAMPLE.md)

### [Entities](docs/ENTITIES_EXAMPLE.md)

## Config

!INCLUDE "docs/src/config/config-yaml.md", 2

!INCLUDE "docs/src/config/config-json.md", 2

## API

!INCLUDE "docs/src/api/rest/rest-api.md", 2

!INCLUDE "docs/src/api/files/files-api.md", 2

## Models

!INCLUDE "docs/src/models/rest/endpoint/rest-endpoint.md", 2

!INCLUDE "docs/src/models/file/file.md", 2

!INCLUDE "docs/src/models/rest/entity/rest-entity.md", 2
