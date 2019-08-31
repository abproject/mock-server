# mock-server

> **mock-server** is a tiny (<10 MB) Web Server that allows to emulate RESTful and WebSocket.
> It could be useful for a quick mocking Back-End endpoints during UI prototype implementation, integration or system tests.
>
> **mock-server** could be configured with config file (YAML or JSON) and/or with API.

## Table of Contents

1\.  [Prerequisites](#prerequisites)  
2\.  [Hello World](#helloworld)  
2.1\.  [YAML](#yaml)  
2.2\.  [JSON](#json)  
2.3\.  [API](#api)  
3\.  [Rest](#rest)  
3.1\.  [API](#api-1)  
3.2\.  [YAML](#yaml-1)  
3.3\.  [JSON](#json-1)  

<a name="prerequisites"></a>

## 1\. Prerequisites

[Install Docker](https://docs.docker.com/install/)

<a name="helloworld"></a>

## 2\. Hello World

<a name="yaml"></a>

### 2.1\. YAML

Create file `config.yml` with content:

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

Run in terminal:

```bash
docker run -p 4242:8000 -v ${PWD}/config.yml:/config.yml abezpalov/mock-server -file=config.yml
```

Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:

```bash
curl -v http://localhost:4242/hello
```

```bash
### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/html
...
Hello, World!
```

<a name="json"></a>

### 2.2\. JSON

Create file `config.yml` with content:

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

Run in terminal:

```bash
docker run -p 4242:8000 -v ${PWD}/config.json:/config.json abezpalov/mock-server -file=config.json
```

Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:

```bash
curl -v http://localhost:4242/hello
```

```bash
### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/html
...
Hello, World!
```

<a name="api"></a>

### 2.3\. API

Another way to get the same `Hello World` configuration without config file but by using API requests only.

Run in terminal:

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:

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

e.g., with `curl` **(please copy all 3 code blocks below and paste in terminal)**:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoint \
-H "Content-Type: application/json" \
-d @- << EOF
```

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

```
EOF
```

Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:

```bash
curl -v http://localhost:4242/hello
```

```bash
### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/html
...
Hello, World!
```

<a name="rest"></a>

## 3\. Rest

<a name="api-1"></a>

### 3.1\. [API](docs/REST-API.md)

<a name="yaml-1"></a>

### 3.2\. [YAML](docs/REST-YAML.md)

<a name="json-1"></a>

### 3.3\. [JSON](docs/REST-JSON.md)
