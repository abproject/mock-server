# Rest Endpoint API

URL: `/_api/rest/endpoints`

## Rest Endpoint API Description

| Path                       | Method   | Description                                                         |                Request Body                 |                    Response Body                    | Success Status | Failed Status |
| -------------------------- | -------- | ------------------------------------------------------------------- | :-----------------------------------------: | :-------------------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoints`     | `GET`    | Returns the list of all endpoints configurations                    |                      -                      | List of [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/endpoints`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **201**     |       -       |
| `/_api/rest/endpoints`     | `DELETE` | Deletes all endpoints configuration                                 |                      -                      |                          -                          |    **204**     |       -       |
| `/_api/rest/endpoints/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                      -                      |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoints/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Model](#restendpointrequestmodel) |     [Endpoint Model](#restendpointrequestmodel)     |    **200**     |    **404**    |
| `/_api/rest/endpoints/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                      -                      |                          -                          |    **204**     |    **404**    |

## Rest Endpoint API Examples

### GET `/_api/rest/endpoints`

**Response**
!INCLUDECODE "examples/api-rest/endpoints/api-get-all-response.json" (json)

### POST `/_api/rest/endpoints`

**Request:**

!INCLUDECODE "examples/api-rest/endpoints/api-post-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/endpoints/api-post-response.json" (json)

### GET `/_api/rest/endpoints/:id`

**Response**
!INCLUDECODE "examples/api-rest/endpoints/api-get-response.json" (json)

### PUT `/_api/rest/endpoints/:id`

**Request:**

!INCLUDECODE "examples/api-rest/endpoints/api-put-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/endpoints/api-put-response.json" (json)
