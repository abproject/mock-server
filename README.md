# mock-server
> **mock-server** is a tiny Web Server (Docker image <10 MB) that allows easily to emulate the RESTful and WebSocket services.
> 
> It could be useful for:
> * FrontEnd Developers: fast mocking the BackEnd endpoints during implementation of the UI prototypes
> * BackEnd Developers: mocking services and applications during integration or system tests 
---

## Table of Contents
* [Run](#run)
* [Examples](#examples)
    * [Hello World](#hello-world)
    * [Files](#files)
    * [CRUD](#crud)
    * [Entity](#entity)
    * [WebSocket](#websocket)
* [File Configuration](#file-configuration)
* [API](#api)

---
## Run

`docker run -p 8000:8000 -v ${PWD}/example:/example abezpalov/mock-server:latest -file=/example/crud/config.yml`

---

## Examples
All example are based on using Docker image, but it is possible to use compiled binary executable file instead.

> It is recommended for probing examples below to create a folder `mkdir example` where all examples will be stored. Or just clone this repository with provided examples.

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
Simple example of basic [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) over `planets` entity.

* Create folder `mkdir example/crud`

* Do steps for `YAML`, `JSON` or `API` defined below in this block.

* Test:
  * GET (all): `curl http://localhost:4242/planets` status `200`
      ```json
      [
        {
          "id": 1,
          "name": "Mercury",
          "type": "Terrestrial planet",
          "period": 0.24,
          "atmosphere": []
        },
        {
          "id": 2,
          "name": "Venus",
          "type": "Terrestrial planet",
          "period": 0.62,
          "atmosphere": ["CO2", "N2"]
        },
        {
          "id": 3,
          "name": "Earth",
          "type": "Terrestrial planet",
          "period": 1,
          "atmosphere": ["N2", "O2", "Ar"]
        },
        {
          "id": 4,
          "name": "Mars",
          "type": "Terrestrial planet",
          "period": 1.88,
          "atmosphere": ["CO2", "N2", "Ar"]
        }
      ]
    ```
  * GET (id): `curl http://localhost:4242/planets/3` status `200`
    ```json
    {
      "id": 3,
      "name": "Earth",
      "type": "Terrestrial planet",
      "period": 1,
      "atmosphere": ["N2", "O2", "Ar"]
    }
    ```
  * POST: `curl -X POST http://localhost:4242/planets` status `201`
    ```json
      {
        "id": 3,
        "name": "Earth",
        "type": "Terrestrial planet",
        "period": 1,
        "atmosphere": ["N2", "O2", "Ar"]
      }
      ```
  * PUT: `curl -X PUT http://localhost:4242/planets/3` status `200`
    ```json
      {
        "id": 3,
        "name": "Earth",
        "type": "Terrestrial planet",
        "period": 1,
        "atmosphere": ["N2", "O2", "Ar"]
      }
      ```  
  * DELETE: `curl -X DELETE http://localhost:4242/planets/3` status `200`
> **In this example for any `id` there is always the same response**

#### YAML
Create file `example/crud/config.yml` with the content:
```yaml
rest:
  global:
    response:
      status: 200
      headers:
        Content-Type: application/json
  endpoints:
    - request:
        method: GET
        path: planets
      response:
        body: >
          [
            {
              "id": 1,
              "name": "Mercury",
              "type": "Terrestrial planet",
              "period": 0.24,
              "atmosphere": []
            },
            {
              "id": 2,
              "name": "Venus",
              "type": "Terrestrial planet",
              "period": 0.62,
              "atmosphere": ["CO2", "N2"]
            },
            {
              "id": 3,
              "name": "Earth",
              "type": "Terrestrial planet",
              "period": 1,
              "atmosphere": ["N2", "O2", "Ar"]
            },
            {
              "id": 4,
              "name": "Mars",
              "type": "Terrestrial planet",
              "period": 1.88,
              "atmosphere": ["CO2", "N2", "Ar"]
            }
          ]
    - request:
        method: GET
        path: planets/:id
      response:
        body: '{"id":3,"name":"Earth","type":"Terrestrial planet","period":1,"atmosphere":["N2","O2","Ar"]}'
    - request:
        method: POST
        path: planets
      response:
        status: 201
        body: "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
    - request:
        method: PUT
        path: planets/:id
      response:
        body: '{"id":3,"name":"Earth","type":"Terrestrial planet","period":1,"atmosphere":["N2","O2","Ar"]}'
    - request:
        method: DELETE
        path: planets/:id
```
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/crud/config.yml

```

#### JSON
Create file `example/crud/config.json` with the content:
```json
{
  "rest": {
    "global": {
      "response": {
        "status": 200,
        "headers": {
          "Content-Type": "application/json"
        }
      }
    },
    "endpoints": [
      {
        "request": {
          "method": "GET",
          "path": "planets"
        },
        "response": {
          "body": "[{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]},{\"id\":2,\"name\":\"Venus\",\"type\":\"Terrestrial planet\",\"period\":0.62,\"atmosphere\":[\"CO2\",\"N2\"]},{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]},{\"id\":4,\"name\":\"Mars\",\"type\":\"Terrestrial planet\",\"period\":1.88,\"atmosphere\":[\"CO2\",\"N2\",\"Ar\"]}]"
        }
      },
      {
        "request": {
          "method": "GET",
          "path": "planets/:id"
        },
        "response": {
          "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
        }
      },
      {
        "request": {
          "method": "POST",
          "path": "planets"
        },
        "response": {
          "status": 201,
          "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
        }
      },
      {
        "request": {
          "method": "PUT",
          "path": "planets/:id"
        },
        "response": {
          "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
        }
      },
      {
        "request": {
          "method": "DELETE",
          "path": "planets/:id"
        }
      }
    ]
  }
}
```
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/crud/config.json

```

#### API
Run:
```console
docker run -p 4242:8000 abezpalov/mock-server
```
Send `POST` request for `GET (all)`:
```console
curl -v -X POST http://localhost:4242/_api/rest \
-H "Content-Type: application/json" \
-d @- << EOF

{
    "request": {
      "method": "GET",
      "path": "planets"
    },
    "response": {
      "status": 200,
      "headers": {
        "Content-Type": "application/json"
      },
      "body": "[{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]},{\"id\":2,\"name\":\"Venus\",\"type\":\"Terrestrial planet\",\"period\":0.62,\"atmosphere\":[\"CO2\",\"N2\"]},{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]},{\"id\":4,\"name\":\"Mars\",\"type\":\"Terrestrial planet\",\"period\":1.88,\"atmosphere\":[\"CO2\",\"N2\",\"Ar\"]}]"
    }
}
EOF
```

Send `POST` request for `GET (id)`:
```console
curl -v -X POST http://localhost:4242/_api/rest \
-H "Content-Type: application/json" \
-d @- << EOF

{
    "request": {
      "method": "GET",
      "path": "planets/:id"
    },
    "response": {
      "status": 200,
      "headers": {
        "Content-Type": "application/json"
      },
      "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
    }
}
EOF
```

Send `POST` request for `POST`:
```console
curl -v -X POST http://localhost:4242/_api/rest \
-H "Content-Type: application/json" \
-d @- << EOF

{
    "request": {
      "method": "POST",
      "path": "planets/:id"
    },
    "response": {
      "status": 201,
      "headers": {
        "Content-Type": "application/json"
      },
      "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
    }
}
EOF
```

Send `POST` request for `PUT`:
```console
curl -v -X POST http://localhost:4242/_api/rest \
-H "Content-Type: application/json" \
-d @- << EOF

{
    "request": {
      "method": "PUT",
      "path": "planets/:id"
    },
    "response": {
      "status": 200,
      "headers": {
        "Content-Type": "application/json"
      },
      "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}"
    }
}
EOF
```

Send `POST` request for `DELETE`:
```console
curl -v -X POST http://localhost:4242/_api/rest \
-H "Content-Type: application/json" \
-d @- << EOF

{
    "request": {
      "method": "DELETE",
      "path": "planets/:id"
    },
    "response": {
      "status": 200,
      "headers": {
        "Content-Type": "application/json"
      }
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
          "path": "planets",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "[{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]},{\"id\":2,\"name\":\"Venus\",\"type\":\"Terrestrial planet\",\"period\":0.62,\"atmosphere\":[\"CO2\",\"N2\"]},{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]},{\"id\":4,\"name\":\"Mars\",\"type\":\"Terrestrial planet\",\"period\":1.88,\"atmosphere\":[\"CO2\",\"N2\",\"Ar\"]}]",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
          "method": "GET",
          "path": "planets/:id",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
            "method": "POST",
            "path": "planets",
            "pathReg": "",
            "headers": {}
        },
        "response": {
            "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}",
            "file": "",
            "status": 201,
            "headers": {
              "Content-Type": "application/json"
            }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
            "method": "PUT",
            "path": "planets/:id",
            "pathReg": "",
            "headers": {}
        },
        "response": {
            "body": "{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]}",
            "file": "",
            "status": 200,
            "headers": {
              "Content-Type": "application/json"
            }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
          "method": "DELETE",
          "path": "planets/:id",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      }
    ]
    ```

### Entity
As CRUD example requires a lot of boilerplate there is the `entity` configuration that simplifies basic CRUD description.
> Entity provides an endpoint for each `id` in the list unlike `/:id` in the CRUD example  

* Create folder `mkdir example/entity`

* Do steps for `YAML`, `JSON` or `API` defined below in this block.

* Test:
  * GET (all): `curl http://localhost:4242/planets` status `200`
  * GET (id): 
    * `id = 1` `curl http://localhost:4242/planets/1` status `200`
      ```json
      {
        "id": 1,
        "name": "Mercury",
        "type": "Terrestrial planet",
        "period": 0.24,
        "atmosphere": []
      }
      ```
    * ...
    * `id = 3` `curl http://localhost:4242/planets/3` status `200`
      ```json
      {
        "id": 3,
        "name": "Earth",
        "type": "Terrestrial planet",
        "period": 1,
        "atmosphere": ["N2", "O2", "Ar"]
      }
      ```
    * ...
  * POST (id): 
    * `id = 1` `curl -X POST http://localhost:4242/planets` status `201`
      ```json
      {
        "id": 1,
        "name": "Mercury",
        "type": "Terrestrial planet",
        "period": 0.24,
        "atmosphere": []
      }
      ```
   * PUT (id): 
     * `id = 1` `curl -X PUT http://localhost:4242/planets/1` status `200`
       ```json
       {
         "id": 1,
         "name": "Mercury",
         "type": "Terrestrial planet",
         "period": 0.24,
         "atmosphere": []
       }
       ```
     * ...
     * `id = 3` `curl -X PUT http://localhost:4242/planets/3` status `200`
       ```json
       {
         "id": 3,
         "name": "Earth",
         "type": "Terrestrial planet",
         "period": 1,
         "atmosphere": ["N2", "O2", "Ar"]
       }
       ```
     * ...
   * DELETE (id): 
      * `id = 1` `curl -X DELETE http://localhost:4242/planets/1` status `200`
      * ...
      * `id = 3` `curl -X DELETE http://localhost:4242/planets/3` status `200`
      * ...

#### YAML
Create file `example/entity/config.yml` with the content:
```yaml
entity:
  path: planets
  data: >
    [
      {
        "id": 1,
        "name": "Mercury",
        "type": "Terrestrial planet",
        "period": 0.24,
        "atmosphere": []
      },
      {
        "id": 2,
        "name": "Venus",
        "type": "Terrestrial planet",
        "period": 0.62,
        "atmosphere": ["CO2", "N2"]
      },
      {
        "id": 3,
        "name": "Earth",
        "type": "Terrestrial planet",
        "period": 1,
        "atmosphere": ["N2", "O2", "Ar"]
      },
      {
        "id": 4,
        "name": "Mars",
        "type": "Terrestrial planet",
        "period": 1.88,
        "atmosphere": ["CO2", "N2", "Ar"]
      },
      {
        "id": 5,
        "name": "Jupiter",
        "type": "Gas giant",
        "period": 11.86,
        "atmosphere": ["H2", "He"]
      },
      {
        "id": 6,
        "name": "Saturn",
        "type": "Gas giant",
        "period": 29.46,
        "atmosphere": ["H2", "He"]
      },
      {
        "id": 7,
        "name": "Uranus",
        "type": "Ice giant",
        "period": 84.01,
        "atmosphere": ["H2", "He", "CH4"]
      },
      {
        "id": 8,
        "name": "Neptune",
        "type": "Ice giant",
        "period": 164.8,
        "atmosphere": ["H2", "He", "CH4"]
      }
    ]
```
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/entity/config.yml

```

> Alternative for best readability: it is possible to use files for data.
Create file `example/entity/config-short.yml` with the content:
```yaml
entity:
  path: planets
  file: example/entity/data.json
```

Create file `example/entity/data.json` with the content:
```json
[
  {
    "id": 1,
    "name": "Mercury",
    "type": "Terrestrial planet",
    "period": 0.24,
    "atmosphere": []
  },
  {
    "id": 2,
    "name": "Venus",
    "type": "Terrestrial planet",
    "period": 0.62,
    "atmosphere": [
      "CO2",
      "N2"
    ]
  },
  {
    "id": 3,
    "name": "Earth",
    "type": "Terrestrial planet",
    "period": 1,
    "atmosphere": [
      "N2",
      "O2",
      "Ar"
    ]
  },
  {
    "id": 4,
    "name": "Mars",
    "type": "Terrestrial planet",
    "period": 1.88,
    "atmosphere": [
      "CO2",
      "N2",
      "Ar"
    ]
  },
  {
    "id": 5,
    "name": "Jupiter",
    "type": "Gas giant",
    "period": 11.86,
    "atmosphere": [
      "H2",
      "He"
    ]
  },
  {
    "id": 6,
    "name": "Saturn",
    "type": "Gas giant",
    "period": 29.46,
    "atmosphere": [
      "H2",
      "He"
    ]
  },
  {
    "id": 7,
    "name": "Uranus",
    "type": "Ice giant",
    "period": 84.01,
    "atmosphere": [
      "H2",
      "He",
      "CH4"
    ]
  },
  {
    "id": 8,
    "name": "Neptune",
    "type": "Ice giant",
    "period": 164.8,
    "atmosphere": [
      "H2",
      "He",
      "CH4"
    ]
  }
]
```

Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/entity/config-short.yml

```


#### JSON
Create file `example/entity/config.json` with the content:
```json
{
  "entity": {
    "path": "planets",
    "data": "[{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]},{\"id\":2,\"name\":\"Venus\",\"type\":\"Terrestrial planet\",\"period\":0.62,\"atmosphere\":[\"CO2\",\"N2\"]},{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]},{\"id\":4,\"name\":\"Mars\",\"type\":\"Terrestrial planet\",\"period\":1.88,\"atmosphere\":[\"CO2\",\"N2\",\"Ar\"]},{\"id\":5,\"name\":\"Jupiter\",\"type\":\"Gas giant\",\"period\":11.86,\"atmosphere\":[\"H2\",\"He\"]},{\"id\":6,\"name\":\"Saturn\",\"type\":\"Gas giant\",\"period\":29.46,\"atmosphere\":[\"H2\",\"He\"]},{\"id\":7,\"name\":\"Uranus\",\"type\":\"Ice giant\",\"period\":84.01,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]},{\"id\":8,\"name\":\"Neptune\",\"type\":\"Ice giant\",\"period\":164.8,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]}]"
  }
}
```
Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/entity/config.json

```


> Alternative for best readability: it is possible to use files for data.
Create file `example/entity/config-short.json` with the content:
```json
{
  "entity": {
    "path": "planets",
    "file": "example/entity/data.json"
  }
}
```

Create file `example/entity/data.json` with the content:
```json
[
  {
    "id": 1,
    "name": "Mercury",
    "type": "Terrestrial planet",
    "period": 0.24,
    "atmosphere": []
  },
  {
    "id": 2,
    "name": "Venus",
    "type": "Terrestrial planet",
    "period": 0.62,
    "atmosphere": [
      "CO2",
      "N2"
    ]
  },
  {
    "id": 3,
    "name": "Earth",
    "type": "Terrestrial planet",
    "period": 1,
    "atmosphere": [
      "N2",
      "O2",
      "Ar"
    ]
  },
  {
    "id": 4,
    "name": "Mars",
    "type": "Terrestrial planet",
    "period": 1.88,
    "atmosphere": [
      "CO2",
      "N2",
      "Ar"
    ]
  },
  {
    "id": 5,
    "name": "Jupiter",
    "type": "Gas giant",
    "period": 11.86,
    "atmosphere": [
      "H2",
      "He"
    ]
  },
  {
    "id": 6,
    "name": "Saturn",
    "type": "Gas giant",
    "period": 29.46,
    "atmosphere": [
      "H2",
      "He"
    ]
  },
  {
    "id": 7,
    "name": "Uranus",
    "type": "Ice giant",
    "period": 84.01,
    "atmosphere": [
      "H2",
      "He",
      "CH4"
    ]
  },
  {
    "id": 8,
    "name": "Neptune",
    "type": "Ice giant",
    "period": 164.8,
    "atmosphere": [
      "H2",
      "He",
      "CH4"
    ]
  }
]
```

Run:
```console
docker run -p 4242:8000 -v ${PWD}/example:/example abezpalov/mock-server -file=example/entity/config-short.json

```

#### API
Run:
```console
docker run -p 4242:8000 abezpalov/mock-server
```
Send `POST` request:
```console
curl -v -X POST http://localhost:4242/_api/entity \
-H "Content-Type: application/json" \
-d @- << EOF

{
  "entity": {
    "path": "planets",
    "data": "[{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]},{\"id\":2,\"name\":\"Venus\",\"type\":\"Terrestrial planet\",\"period\":0.62,\"atmosphere\":[\"CO2\",\"N2\"]},{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]},{\"id\":4,\"name\":\"Mars\",\"type\":\"Terrestrial planet\",\"period\":1.88,\"atmosphere\":[\"CO2\",\"N2\",\"Ar\"]},{\"id\":5,\"name\":\"Jupiter\",\"type\":\"Gas giant\",\"period\":11.86,\"atmosphere\":[\"H2\",\"He\"]},{\"id\":6,\"name\":\"Saturn\",\"type\":\"Gas giant\",\"period\":29.46,\"atmosphere\":[\"H2\",\"He\"]},{\"id\":7,\"name\":\"Uranus\",\"type\":\"Ice giant\",\"period\":84.01,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]},{\"id\":8,\"name\":\"Neptune\",\"type\":\"Ice giant\",\"period\":164.8,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]}]"
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
          "path": "planets",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "[{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]},{\"id\":2,\"name\":\"Venus\",\"type\":\"Terrestrial planet\",\"period\":0.62,\"atmosphere\":[\"CO2\",\"N2\"]},{\"id\":3,\"name\":\"Earth\",\"type\":\"Terrestrial planet\",\"period\":1,\"atmosphere\":[\"N2\",\"O2\",\"Ar\"]},{\"id\":4,\"name\":\"Mars\",\"type\":\"Terrestrial planet\",\"period\":1.88,\"atmosphere\":[\"CO2\",\"N2\",\"Ar\"]},{\"id\":5,\"name\":\"Jupiter\",\"type\":\"Gas giant\",\"period\":11.86,\"atmosphere\":[\"H2\",\"He\"]},{\"id\":6,\"name\":\"Saturn\",\"type\":\"Gas giant\",\"period\":29.46,\"atmosphere\":[\"H2\",\"He\"]},{\"id\":7,\"name\":\"Uranus\",\"type\":\"Ice giant\",\"period\":84.01,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]},{\"id\":8,\"name\":\"Neptune\",\"type\":\"Ice giant\",\"period\":164.8,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]}]",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
          "method": "GET",
          "path": "planets/1",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]}",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      "...",
      {
        "id": "${unique-id}",
        "request": {
          "method": "GET",
          "path": "planets/8",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "{\"id\":8,\"name\":\"Neptune\",\"type\":\"Ice giant\",\"period\":164.8,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]}",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
            "method": "POST",
            "path": "planets",
            "pathReg": "",
            "headers": {}
        },
        "response": {
            "body": "{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]}",
            "file": "",
            "status": 201,
            "headers": {
              "Content-Type": "application/json"
            }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
          "method": "PUT",
          "path": "planets/1",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "{\"id\":1,\"name\":\"Mercury\",\"type\":\"Terrestrial planet\",\"period\":0.24,\"atmosphere\":[]}",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      "...",
      {
        "id": "${unique-id}",
        "request": {
          "method": "PUT",
          "path": "planets/8",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "{\"id\":8,\"name\":\"Neptune\",\"type\":\"Ice giant\",\"period\":164.8,\"atmosphere\":[\"H2\",\"He\",\"CH4\"]}",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      {
        "id": "${unique-id}",
        "request": {
          "method": "DELETE",
          "path": "planets/1",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      },
      "...",
      {
        "id": "${unique-id}",
        "request": {
          "method": "DELETE",
          "path": "planets/8",
          "pathReg": "",
          "headers": {}
        },
        "response": {
          "body": "",
          "file": "",
          "status": 200,
          "headers": {
            "Content-Type": "application/json"
          }
        }
      }
    ]
    ```


### WebSocket

---
## File Configuration

---
## API

