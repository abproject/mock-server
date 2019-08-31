# YAML

Create file `config.yml` with content:

!INCLUDECODE "examples/hello/config.yml" (yaml)

Run in terminal:

```bash
docker run -p 4242:8000 -v ${PWD}/config.yml:/config.yml abezpalov/mock-server -file=config.yml
```

!INCLUDE "hello-world-check.md"

# JSON

Create file `config.yml` with content:

!INCLUDECODE "examples/hello/config.json" (json)

Run in terminal:

```bash
docker run -p 4242:8000 -v ${PWD}/config.json:/config.json abezpalov/mock-server -file=config.json
```

!INCLUDE "hello-world-check.md"

# API

Another way to get the same `Hello World` configuration without config file but by using API requests only.

Run in terminal:

```bash
docker run -p 4242:8000 abezpalov/mock-server
```

Make `POST` request to URL `http://localhost:4242/_api/rest/endpoint` with body:

!INCLUDECODE "examples/hello/config-api.json" (json)

e.g., with `curl` **(please copy all 3 code blocks below and paste in terminal)**:

```bash
curl -X POST http://localhost:4242/_api/rest/endpoint \
-H "Content-Type: application/json" \
-d @- << EOF
```

!INCLUDECODE "examples/hello/config-api.json" (json)

```
EOF
```

!INCLUDE "hello-world-check.md"
