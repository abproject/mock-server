# Entities

!TOC

## YAML

!INCLUDE "docs/src/examples/entities/entities-create.md"

**Create file `examples/entities/config.yaml` with content:**

!INCLUDECODE "examples/entities/config.yaml" (yaml)

**Final structure:**

```bash
./examples/entities/data-all.json
./examples/entities/data-new.json
./examples/entities/config.yaml
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/entities/config.yaml
```

!INCLUDE "docs/src/examples/entities/entities-check.md", 2

## JSON

!INCLUDE "docs/src/examples/entities/entities-create.md"

**Create file `examples/entities/config.json` with content:**

!INCLUDECODE "examples/entities/config.json" (json)

**Final structure:**

```bash
./examples/entities/data-all.json
./examples/entities/data-new.json
./examples/entities/config.json
```

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/entities/config.json
```

!INCLUDE "docs/src/examples/entities/entities-check.md", 2

!INCLUDE "docs/src/examples/entities/entities-api.md", 1

## Entity Model

!INCLUDECODE "examples/entities/data-id-3.json" (json)
