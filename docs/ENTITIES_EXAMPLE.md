# Entities

1\.  [YAML](#yaml)  
2\.  [JSON](#json)  
3\.  [API](#api)  
4\.  [Entity Model](#entitymodel)  

<a name="yaml"></a>

## 1\. YAML

**Create file `examples/entities/data-all.json` with content:**

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

**Create file `examples/entities/data-new.json` with content:**

```json
{
  "atmosphere": ["N2", "CH4", "CO"],
  "id": 9,
  "name": "Pluto",
  "period": 248,
  "type": "Dwarf planet"
}

```

**Create file `examples/entities/config.yaml` with content:**

```yaml
rest:
  entities:
    - name: planets
      dataAll: examples/entities/data-all.json
      dataNew: examples/entities/data-new.json
      id: id

```

**Final structure:**

```bash
./examples/entities/data-all.json
./examples/entities/data-new.json
./examples/entities/config.yaml
```

**Run in terminal:**

```bash
docker run -p 4242:8000 -v ${PWD}/examples:/examples abezpalov/mock-server -file=/examples/entities/config.yaml
```

**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
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
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center">List of <a href="#entitymodel">Models</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>

<a name="json"></a>

## 2\. JSON

**Create file `examples/entities/data-all.json` with content:**

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

**Create file `examples/entities/data-new.json` with content:**

```json
{
  "atmosphere": ["N2", "CH4", "CO"],
  "id": 9,
  "name": "Pluto",
  "period": 248,
  "type": "Dwarf planet"
}

```

**Create file `examples/entities/config.json` with content:**

```json
{
  "rest": {
    "entities": [
      {
        "name": "planets",
        "dataAll": "examples/entities/data-all.json",
        "dataNew": "examples/entities/data-new.json",
        "id": "id"
      }
    ]
  }
}

```

**Final structure:**

```bash
./examples/entities/data-all.json
./examples/entities/data-new.json
./examples/entities/config.json
```

**Run in terminal:**

```bash
docker run -p 4242:8000 -v ${PWD}/examples:/examples abezpalov/mock-server -file=/examples/entities/config.json
```

**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
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
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center">List of <a href="#entitymodel">Models</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
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

**Create file `examples/entities/data-all.json` with content:**

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

**Create file `examples/entities/data-new.json` with content:**

```json
{
  "atmosphere": ["N2", "CH4", "CO"],
  "id": 9,
  "name": "Pluto",
  "period": 248,
  "type": "Dwarf planet"
}

```

**Final structure:**

```bash
./examples/entities/data-all.json
./examples/entities/data-new.json
```

**Create file for all data**
**Send `POST` request to URL `http://localhost:4242/_api/files` with:**

- body as `form-data`
  - `file: data-all.json`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/entities/data-all.json' http://localhost:4242/_api/files

## Response, e.g.:
{"id":"5fb90badb37c4821b6d95526a41a9504","name":"data-all.json","length":1166}
```

> **`file-all-data-id`** ("5fb90badb37c4821b6d95526a41a9504") should be used in the request below

**Create file for new data**
**Send `POST` request to URL `http://localhost:4242/_api/files` with:**

- body as `form-data`
  - `file: data-new.json`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/entities/data-new.json' http://localhost:4242/_api/files

## Response, e.g.:
{"id":"95af5a25367941baa2ff6cd471c483f1","name":"data-new.json","length":113}
```

> **`file-new-data-id`** ("95af5a25367941baa2ff6cd471c483f1") should be used in the request below

**Send `POST` request to URL `http://localhost:4242/_api/rest/entities` with body:**

```json
{
  "name": "planets",
  "dataAll": "<changed-to-id-all>",
  "dataNew": "<changed-to-id-new>",
  "id": "id"
}

```

e.g., with `curl` _(copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/entities \
-H "Content-Type: application/json" \
-d @- << EOF
```

```json
{
  "name": "planets",
  "dataAll": "<changed-to-id-all>",
  "dataNew": "<changed-to-id-new>",
  "id": "id"
}

```

```
EOF
```

**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
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
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center">List of <a href="#entitymodel">Models</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>

<a name="entitymodel"></a>

## 4\. Entity Model

```json
{
  "atmosphere": ["N2", "O2", "Ar"],
  "id": 3,
  "name": "Earth",
  "period": 1,
  "type": "Terrestrial planet"
}

```
