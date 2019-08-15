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
            'Content-Type': 'text/html'
  ```

- Run in terminal

  ```bash
  docker run -p 4242:8000 -v ${PWD}/config.yml:/config.yml abezpalov/mock-server -file=config.yml
  ```

- Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g. with `curl`
  ```bash
  curl -v http://localhost:4242/hello
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
  <b>CLICK HERE</b> If you prefer JSON over YML
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

- Run in terminal

  ```bash
  docker run -p 4242:8000 -v ${PWD}/config.json:/config.json abezpalov/mock-server -file=config.json
  ```

</details>

#### <a name="hello-world-api"></a>API

It is possible to configure the server after starting by using API calls.

- Run in terminal

  ```bash
  docker run -p 4242:8000 abezpalov/mock-server
  ```

- Make `POST` request to URL `http://localhost:4242/_api/rest` with body

  ```json
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
  ```

- e.g with `curl`

  ```bash
  curl -X POST http://localhost:4242/_api/rest \
  -H "Content-Type: application/json" \
  -d @- << EOF
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
  EOF
  ```
