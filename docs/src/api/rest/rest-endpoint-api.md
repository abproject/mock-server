# Rest Endpoint API

URL: `/_api/rest/endpoint`

## Rest Endpoint API Description

| Path                      | Method   | Description                                                         |                Request Body                 |                    Response Body                    | Success Status | Failed Status |
| ------------------------- | -------- | ------------------------------------------------------------------- | :-----------------------------------------: | :-------------------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoint`     | `GET`    | Returns the list of all endpoints configurations                    |                      -                      | List of [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/endpoint`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **201**     |       -       |
| `/_api/rest/endpoint`     | `DELETE` | Deletes all endpoints configuration                                 |                      -                      |                          -                          |    **204**     |       -       |
| `/_api/rest/endpoint/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                      -                      |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                      -                      |                          -                          |    **204**     |    **404**    |

## Rest Endpoint API Examples

### GET `/_api/rest/endpoint`

**Response**
!INCLUDECODE "examples/api-rest/api-get-all-response.json" (json)

### POST `/_api/rest/endpoint`

**Request:**

!INCLUDECODE "examples/api-rest/api-post-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/api-post-response.json" (json)

### GET `/_api/rest/endpoint/:id`

**Response**
!INCLUDECODE "examples/api-rest/api-get-response.json" (json)

### PUT `/_api/rest/endpoint/:id`

**Request:**

!INCLUDECODE "examples/api-rest/api-put-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/api-put-response.json" (json)
