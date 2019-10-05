# mock-server

> **mock-server** is a tiny (<10 MB) Web Server that allows to emulate RESTful and WebSocket.
> It could be useful for a quick mocking Back-End endpoints during UI prototype implementation, integration or system tests.
> It could be configured by config file (YAML or JSON) and/or with API.

## Table of Contents

1\.  [Prerequisites](#prerequisites)  
1.1\.  [Install](#install)  
1.2\.  [Example files](#examplefiles)  
2\.  [Hello World](#helloworld)  
2.1\.  [YAML](#yaml)  
2.2\.  [JSON](#json)  
2.3\.  [API](#api)  
3\.  [Examples](#examples)  
3.1\.  [Files](#files)  
3.2\.  [CRUD](#crud)  
4\.  [Config](#config)  
4.1\.  [Config YAML](#configyaml)  
4.2\.  [Config JSON](#configjson)  
5\.  [API](#api-1)  
5.1\.  [Rest Endpoint API](#restendpointapi)  
5.1.1\.  [Rest Endpoint API Description](#restendpointapidescription)  
5.1.2\.  [Rest Endpoint API Examples](#restendpointapiexamples)  
5.1.2.1\.  [GET `/_api/rest/endpoints`](#get`/_api/rest/endpoints`)  
5.1.2.2\.  [POST `/_api/rest/endpoints`](#post`/_api/rest/endpoints`)  
5.1.2.3\.  [GET `/_api/rest/endpoints/:id`](#get`/_api/rest/endpoints/:id`)  
5.1.2.4\.  [PUT `/_api/rest/endpoints/:id`](#put`/_api/rest/endpoints/:id`)  
5.2\.  [Rest Global API](#restglobalapi)  
5.2.1\.  [Rest Global API Description](#restglobalapidescription)  
5.2.2\.  [Rest Global API Examples](#restglobalapiexamples)  
5.2.2.1\.  [GET `/_api/rest/global`](#get`/_api/rest/global`)  
5.2.2.2\.  [POST `/_api/rest/global`](#post`/_api/rest/global`)  
5.3\.  [Files API](#filesapi)  
5.3.1\.  [Files API Description](#filesapidescription)  
5.3.2\.  [Files API Examples](#filesapiexamples)  
5.3.2.1\.  [GET `/_api/files`](#get`/_api/files`)  
5.3.2.2\.  [GET `/_api/files/:id`](#get`/_api/files/:id`)  
6\.  [Models](#models)  
6.1\.  [Rest Endpoint Model](#restendpointmodel)  
6.1.1\.  [Rest Endpoint Request Model](#restendpointrequestmodel)  
6.1.2\.  [Rest Endpoint Response Model](#restendpointresponsemodel)  
6.2\.  [File Model](#filemodel)  

<a name="prerequisites"></a>

## 1\. Prerequisites

<a name="install"></a>

### 1.1\. Install

- [Docker](https://docs.docker.com/install/)
- _Optional_ [go](https://golang.org/doc/install)

<a name="examplefiles"></a>

### 1.2\. Example files

> It is possible just to clone current repositiory with examples instead of creating files manually:
>
> ```bash
> git clone https://github.com/abproject/mock-server.git
> cd mock-server
> ```
>
> and then run docker commands with examples from repository (no path changes needed).

<a name="helloworld"></a>

## 2\. Hello World

<a name="yaml"></a>

### 2.1\. YAML

**Create file `examples/hello/config.yaml` with content:**

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

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/hello/config.yaml
```

**Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:**

```bash
curl -v http://localhost:4242/hello

### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/html
...
Hello, World!
```

<a name="json"></a>

### 2.2\. JSON

**Create file `examples/hello/config.json` with content:**

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

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/hello/config.json
```

**Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:**

```bash
curl -v http://localhost:4242/hello

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

**Run in terminal:**

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:**

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

e.g., with `curl` _(copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoints \
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

**Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:**

```bash
curl -v http://localhost:4242/hello

### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/html
...
Hello, World!
```

<a name="examples"></a>

## 3\. Examples

<a name="files"></a>

### 3.1\. [Files](docs/FILES_EXAMPLE.md)

<a name="crud"></a>

### 3.2\. [CRUD](docs/CRUD_EXAMPLE.md)

<a name="config"></a>

## 4\. Config

<a name="configyaml"></a>

### 4.1\. Config YAML

```yaml
rest:
  global:
    request:
      method: method name
      path: url path
    response:
      body: response body
      bodyFile: response body as file
      status: http status
      headers: # Map
        header-key: header-value
  endpoints: # List of Endpoint Models
    - request:
        method: method name
        path: url path
      response:
        body: response body
        bodyFile: response body as file
        status: http status
        headers: # Map
          header-key: header-value

```

**Details of model description:**

| Path             | Model                                               |
| ---------------- | --------------------------------------------------- |
| `rest.global`    | [Endpoint Model](#restendpointrequestmodel)         |
| `rest.endpoints` | List of [Endpoint Model](#restendpointrequestmodel) |

<a name="configjson"></a>

### 4.2\. Config JSON

```json
{
  "rest": {
    "global": {
      "request": {
        "method": "method name",
        "path": "url path"
      },
      "response": {
        "body": "response body",
        "bodyFile": "response body as file",
        "status": "http status",
        "headers": {
          "header-key": "header-value"
        }
      }
    },
    "endpoints": [
      {
        "request": {
          "method": "method name",
          "path": "url path"
        },
        "response": {
          "body": "response body",
          "bodyFile": "response body as file",
          "status": "http status",
          "headers": {
            "header-key": "header-value"
          }
        }
      }
    ]
  }
}

```

**Details of model description:**

| Path             | Model                                               |
| ---------------- | --------------------------------------------------- |
| `rest.global`    | [Endpoint Model](#restendpointrequestmodel)         |
| `rest.endpoints` | List of [Endpoint Model](#restendpointrequestmodel) |

<a name="api-1"></a>

## 5\. API

<a name="restendpointapi"></a>

### 5.1\. Rest Endpoint API

URL: `/_api/rest/endpoints`

<a name="restendpointapidescription"></a>

#### 5.1.1\. Rest Endpoint API Description

| Path                       | Method   | Description                                                         |                Request Body                 |                    Response Body                    | Success Status | Failed Status |
| -------------------------- | -------- | ------------------------------------------------------------------- | :-----------------------------------------: | :-------------------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoints`     | `GET`    | Returns the list of all endpoints configurations                    |                      -                      | List of [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/endpoints`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **201**     |       -       |
| `/_api/rest/endpoints`     | `DELETE` | Deletes all endpoints configuration                                 |                      -                      |                          -                          |    **204**     |       -       |
| `/_api/rest/endpoints/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                      -                      |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoints/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoints/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                      -                      |                          -                          |    **204**     |    **404**    |

<a name="restendpointapiexamples"></a>

#### 5.1.2\. Rest Endpoint API Examples

<a name="get`/_api/rest/endpoints`"></a>

##### 5.1.2.1\. GET `/_api/rest/endpoints`

**Response**
```json
[
  {
    "id": ":id",
    "request": {
      "method": "GET",
      "path": "hello",
      "pathReg": "",
      "headers": null
    },
    "response": {
      "status": 200,
      "body": "Hello, World!",
      "bodyFile": "",
      "headers": {
        "Content-Type": "text/html"
      }
    }
  }
]

```

<a name="post`/_api/rest/endpoints`"></a>

##### 5.1.2.2\. POST `/_api/rest/endpoints`

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
      "Content-Type": "text/html"
    }
  }
}

```

**Response:**

```json
{
  "id": ":id",
  "request": {
    "method": "GET",
    "path": "hello",
    "pathReg": "",
    "headers": null
  },
  "response": {
    "status": 200,
    "body": "Hello, World!",
    "bodyFile": "",
    "headers": {
      "Content-Type": "text/html"
    }
  }
}

```

<a name="get`/_api/rest/endpoints/:id`"></a>

##### 5.1.2.3\. GET `/_api/rest/endpoints/:id`

**Response**
```json
{
  "id": ":id",
  "request": {
    "method": "GET",
    "path": "hello",
    "pathReg": "",
    "headers": null
  },
  "response": {
    "status": 200,
    "body": "Hello, World!",
    "bodyFile": "",
    "headers": {
      "Content-Type": "text/html"
    }
  }
}

```

<a name="put`/_api/rest/endpoints/:id`"></a>

##### 5.1.2.4\. PUT `/_api/rest/endpoints/:id`

**Request:**

```json
{
  "request": {
    "method": "GET",
    "path": "hello-new"
  },
  "response": {
    "status": 200,
    "body": "Hello, New World!",
    "headers": {
      "Content-Type": "text/html"
    }
  }
}

```

**Response:**

```json
{
  "request": {
    "method": "GET",
    "path": "hello-new",
    "pathReg": "",
    "headers": null
  },
  "response": {
    "status": 200,
    "body": "Hello, New World!",
    "bodyFile": "",
    "headers": {
      "Content-Type": "text/html"
    }
  }
}

```

<a name="restglobalapi"></a>

### 5.2\. Rest Global API

URL: `/_api/rest/global`

<a name="restglobalapidescription"></a>

#### 5.2.1\. Rest Global API Description

| Path                | Method   | Description                            |                Request Body                 |                Response Body                | Success Status | Failed Status |
| ------------------- | -------- | -------------------------------------- | :-----------------------------------------: | :-----------------------------------------: | :------------: | :-----------: |
| `/_api/rest/global` | `GET`    | Returns global endpoint configurations |                      -                      | [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/global` | `POST`   | Creates new global endpoint entity     | [Endpoint Model](#restendpointrequestmodel) | [Endpoint Model](#restendpointrequestmodel) |    **201**     |       -       |
| `/_api/rest/global` | `DELETE` | Deletes global endpoint configuration  |                      -                      |                      -                      |    **204**     |       -       |

<a name="restglobalapiexamples"></a>

#### 5.2.2\. Rest Global API Examples

<a name="get`/_api/rest/global`"></a>

##### 5.2.2.1\. GET `/_api/rest/global`

**Response**
```json
{
  "id": "",
  "request": {
    "method": "",
    "path": "",
    "pathReg": "",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "response": {
    "status": 0,
    "body": "",
    "bodyFile": "",
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

<a name="post`/_api/rest/global`"></a>

##### 5.2.2.2\. POST `/_api/rest/global`

**Request:**

```json
{
  "request": {
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "response": {
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

**Response:**

```json
{
  "id": "",
  "request": {
    "method": "",
    "path": "",
    "pathReg": "",
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "response": {
    "status": 0,
    "body": "",
    "bodyFile": "",
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

<a name="filesapi"></a>

### 5.3\. Files API

URL: `/_api/files`

<a name="filesapidescription"></a>

#### 5.3.1\. Files API Description

| Path              | Method   | Description                                                     |                                                                        Request Body                                                                         |          Response Body           | Success Status | Failed Status |
| ----------------- | -------- | --------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------: | :------------------------------: | :------------: | :-----------: |
| `/_api/files`     | `GET`    | Returns the list of all files configurations                    |                                                                              -                                                                              | List of [File Model](#filemodel) |    **200**     |       -       |
| `/_api/files`     | `POST`   | Creates new file entity                                         | <div style="text-align: left">**Body** as `form-data`:<br>`file: <file content>`<br>**Headers**:<br>`Content-Type: application/x-www-form-urlencoded`</div> |     [File Model](#filemodel)     |    **201**     |       -       |
| `/_api/files`     | `DELETE` | Deletes all files configuration                                 |                                                                              -                                                                              |                -                 |    **204**     |       -       |
| `/_api/files/:id` | `GET`    | Returns file by `id` or error if not found                      |                                                                              -                                                                              |     [File Model](#filemodel)     |    **200**     |    **404**    |
| `/_api/files/:id` | `PUT`    | Sets new file configuration by `id`, returns error if not found | <div style="text-align: left">**Body** as `form-data`:<br>`file: <file content>`<br>**Headers**:<br>`Content-Type: application/x-www-form-urlencoded`</div> |     [File Model](#filemodel)     |    **200**     |    **404**    |
| `/_api/files/:id` | `DELETE` | Deletes file configuration by `id`, returns error if not found  |                                                                              -                                                                              |                -                 |    **204**     |    **404**    |

<a name="filesapiexamples"></a>

#### 5.3.2\. Files API Examples

<a name="get`/_api/files`"></a>

##### 5.3.2.1\. GET `/_api/files`

**Response**
```json
[
  {
    "id": ":id",
    "name": "api-post-request.txt",
    "length": 9
  }
]

```

<a name="get`/_api/files/:id`"></a>

##### 5.3.2.2\. GET `/_api/files/:id`

**Response**
```json
{
  "id": ":id",
  "name": "api-post-request.txt",
  "length": 9
}

```

<a name="models"></a>

## 6\. Models

<a name="restendpointmodel"></a>

### 6.1\. Rest Endpoint Model

| Field Name | Type                                                       | Description                                      |
| ---------- | ---------------------------------------------------------- | ------------------------------------------------ |
| `id`       | `string`                                                   | Unique Endpoint ID. **Generates by mock-server** |
| `request`  | [Rest Endpoint Request Model](#restendpointrequestmodel)   | Request configuration model                      |
| `response` | [Rest Endpoint Response Model](#restendpointresponsemodel) | Response configuration model                     |

<a name="restendpointrequestmodel"></a>

#### 6.1.1\. Rest Endpoint Request Model

| Field Name | Type                  | Description                                                                                                                              |
| ---------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `method`   | `string`              | Method name, e.g., `GET`, `POST`, `DELETE`, `PUT`. <br>Empty string means all type of requests. <br>**Default:** empty string            |
| `path`     | `string`              | Request Endpoint path, e.g., `/my-path`. <br>**Default:** empty string                                                                   |
| `pathReg`  | `string`              | Request Endpoint path as regular expression. If `pathReg` value is not empty then `path` value is ignored. <br>**Default:** empty string |
| `headers`  | `map<string, string>` | Request Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                              |

<a name="restendpointresponsemodel"></a>

#### 6.1.2\. Rest Endpoint Response Model

| Field Name | Type                  | Description                                                                                                                            |
| ---------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `body`     | `any`                 | Response body could be type including JSON objects. <br>**Default:** empty string                                                      |
| `bodyFile` | `string`              | Response body is the content of the file. If `bodyFile` value is not empty then `body` value is ignored. <br>**Default:** empty string |
| `status`   | `integer`             | Response HTTP status code. <br>**Default:** `0`                                                                                        |
| `headers`  | `map<string, string>` | Response Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                           |

<a name="filemodel"></a>

### 6.2\. File Model

| Field Name | Type     | Description                                                                                |
| ---------- | -------- | ------------------------------------------------------------------------------------------ |
| `id`       | `string` | Unique Endpoint ID or file name in case of config parsing <br>**Generates by mock-server** |
| `name`     | `string` | Provided file name                                                                         |
| `length`   | `number` | Size of file in bytes                                                                      |
