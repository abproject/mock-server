# mock-server
> **mock-server** is a tiny Web Server (Docker image <10 MB) that allows easily to emulate the REST and WebSocket services.
> It could be useful for:
> * FrontEnd Developers: for fast mocking the BackEnd endpoints during implementation of the UI prototypes
> * BackEnd Developers: for mocking applications in environment during integration or system tests 
---

## Table of Contents
* [Run](#run)
* [Examples](#examples)
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

## Examples
All examples are based on using Docker image, but it is possible to use compiled binary executable file instead.

### Hello World
Introductory description of simple `Hello World` application.
* Do steps for `YAML`, `JSON` or `API` defined below in this block.

* Test via `GET` request:
  * `curl http://localhost:8123/hello`
  * or open in browser `localhost:8123/hello`

The response with status `200` is:
```
Hello, World!
```

#### YAML
Create file `config.yml` with the content:
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
docker run -p 8123:8000 -v ${PWD}/config.yml:/config.yml abezpalov/mock-server -file=config.yml

```

#### JSON
Create file `config.json` with the content:
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
docker run -p 8123:8000 -v ${PWD}/config.json:/config.json abezpalov/mock-server -file=config.json

```

#### API
Run:
```console
docker run -p 8123:8000 abezpalov/mock-server
```
Send `POST` request:
```console
curl -v -X POST http://localhost:8123/_api/rest \
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
* `curl http://localhost:8123/_api/rest`
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
### CRUD
### Entity
### WebSocket

---
## API

