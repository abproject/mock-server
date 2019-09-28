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
3\.  [Config](#config)  
3.1\.  [Config YAML](#configyaml)  
3.2\.  [Config JSON](#configjson)  
4\.  [API](#api-1)  
4.1\.  [Rest Endpoint API](#restendpointapi)  
4.1.1\.  [Rest Endpoint API Description](#restendpointapidescription)  
4.1.2\.  [Rest Endpoint API Examples](#restendpointapiexamples)  
4.1.2.1\.  [GET `/_api/rest/endpoints`](#get`/_api/rest/endpoints`)  
4.1.2.2\.  [POST `/_api/rest/endpoints`](#post`/_api/rest/endpoints`)  
4.1.2.3\.  [GET `/_api/rest/endpoints/:id`](#get`/_api/rest/endpoints/:id`)  
4.1.2.4\.  [PUT `/_api/rest/endpoints/:id`](#put`/_api/rest/endpoints/:id`)  
4.2\.  [Rest Global API](#restglobalapi)  
4.2.1\.  [Rest Global API Description](#restglobalapidescription)  
4.2.2\.  [Rest Global API Examples](#restglobalapiexamples)  
4.2.2.1\.  [GET `/_api/rest/global`](#get`/_api/rest/global`)  
4.2.2.2\.  [POST `/_api/rest/global`](#post`/_api/rest/global`)  
4.3\.  [Files API](#filesapi)  
4.3.1\.  [Files API Description](#filesapidescription)  
4.3.2\.  [Files API Examples](#filesapiexamples)  
4.3.2.1\.  [GET `/_api/files`](#get`/_api/files`)  
4.3.2.2\.  [GET `/_api/files/:id`](#get`/_api/files/:id`)  
5\.  [Models](#models)  
5.1\.  [Rest Endpoint Model](#restendpointmodel)  
5.1.1\.  [Rest Endpoint Request Model](#restendpointrequestmodel)  
5.1.2\.  [Rest Endpoint Response Model](#restendpointresponsemodel)  
5.2\.  [File Model](#filemodel)  
6\.  [Examples](#examples)  
6.1\.  [Files](#files)  
6.2\.  [YAML](#yaml-1)  
6.3\.  [JSON](#json-1)  
6.4\.  [API](#api-2)  

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

Create file `examples/hello/config.yaml` with content:

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
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/hello/config.yaml
```

Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:

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

Create file `examples/hello/config.json` with content:

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
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/hello/config.json
```

Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:

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

Run in terminal:

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:

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

Check by opening in browser http://localhost:4242/hello or making `GET` request, e.g., with `curl`:

```bash
curl -v http://localhost:4242/hello

### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/html
...
Hello, World!
```

<a name="config"></a>

## 3\. Config

<a name="configyaml"></a>

### 3.1\. Config YAML

```yaml
rest:
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
| `rest.endpoints` | List of [Endpoint Model](#restendpointrequestmodel) |

<a name="configjson"></a>

### 3.2\. Config JSON

```json
{
  "rest": {
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
| `rest.endpoints` | List of [Endpoint Model](#restendpointrequestmodel) |

<a name="api-1"></a>

## 4\. API

<a name="restendpointapi"></a>

### 4.1\. Rest Endpoint API

URL: `/_api/rest/endpoints`

<a name="restendpointapidescription"></a>

#### 4.1.1\. Rest Endpoint API Description

| Path                       | Method   | Description                                                         |                Request Body                 |                    Response Body                    | Success Status | Failed Status |
| -------------------------- | -------- | ------------------------------------------------------------------- | :-----------------------------------------: | :-------------------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoints`     | `GET`    | Returns the list of all endpoints configurations                    |                      -                      | List of [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/endpoints`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **201**     |       -       |
| `/_api/rest/endpoints`     | `DELETE` | Deletes all endpoints configuration                                 |                      -                      |                          -                          |    **204**     |       -       |
| `/_api/rest/endpoints/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                      -                      |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoints/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoints/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                      -                      |                          -                          |    **204**     |    **404**    |

<a name="restendpointapiexamples"></a>

#### 4.1.2\. Rest Endpoint API Examples

<a name="get`/_api/rest/endpoints`"></a>

##### 4.1.2.1\. GET `/_api/rest/endpoints`

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

##### 4.1.2.2\. POST `/_api/rest/endpoints`

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

##### 4.1.2.3\. GET `/_api/rest/endpoints/:id`

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

##### 4.1.2.4\. PUT `/_api/rest/endpoints/:id`

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

### 4.2\. Rest Global API

URL: `/_api/rest/global`

<a name="restglobalapidescription"></a>

#### 4.2.1\. Rest Global API Description

| Path                | Method   | Description                            |                Request Body                 |                Response Body                | Success Status | Failed Status |
| ------------------- | -------- | -------------------------------------- | :-----------------------------------------: | :-----------------------------------------: | :------------: | :-----------: |
| `/_api/rest/global` | `GET`    | Returns global endpoint configurations |                      -                      | [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/global` | `POST`   | Creates new global endpoint entity     | [Endpoint Model](#restendpointrequestmodel) | [Endpoint Model](#restendpointrequestmodel) |    **201**     |       -       |
| `/_api/rest/global` | `DELETE` | Deletes global endpoint configuration  |                      -                      |                      -                      |    **204**     |       -       |

<a name="restglobalapiexamples"></a>

#### 4.2.2\. Rest Global API Examples

<a name="get`/_api/rest/global`"></a>

##### 4.2.2.1\. GET `/_api/rest/global`

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
    "status": 200,
    "body": "",
    "bodyFile": "",
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

<a name="post`/_api/rest/global`"></a>

##### 4.2.2.2\. POST `/_api/rest/global`

**Request:**

```json
{
  "request": {
    "headers": {
      "Content-Type": "application/json"
    }
  },
  "response": {
    "status": 200,
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
    "status": 200,
    "body": "",
    "bodyFile": "",
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

<a name="filesapi"></a>

### 4.3\. Files API

URL: `/_api/files`

<a name="filesapidescription"></a>

#### 4.3.1\. Files API Description

| Path              | Method   | Description                                                     |                                                                        Request Body                                                                         |          Response Body           | Success Status | Failed Status |
| ----------------- | -------- | --------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------: | :------------------------------: | :------------: | :-----------: |
| `/_api/files`     | `GET`    | Returns the list of all files configurations                    |                                                                              -                                                                              | List of [File Model](#filemodel) |    **200**     |       -       |
| `/_api/files`     | `POST`   | Creates new file entity                                         | <div style="text-align: left">**Body** as `form-data`:<br>`file: <file content>`<br>**Headers**:<br>`Content-Type: application/x-www-form-urlencoded`</div> |     [File Model](#filemodel)     |    **201**     |       -       |
| `/_api/files`     | `DELETE` | Deletes all files configuration                                 |                                                                              -                                                                              |                -                 |    **204**     |       -       |
| `/_api/files/:id` | `GET`    | Returns file by `id` or error if not found                      |                                                                              -                                                                              |     [File Model](#filemodel)     |    **200**     |    **404**    |
| `/_api/files/:id` | `PUT`    | Sets new file configuration by `id`, returns error if not found | <div style="text-align: left">**Body** as `form-data`:<br>`file: <file content>`<br>**Headers**:<br>`Content-Type: application/x-www-form-urlencoded`</div> |     [File Model](#filemodel)     |    **200**     |    **404**    |
| `/_api/files/:id` | `DELETE` | Deletes file configuration by `id`, returns error if not found  |                                                                              -                                                                              |                -                 |    **204**     |    **404**    |

<a name="filesapiexamples"></a>

#### 4.3.2\. Files API Examples

<a name="get`/_api/files`"></a>

##### 4.3.2.1\. GET `/_api/files`

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

##### 4.3.2.2\. GET `/_api/files/:id`

**Response**
```json
{
  "id": ":id",
  "name": "api-post-request.txt",
  "length": 9
}

```

<a name="models"></a>

## 5\. Models

<a name="restendpointmodel"></a>

### 5.1\. Rest Endpoint Model

| Field Name | Type                                                       | Description                                      |
| ---------- | ---------------------------------------------------------- | ------------------------------------------------ |
| `id`       | `string`                                                   | Unique Endpoint ID. **Generates by mock-server** |
| `request`  | [Rest Endpoint Request Model](#restendpointrequestmodel)   | Request configuration model                      |
| `response` | [Rest Endpoint Response Model](#restendpointresponsemodel) | Response configuration model                     |

<a name="restendpointrequestmodel"></a>

#### 5.1.1\. Rest Endpoint Request Model

| Field Name | Type                  | Description                                                                                                                              |
| ---------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `method`   | `string`              | Method name, e.g., `GET`, `POST`, `DELETE`, `PUT`. <br>Empty string means all type of requests. <br>**Default:** empty string            |
| `path`     | `string`              | Request Endpoint path, e.g., `/my-path`. <br>**Default:** empty string                                                                   |
| `pathReg`  | `string`              | Request Endpoint path as regular expression. If `pathReg` value is not empty then `path` value is ignored. <br>**Default:** empty string |
| `headers`  | `map<string, string>` | Request Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                              |

<a name="restendpointresponsemodel"></a>

#### 5.1.2\. Rest Endpoint Response Model

| Field Name | Type                  | Description                                                                                                                            |
| ---------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `body`     | `any`                 | Response body could be type including JSON objects. <br>**Default:** empty string                                                      |
| `bodyFile` | `string`              | Response body is the content of the file. If `bodyFile` value is not empty then `body` value is ignored. <br>**Default:** empty string |
| `status`   | `integer`             | Response HTTP status code. <br>**Default:** `0`                                                                                        |
| `headers`  | `map<string, string>` | Response Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                           |

<a name="filemodel"></a>

### 5.2\. File Model

| Field Name | Type     | Description                                                                                |
| ---------- | -------- | ------------------------------------------------------------------------------------------ |
| `id`       | `string` | Unique Endpoint ID or file name in case of config parsing <br>**Generates by mock-server** |
| `name`     | `string` | Provided file name                                                                         |
| `length`   | `number` | Size of file in bytes                                                                      |

<a name="examples"></a>

## 6\. Examples

<a name="files"></a>

### 6.1\. Files

<a name="yaml-1"></a>

### 6.2\. YAML

Create file `examples/files/hello.txt` with content:

```txt
Hello from file!
```

Create file `examples/files/hello.json` with content:

```json
{
  "message": "Hello, World!"
}

```

Create file `examples/files/config.yaml` with content:

```yaml
rest:
  endpoints:
    - request:
        method: GET
        path: hello-txt
      response:
        bodyFile: examples/files/hello.txt
        status: 200
        headers:
          Content-Type: text/plain
    - request:
        method: GET
        path: hello-json
      response:
        bodyFile: examples/files/hello.json
        status: 200
        headers:
          Content-Type: application/json

```

Final structure:

```bash
 <your-path>/examples/files/config.yaml
 <your-path>/examples/files/hello.json
 <your-path>/examples/files/hello.txt
```

Run in terminal:

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/files/config.yaml
```

Check `hello.txt`:

```bash
curl -v http://localhost:4242/hello-txt

### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

Check `hello.json`:

```bash
curl -v http://localhost:4242/hello-json

### Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```

<a name="json-1"></a>

### 6.3\. JSON

Create file `examples/files/hello.txt` with content:

```txt
Hello from file!
```

Create file `examples/files/hello.json` with content:

```json
{
  "message": "Hello, World!"
}

```

Create file `examples/files/config.json` with content:

```yaml
{
  "rest": {
    "endpoints": [
      {
        "request": {
          "method": "GET",
          "path": "hello-txt"
        },
        "response": {
          "bodyFile": "examples/files/hello.txt",
          "status": 200,
          "headers": {
            "Content-Type": "text/plain"
          }
        }
      },
      {
        "request": {
          "method": "GET",
          "path": "hello-json"
        },
        "response": {
          "bodyFile": "examples/files/hello.json",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      }
    ]
  }
}

```

Final structure:

```bash
 <your-path>/examples/files/config.json
 <your-path>/examples/files/hello.json
 <your-path>/examples/files/hello.txt
```

Run in terminal:

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/files/config.json
```

Check `hello.txt`:

```bash
curl -v http://localhost:4242/hello-txt

### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

Check `hello.json`:

```bash
curl -v http://localhost:4242/hello-json

### Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```

<a name="api-2"></a>

### 6.4\. API

Run in terminal:

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

Create file `examples/files/hello.txt` with content:

```txt
Hello from file!
```

Create file `examples/files/hello.json` with content:

```json
{
  "message": "Hello, World!"
}

```

Final structure:

```bash
 <your-path>/examples/files/hello.json
 <your-path>/examples/files/hello.txt
```

**hello.txt**

Send `POST` request to URL `http://localhost:4242/_api/files` with:

- body as `form-data`
  - `file: hello.txt`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/files/hello.txt' http://localhost:4242/_api/files

### Response, e.g.:
{"id":"6694d2c422ac4208a0072939487f6999","name":"hello.txt","length":16}
```

Copy `id` from response, e.g., `6694d2c422ac4208a0072939487f6999`.

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:

```json
{
  "request": {
    "method": "GET",
    "path": "hello-txt"
  },
  "response": {
    "status": 200,
    "bodyFile": "<changed-to-id>",
    "headers": {
      "Content-Type": "text/plain"
    }
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoints \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "GET",
    "path": "hello-txt"
  },
  "response": {
    "status": 200,
    "bodyFile": "<changed-to-id>",
    "headers": {
      "Content-Type": "text/plain"
    }
  }
}

```

```
EOF
```

**hello.json**

Send `POST` request to URL `http://localhost:4242/_api/files` with:

- body as `form-data`
  - `file: hello.json`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/files/hello.json' http://localhost:4242/_api/files

### Response, e.g.:
{"id":"9566c74d10034c4dbbbb0407d1e2c649","name":"hello.json","length":16}
```

Copy `id` from response, e.g., `9566c74d10034c4dbbbb0407d1e2c649`.

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:

```json
{
  "request": {
    "method": "GET",
    "path": "hello-json"
  },
  "response": {
    "status": 200,
    "bodyFile": "<changed-to-id>",
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoints \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "GET",
    "path": "hello-json"
  },
  "response": {
    "status": 200,
    "bodyFile": "<changed-to-id>",
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

```
EOF
```

Check `hello.txt`:

```bash
curl -v http://localhost:4242/hello-txt

### Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

Check `hello.json`:

```bash
curl -v http://localhost:4242/hello-json

### Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```
