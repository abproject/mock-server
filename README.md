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

#### Config

##### <a name="config-yaml"></a>YAML

##### <a name="config-json"></a>JSON

#### <a name="hello-world-api"></a>API

Another way to get the same `Hello World` configuration without config file but by using API requests only.

- Run in terminal:

  ```bash
  docker run -p 4242:8000 abezpalov/mock-server
  ```

- Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:

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

- e.g., with `curl`:

  ```bash
  curl -X POST http://localhost:4242/_api/rest/endpoint \
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

## API

#### <a name="api-rest"></a>REST

##### <a name="api-rest-endpoint"></a>Endpoint

> **Response Headers:**
>
> - `Content-Type: application/json`

###### <a name="api-rest-endpoint-description"></a>Description

| Path                      | Method   | Description                                                         |                 Request Body                 |                    Response Body                     | Success Status | Failed Status |
| ------------------------- | -------- | ------------------------------------------------------------------- | :------------------------------------------: | :--------------------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoint`     | `GET`    | Returns the list of all endpoints configurations                    |                      -                       | List of [Endpoint Object](#api-rest-endpoint-object) |    **200**     |       -       |
| `/_api/rest/endpoint`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Object](#api-rest-endpoint-object) |     [Endpoint Object](#api-rest-endpoint-object)     |    **201**     |       -       |
| `/_api/rest/endpoint`     | `DELETE` | Deletes all endpoints configuration                                 |                      -                       |                          -                           |    **204**     |       -       |
| `/_api/rest/endpoint/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                      -                       |     [Endpoint Object](#api-rest-endpoint-object)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Object](#api-rest-endpoint-object) |     [Endpoint Object](#api-rest-endpoint-object)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                      -                       |                          -                           |    **204**     |    **404**    |

###### <a name="api-rest-endpoint-object"></a>Endpoint Object

| Field Name | Type                                                           | Description                           |
| ---------- | -------------------------------------------------------------- | ------------------------------------- |
| `id`       | `string`                                                       | Unique Endpoint ID. **Response only** |
| `request`  | [Endpoint Request Object](#api-rest-endpoint-request-object)   | Request configuration object          |
| `response` | [Endpoint Response Object](#api-rest-endpoint-response-object) | Response configuration object         |

###### <a name="api-rest-endpoint-object-examples"></a>Endpoint Object Examples

**Request:**

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
      "Content-Type": "plain/text"
    }
  }
}
```

**Response:**

```json
{
  "id": "52fdfc072182454f963f5f0f9a621d72",
  "request": {
    "method": "GET",
    "path": "hello",
    "pathReg": "",
    "headers": null
  },
  "response": {
    "body": "Hello, World!",
    "bodyFile": "",
    "status": 200,
    "headers": {
      "Content-Type": "plain/text"
    }
  }
}
```

###### <a name="api-rest-endpoint-request-object"></a>Endpoint Request Object

| Field Name | Type                  | Description                                                                                                                              |
| ---------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `method`   | `string`              | Method name, e.g., `GET`, `POST`, `DELETE`, `PUT`. <br>Empty string means all type of request. <br>**Default:** empty string             |
| `path`     | `string`              | Request Endpoint path, e.g., `/my-path`. <br>**Default:** empty string                                                                   |
| `pathReg`  | `string`              | Request Endpoint path as regular expression. If `pathReg` value is not empty then `path` value is ignored. <br>**Default:** empty string |
| `headers`  | `map<string, string>` | Request Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                              |

###### <a name="api-rest-endpoint-response-object"></a>Endpoint Response Object

| Field Name | Type                  | Description                                                                                                                        |
| ---------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `body`     | `any`                 | Response body, could be anything including JSON objects. <br>**Default:** empty string                                             |
| `bodyFile` | `string`              | Response will be content of the file. If `bodyFile` value is not empty then `body` value is ignored. <br>**Default:** empty string |
| `status`   | `integer`             | Response HTTP status code. <br>**Default:** `0`                                                                                    |
| `headers`  | `map<string, string>` | Response Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                       |
