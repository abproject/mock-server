# mock-server

> **mock-server** is a tiny Web Server (Docker image <10 MB) that allows easily to emulate the RESTful and WebSocket services.
>
> It could be useful for:
>
> - FrontEnd Developers: fast mocking the BackEnd endpoints during implementation of the UI prototypes
> - BackEnd Developers: mocking services and applications during integration or system tests

---

## Table of Contents

- [Run](#run)
- [File Configuration](#file-configuration)
- [API](#api)

## Run

`docker run -p 8000:8000 -v ${PWD}/example:/example abezpalov/mock-server:latest -file=/example/crud/config.yaml`

---

There are more [examples](example/README.md):

- [Files](example/README.md#files)
- [CRUD](example/README.md#crud)
- [Entity](example/README.md#entity)
- [WebSocket](example/README.md#websocket)

### Hello World

Introductory description of simple `Hello World` endpoint.

- Create folder `mkdir example/hello`

- Do steps for `YAML`, `JSON` or `API` defined below in this block.

- Test via `GET` request:

  - `curl http://localhost:4242/hello`
  - or open in browser [http://localhost:4242/hello](http://localhost:4242/hello)

- The response with status `200` is:
  ```
  Hello, World!
  ```

#### YAML

Create file `example/hello/config.yaml` with the content:

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
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/hello/config.yaml

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

- `curl http://localhost:4242/_api/rest`
- Response:
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

---

## File Configuration

---

## API
