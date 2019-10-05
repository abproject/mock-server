# CRUD

There are different combinations of `body` and `bodyFile` are shown in the examples, it is recommended to select one and stick to it.

> `:id` doesn't analyse the `id` of object `body` or `bodyFile`, so if there is need to have `404 Status` in case of not existing `id` or dynamic answer depending on `id` check [Entries](#entries) example.

1\.  [YAML](#yaml)  
2\.  [JSON](#json)  
3\.  [API](#api)  
4\.  [Models](#models)  
4.1\.  [CRUD Model All](#crudmodelall)  
4.2\.  [CRUD Model One](#crudmodelone)  

<a name="yaml"></a>

## 1\. YAML

**Create file `examples/crud/data.json` with content:**

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

**Create file `examples/crud/data-id-3.json` with content:**

```json
{
  "id": 3,
  "name": "Earth",
  "type": "Terrestrial planet",
  "period": 1,
  "atmosphere": ["N2", "O2", "Ar"]
}

```

**Create file `examples/crud/config.yaml` with content:**

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
        bodyFile: examples/crud/data.json
    - request:
        method: GET
        path: planets/:id
      response:
        body: >
          { 
            "id": 3,
            "name": "Earth",
            "type": "Terrestrial planet",
            "period": 1,
            "atmosphere": ["N2", "O2", "Ar"]
          }
    - request:
        method: POST
        path: planets
      response:
        status: 201
        body: '{"id":3, "name":"Earth", "type":"Terrestrial planet", "period":1, "atmosphere": ["N2", "O2", "Ar"]}'
    - request:
        method: PUT
        path: planets/:id
      response:
        bodyFile: examples/crud/data-id-3.json
    - request:
        method: DELETE
        path: planets/:id
      response:
        status: 204

```

**Final structure:**

```bash
<your-path>/examples/crud/data.json
<your-path>/examples/crud/data-id-3.json
<your-path>/examples/crud/config.yaml
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/crud/config.yaml
```

**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
    <th rowspan="2">cURL</th>
    <th colspan="3">Response</th>
  </tr>
  <tr>
    <th>Type</th>
    <th>Endpoint</th>
    <th>Status</th>
    <th>Body</th>
    <th>Headers</th>
  </tr>
  <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelall">Body All</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v -X POST http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X PUT http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v -X PUT  http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X DELETE http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>

<a name="json"></a>

## 2\. JSON

**Create file `examples/crud/data.json` with content:**

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

**Create file `examples/crud/data-id-3.json` with content:**

```json
{
  "id": 3,
  "name": "Earth",
  "type": "Terrestrial planet",
  "period": 1,
  "atmosphere": ["N2", "O2", "Ar"]
}

```

**Create file `examples/crud/config.json` with content:**

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
          "bodyFile": "examples/crud/data.json"
        }
      },
      {
        "request": {
          "method": "GET",
          "path": "planets/:id"
        },
        "response": {
          "body": "{  \"id\": 3,\n  \"name\": \"Earth\",\n  \"type\": \"Terrestrial planet\",\n  \"period\": 1,\n  \"atmosphere\": [\"N2\", \"O2\", \"Ar\"]\n}\n"
        }
      },
      {
        "request": {
          "method": "POST",
          "path": "planets"
        },
        "response": {
          "status": 201,
          "body": "{\"id\":3, \"name\":\"Earth\", \"type\":\"Terrestrial planet\", \"period\":1, \"atmosphere\": [\"N2\", \"O2\", \"Ar\"]}"
        }
      },
      {
        "request": {
          "method": "PUT",
          "path": "planets/:id"
        },
        "response": {
          "bodyFile": "examples/crud/data-id-3.json"
        }
      },
      {
        "request": {
          "method": "DELETE",
          "path": "planets/:id"
        },
        "response": {
          "status": 204
        }
      }
    ]
  }
}

```

**Final structure:**

```bash
<your-path>/examples/crud/data.json
<your-path>/examples/crud/data-id-3.json
<your-path>/examples/crud/config.json
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/crud/config.json
```

**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
    <th rowspan="2">cURL</th>
    <th colspan="3">Response</th>
  </tr>
  <tr>
    <th>Type</th>
    <th>Endpoint</th>
    <th>Status</th>
    <th>Body</th>
    <th>Headers</th>
  </tr>
  <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelall">Body All</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v -X POST http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X PUT http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v -X PUT  http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X DELETE http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>

<a name="api"></a>

## 3\. API

**Run in terminal:**

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

**Create file `examples/crud/data.json` with content:**

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

**Create file `examples/crud/data-id-3.json` with content:**

```json
{
  "id": 3,
  "name": "Earth",
  "type": "Terrestrial planet",
  "period": 1,
  "atmosphere": ["N2", "O2", "Ar"]
}

```

**Final structure:**

```bash
<your-path>/examples/crud/data.json
<your-path>/examples/crud/data-id-3.json
```

**Create file for all data**
**Send `POST` request to URL `http://localhost:4242/_api/files` with:**

- body as `form-data`
  - `file: data.json`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/crud/data.json' http://localhost:4242/_api/files

## Response, e.g.:
{"id":"95af5a25367941baa2ff6cd471c483f1","name":"data.json","length":517}
```

> **`file-all-data-id`** ("95af5a25367941baa2ff6cd471c483f1") should be used in the request below

**Create file for one instance**
**Send `POST` request to URL `http://localhost:4242/_api/files` with:**

- body as `form-data`
  - `file: data-id-3.json`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/crud/data-id-3.json' http://localhost:4242/_api/files

## Response, e.g.:
{"id":"5fb90badb37c4821b6d95526a41a9504","name":"data.json","length":116}
```

> **`file-one-instance-id`** ("5fb90badb37c4821b6d95526a41a9504") should be used in the requests below

**Global Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/global` with body:**

```json
{
  "response": {
    "status": 200,
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

e.g., with `curl` _(copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "response": {
    "status": 200,
    "headers": {
      "Content-Type": "application/json"
    }
  }
}

```

```
EOF
```

**Get All Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

```json
{
  "request": {
    "method": "GET",
    "path": "planets"
  },
  "response": {
    "bodyFile": "<changed-to-id>"
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `file-all-data-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "GET",
    "path": "planets"
  },
  "response": {
    "bodyFile": "<changed-to-id>"
  }
}

```

```
EOF
```

**Get One Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

```json
{
  "request": {
    "method": "GET",
    "path": "planets/:id"
  },
  "response": {
    "bodyFile": "<changed-to-id>"
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "GET",
    "path": "planets/:id"
  },
  "response": {
    "bodyFile": "<changed-to-id>"
  }
}

```

```
EOF
```

**Post Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

```json
{
  "request": {
    "method": "POST",
    "path": "planets"
  },
  "response": {
    "status": 201,
    "bodyFile": "<changed-to-id>"
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "POST",
    "path": "planets"
  },
  "response": {
    "status": 201,
    "bodyFile": "<changed-to-id>"
  }
}

```

```
EOF
```

**Put Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

```json
{
  "request": {
    "method": "PUT",
    "path": "planets/:id"
  },
  "response": {
    "bodyFile": "<changed-to-id>"
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "PUT",
    "path": "planets/:id"
  },
  "response": {
    "bodyFile": "<changed-to-id>"
  }
}

```

```
EOF
```

**Delete Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

```json
{
  "request": {
    "method": "DELETE",
    "path": "planets/:id"
  },
  "response": {
    "status": 204
  }
}

```

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "request": {
    "method": "DELETE",
    "path": "planets/:id"
  },
  "response": {
    "status": 204
  }
}

```

```
EOF
```

**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
    <th rowspan="2">cURL</th>
    <th colspan="3">Response</th>
  </tr>
  <tr>
    <th>Type</th>
    <th>Endpoint</th>
    <th>Status</th>
    <th>Body</th>
    <th>Headers</th>
  </tr>
  <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelall">Body All</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v -X POST http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X PUT http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v -X PUT  http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X DELETE http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>

<a name="models"></a>

## 4\. Models

<a name="crudmodelall"></a>

### 4.1\. CRUD Model All

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

<a name="crudmodelone"></a>

### 4.2\. CRUD Model One

```json
{
  "id": 3,
  "name": "Earth",
  "type": "Terrestrial planet",
  "period": 1,
  "atmosphere": ["N2", "O2", "Ar"]
}

```
