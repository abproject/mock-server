# mock-server

> **mock-server** is a tiny (<10 MB) Web Server that allows to emulate RESTful and WebSocket.
> It could be useful for a quick mocking Back-End endpoints during UI prototype implementation, integration or system tests.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Hello World](#hello-world)
  - [YAML](#hello-world-yaml)
  - [JSON](#hello-world-json)
  - [API](#hello-world-api)
- [Config](#config)
  - [YAML](#config-yaml)
  - [JSON](#config-json)
- [API](#api)
  - [REST](#api-rest)
    - [Endpoint](#api-rest-endpoint)
      - [Description](#api-rest-endpoint-description)
      - [Endpoint Object](#api-rest-endpoint-object)
      - [Endpoint Object Examples](#api-rest-endpoint-object-examples)
      - [Endpoint Request Object](#api-rest-endpoint-request-object)
      - [Endpoint Response Object](#api-rest-endpoint-response-object)
- [Examples](#examples)

## Prerequisites

[Install Docker](https://docs.docker.com/install/)

## Hello World

#### <a name="hello-world-yaml"></a>YAML

- Create file `config.yml` with content:

  ```yaml
  rest:
    endpoints:
      - request:
          method: GET
          path: hello
        response:
          status: 200
          body: Hello, World!
          headers:
            Content-Type: text/html
  ```

- Run in terminal:

  ```bash
  docker run -p 4242:8000 -v ${PWD}/config.yml:/config.yml abezpalov/mock-server -file=config.yml
  ```

- Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:
  ```bash
  curl -v http://localhost:4242/hello
  ```
  ```bash
  # Response
  ...
  < HTTP/1.1 200 OK
  < Content-Type: text/html
  ...
  Hello, World!
  ```

#### <a name="hello-world-json"></a>JSON

<details>
<summary>
  <b>CLICK HERE</b> If you prefer JSON over YAML
</summary>

- Create file `config.yml` with content:

  ```json
  {
    "rest": {
      "endpoints": [
        {
          "request": {
            "method": "GET",
            "path": "hello"
          },
          "response": {
            "status": 200,
            "body": "Hello, World!",
            "headers": {
              "Content-Type": "text/html"
            }
          }
        }
      ]
    }
  }
  ```

- Run in terminal:

  ```bash
  docker run -p 4242:8000 -v ${PWD}/config.json:/config.json abezpalov/mock-server -file=config.json
  ```

</details>
