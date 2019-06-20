# mock-server
> **mock-server** is a tiny Web Server (Docker image <10 MB) that allows easily to emulate the RESTful and WebSocket services.
> 
> It could be useful for:
> * FrontEnd Developers: fast mocking the BackEnd endpoints during implementation of the UI prototypes
> * BackEnd Developers: mocking services and applications during integration or system tests 
---

## Table of Contents
* [Run](#run)
* [example](#example)
    * [Hello World](#hello-world)
    * [Files](#files)
    * [CRUD](#crud)
    * [Entity](#entity)
    * [WebSocket](#websocket)
* [API](#api)

---
## Run

`docker run -p 8000:8000 -v ${PWD}/example:/example abezpalov/mock-server:latest -file=/example/crud/config.yml`

---

## example
All example are based on using Docker image, but it is possible to use compiled binary executable file instead.

> It is recommended for probing  example below to create a folder `mkdir example` where all example will be stored. Or just clone this repository with provided example.

### Hello World
Introductory description of simple `Hello World` endpoint.
* Create folder `mkdir example/hello`

* Do steps for `YAML`, `JSON` or `API` defined below in this block.

* Test via `GET` request:
  * `curl http://localhost:4242/hello`
  * or open in browser [http://localhost:4242/hello](http://localhost:4242/hello)

* The response with status `200` is:
    ```
    Hello, World!
    ```

#### YAML
Create file `example/hello/config.yml` with the content:
```yaml
rest:
  endpoints:
    - request:
        method: GET
        path: hello
      response:
        status: 200
        body: Hello, World!
```
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/hello/config.yml

```

#### JSON
Create file `example/hello/config.json` with the content:
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
          "body": "Hello, World!"
        }
      }
    ]
  }
}
```
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/hello/config.json

```

#### API
Run:
```console
docker run -p 4242:8000 abezpalov/mock-server
```
Send `POST` request:
```console
curl -v -X POST http://localhost:4242/_api/rest \
-H "Content-Type: application/json" \
-d @- << EOF

{
    "request": {
      "method": "GET",
      "path": "hello"
    },
    "response": {
      "status": 200,
      "body": "Hello, World!"
    }
}
EOF
```

#### Checking the configuration via API
* `curl http://localhost:4242/_api/rest`
* Response:
    ```json
    [
      {
        "id": "${unique-id}",
        "request": {
          "method": "GET",
          "path": "hello",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "Hello, World!",
          "file": "",
          "status": 200,
          "headers": {}
        }
      }
    ]
    ```

### Files
> If server must return a file or the response body JSON is stored in a separate file it is possible to use `file` configuration.

In this example the server will return `file.txt` via [http://localhost:4242/file](http://localhost:4242/file) and return a JSON `{ "message": "Hello, World!" }` via [http://localhost:4242/hello](http://localhost:4242/hello) 

* Create folder `mkdir example/files`

* Create file `example/files/file.txt` with the content:
    ```
    Hello from file!
    ```
* Create file `example/files/hello.json` with the content:
    ```json
    {
      "message": "Hello, World!"
    }
    ```
* Do steps for `YAML` or `JSON` defined below in this block.

* Test via `GET` request:
  * file.txt
     * `curl http://localhost:4242/file`
     * or open in browser [http://localhost:4242/file](http://localhost:4242/file)
     * the response with status `200` is a file containing:
        ```
        Hello from file!
        ``` 
  * Response body stored in `hello.json` file
     * `curl http://localhost:4242/hello`
     * or open in browser [http://localhost:4242/hello](http://localhost:4242/hello)
     * the response with status `200` is:
         ```json
         { "message": "Hello, World!" }
         ```
         
#### YAML
Create file `example/files/config.yml` with the content:
```yaml
rest:
  endpoints:
    - request:
        method: GET
        path: file
      response:
        file: example/files/file.txt
        status: 200
        headers:
          Content-Type: application/octet-stream
    - request:
        method: GET
        path: hello
      response:
        file: example/files/hello.json
        status: 200
        headers:
          Content-Type: application/json
```

Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/files/config.yml

```


#### JSON
Create file `example/files/config.json` with the content:
  ```json
  {
    "rest": {
      "endpoints": [
        {
          "request": {
            "method": "GET",
            "path": "file"
          },
          "response": {
            "file": "example/files/file.txt",
            "status": 200,
            "headers": {
              "Content-Type": "application/octet-stream"
            }
          }
        },
        {
          "request": {
            "method": "GET",
            "path": "hello"
          },
          "response": {
            "file": "example/files/hello.json",
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
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/hello/config.json

```

#### Checking the configuration via API
* `curl http://localhost:4242/_api/rest`
* Response:
  ```json
  [
    {
      "id": "${unique-id}",
      "request": {
        "method": "GET",
        "path": "file",
        "pathReg": "",
        "headers": {}
      },
      "response": {
        "body": "",
        "file": "example/files/file.txt",
        "status": 200,
        "headers": {
            "Content-Type": "application/octet-stream"
        }
      }
    },
    {
      "id": "${unique-id}",
      "request": {
        "method": "GET",
        "path": "hello",
        "pathReg": "",
        "headers": {}
      },
      "response": {
        "body": "",
        "file": "example/files/hello.json",
        "status": 200,
        "headers": {
            "Content-Type": "application/json"
        }
      }
    }
  ]
  ```
### CRUD
### Entity
### WebSocket

---
## API

