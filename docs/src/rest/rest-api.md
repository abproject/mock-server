# [mock-server](../README.md): REST API

!TOC

## Endpoint

`/_api/rest/endpoint`

> Configures endpoints of mock-server

### Endpoint API Description

| Path                      | Method   | Description                                                         |            Request Body            |               Response Body                | Success Status | Failed Status |
| ------------------------- | -------- | ------------------------------------------------------------------- | :--------------------------------: | :----------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoint`     | `GET`    | Returns the list of all endpoints configurations                    |                 -                  | List of [Endpoint Object](#endpointobject) |    **200**     |       -       |
| `/_api/rest/endpoint`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Object](#endpointobject) |     [Endpoint Object](#endpointobject)     |    **201**     |       -       |
| `/_api/rest/endpoint`     | `DELETE` | Deletes all endpoints configuration                                 |                 -                  |                     -                      |    **204**     |       -       |
| `/_api/rest/endpoint/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                 -                  |     [Endpoint Object](#endpointobject)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Object](#endpointobject) |     [Endpoint Object](#endpointobject)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                 -                  |                     -                      |    **204**     |    **404**    |

### Endpoint Body Examples

**Request:**

!INCLUDECODE "examples/api-rest/api-post-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/api-post-response.json" (json)

!INCLUDE "docs/src/rest/rest-endpoint.md", 2

!INCLUDE "docs/src/rest/rest-endpoint-request.md", 2

!INCLUDE "docs/src/rest/rest-endpoint-response.md", 2
