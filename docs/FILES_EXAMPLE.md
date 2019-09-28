# Files

1\.  [YAML](#yaml)  
2\.  [JSON](#json)  
3\.  [API](#api)  

<a name="yaml"></a>

## 1\. YAML

**Create file `examples/files/hello.txt` with content:**

```txt
Hello from file!
```

**Create file `examples/files/hello.json` with content:**

```json
{
  "message": "Hello, World!"
}

```

**Create file `examples/files/config.yaml` with content:**

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

**Final structure:**

```bash
<your-path>/examples/files/hello.json
<your-path>/examples/files/hello.txt
<your-path>/examples/files/config.yaml
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/files/config.yaml
```

**Check `hello.txt`:**

```bash
curl -v http://localhost:4242/hello-txt

# Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

**Check `hello.json`:**

```bash
curl -v http://localhost:4242/hello-json

# Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```

<a name="json"></a>

## 2\. JSON

**Create file `examples/files/hello.txt` with content:**

```txt
Hello from file!
```

**Create file `examples/files/hello.json` with content:**

```json
{
  "message": "Hello, World!"
}

```

**Create file `examples/files/config.json` with content:**

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

**Final structure:**

```bash
<your-path>/examples/files/hello.json
<your-path>/examples/files/hello.txt
<your-path>/examples/files/config.json
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/files/config.json
```

**Check `hello.txt`:**

```bash
curl -v http://localhost:4242/hello-txt

# Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

**Check `hello.json`:**

```bash
curl -v http://localhost:4242/hello-json

# Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```

<a name="api"></a>

## 3\. API

**Run in terminal:**

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

**Create file `examples/files/hello.txt` with content:**

```txt
Hello from file!
```

**Create file `examples/files/hello.json` with content:**

```json
{
  "message": "Hello, World!"
}

```

**Final structure:**

```bash
<your-path>/examples/files/hello.json
<your-path>/examples/files/hello.txt
```

**hello.txt**

**Send `POST` request to URL `http://localhost:4242/_api/files` with:**

- body as `form-data`
  - `file: hello.txt`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/files/hello.txt' http://localhost:4242/_api/files

# Response, e.g.:
{"id":"6694d2c422ac4208a0072939487f6999","name":"hello.txt","length":16}
```

**Copy `id` from response, e.g., `6694d2c422ac4208a0072939487f6999`.**

**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:**

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

**Send `POST` request to URL `http://localhost:4242/_api/files` with:**

- body as `form-data`
  - `file: hello.json`
- headers:
  - `ContentType: application/x-www-form-urlencoded`

e.g., with `curl`:

```bash
curl -F 'file=@examples/files/hello.json' http://localhost:4242/_api/files

# Response, e.g.:
{"id":"9566c74d10034c4dbbbb0407d1e2c649","name":"hello.json","length":16}
```

**Copy `id` from response, e.g., `9566c74d10034c4dbbbb0407d1e2c649`.**

**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:**

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

**Check `hello.txt`:**

```bash
curl -v http://localhost:4242/hello-txt

# Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

**Check `hello.json`:**

```bash
curl -v http://localhost:4242/hello-json

# Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```
