# API

**Run in terminal:**

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

!INCLUDE "crud-create.md"

**Final structure:**

```bash
./examples/crud/data.json
./examples/crud/data-id-3.json
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

# Response, e.g.:
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

# Response, e.g.:
{"id":"5fb90badb37c4821b6d95526a41a9504","name":"data.json","length":116}
```

> **`file-one-instance-id`** ("5fb90badb37c4821b6d95526a41a9504") should be used in the requests below

**Global Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/global` with body:**

!INCLUDECODE "examples/crud/config-api-global.json" (json)

e.g., with `curl` _(copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/crud/config-api-global.json" (json)

```
EOF
```

**Get All Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

!INCLUDECODE "examples/crud/config-api-get-all.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `file-all-data-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/crud/config-api-get-all.json" (json)

```
EOF
```

**Get One Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

!INCLUDECODE "examples/crud/config-api-get.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/crud/config-api-get.json" (json)

```
EOF
```

**Post Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

!INCLUDECODE "examples/crud/config-api-post.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/crud/config-api-post.json" (json)

```
EOF
```

**Put Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

!INCLUDECODE "examples/crud/config-api-put.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/crud/config-api-put.json" (json)

```
EOF
```

**Delete Endpoint**
**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:**

!INCLUDECODE "examples/crud/config-api-delete.json" (json)

e.g., with `curl` _(replace `changed-to-id` to `file-one-instance-id` above, copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/global \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/crud/config-api-delete.json" (json)

```
EOF
```

!INCLUDE "crud-check.md", 2
