# CRUD

In example there are only `bodyFile` configuration, but it is possible to use `body` as text instead.

> `:id` doesn't analyse the `id` of object `body` or `bodyFile`, so if there is need to have `404 Status` in case of not existing `id` or dynamic answer depending on `id` check [Entries](#entries) example.

!TOC

## YAML

!INCLUDE "docs/src/examples/crud/crud-create.md"

**Create file `examples/crud/config.yaml` with content:**

!INCLUDECODE "examples/crud/config.yaml" (yaml)

**Final structure:**

```bash
./examples/crud/data.json
./examples/crud/data-id-3.json
./examples/crud/config.yaml
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/crud/config.yaml
```

!INCLUDE "docs/src/examples/crud/crud-check.md", 2

## JSON

!INCLUDE "docs/src/examples/crud/crud-create.md"

**Create file `examples/crud/config.json` with content:**

!INCLUDECODE "examples/crud/config.json" (json)

**Final structure:**

```bash
./examples/crud/data.json
./examples/crud/data-id-3.json
./examples/crud/config.json
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/crud/config.json
```

!INCLUDE "docs/src/examples/crud/crud-check.md", 2

!INCLUDE "docs/src/examples/crud/crud-api.md", 1

## CRUD Model

!INCLUDECODE "examples/crud/data-id-3.json" (json)
