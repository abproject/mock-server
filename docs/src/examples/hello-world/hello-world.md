# YAML

**Create file `examples/hello/config.yaml` with content:**

!INCLUDECODE "examples/hello/config.yaml" (yaml)

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/hello/config.yaml
```

!INCLUDE "hello-world-check.md"

# JSON

**Create file `examples/hello/config.json` with content:**

!INCLUDECODE "examples/hello/config.json" (json)

**Run in terminal:**

```bash
docker run -p 4242:8000 \
-v ${PWD}/examples:/examples \
abezpalov/mock-server \
-file=/examples/hello/config.json
```

!INCLUDE "hello-world-check.md"

# API

Another way to get the same `Hello World` configuration without config file but by using API requests only.

**Run in terminal:**

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

**Make `POST` request to URL `http://localhost:4242/_api/rest/endpoints` with body:**

!INCLUDECODE "examples/hello/config-api.json" (json)

e.g., with `curl` _(copy all 3 code blocks below and paste in terminal)_:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoints \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/hello/config-api.json" (json)

```
EOF
```

!INCLUDE "hello-world-check.md"
