# API

**Run in terminal:**

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

!INCLUDE "entities-create.md"

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

# Response, e.g.:
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

# Response, e.g.:
{"id":"95af5a25367941baa2ff6cd471c483f1","name":"data-new.json","length":113}
```

> **`file-new-data-id`** ("95af5a25367941baa2ff6cd471c483f1") should be used in the request below

**Send `POST` request to URL `http://localhost:4242/_api/rest/entities` with body:**

!INCLUDECODE "examples/entities/config-api.json" (json)

e.g., with `curl` _(copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/entities \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/entities/config-api.json" (json)

```
EOF
```

!INCLUDE "entities-check.md", 2
