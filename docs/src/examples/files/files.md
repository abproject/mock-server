# YAML

!INCLUDE "files-create.md"

Create file `examples/files/config.yaml` with content:

!INCLUDECODE "examples/files/config.yaml" (yaml)

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

!INCLUDE "files-check.md"

# JSON

!INCLUDE "files-create.md"

Create file `examples/files/config.json` with content:

!INCLUDECODE "examples/files/config.json" (yaml)

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

!INCLUDE "files-check.md"

# API

Run in terminal:

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

!INCLUDE "files-create.md"

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

# Response, e.g.:
{"id":"6694d2c422ac4208a0072939487f6999","name":"hello.txt","length":16}
```

Copy `id` from response, e.g., `6694d2c422ac4208a0072939487f6999`.

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:

!INCLUDECODE "examples/files/config-api-txt.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoints \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/files/config-api-txt.json" (json)

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

# Response, e.g.:
{"id":"9566c74d10034c4dbbbb0407d1e2c649","name":"hello.json","length":16}
```

Copy `id` from response, e.g., `9566c74d10034c4dbbbb0407d1e2c649`.

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:

!INCLUDECODE "examples/files/config-api-json.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoints \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/files/config-api-json.json" (json)

```
EOF
```

!INCLUDE "files-check.md"
