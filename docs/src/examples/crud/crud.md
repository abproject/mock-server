# CRUD

There are different combinations of `body` and `bodyFile` are shown in the examples, it is recommended to select one and stick to it.

> `:id` doesn't analyse the amount of objects in `body` or `bodyFile`, if there is need to have `404 Status` in case of not existing `id`, check [Entries](#entries) example.

!TOC

## YAML

!INCLUDE "docs/src/examples/crud/crud-create.md"

**Create file `examples/crud/config.yaml` with content:**

!INCLUDECODE "examples/crud/config.yaml" (yaml)

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

!INCLUDE "docs/src/examples/crud/crud-check.md", 2

## JSON

!INCLUDE "docs/src/examples/crud/crud-create.md"

**Create file `examples/crud/config.json` with content:**

!INCLUDECODE "examples/crud/config.json" (json)

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

!INCLUDE "docs/src/examples/crud/crud-check.md", 2

!INCLUDE "docs/src/examples/crud/crud-api.md", 1

## Models

### CRUD Model All

!INCLUDECODE "examples/crud/data.json" (json)

### CRUD Model One

!INCLUDECODE "examples/crud/data-id-3.json" (json)
