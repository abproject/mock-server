# Rest Entity API

URL: `/_api/rest/entities`

## Rest Entity API Description

| Path                        | Method   | Description                                           |           Request Body           |              Response Body               | Success Status | Failed Status |
| --------------------------- | -------- | ----------------------------------------------------- | :------------------------------: | :--------------------------------------: | :------------: | :-----------: |
| `/_api/rest/entities`       | `GET`    | Returns the list of all entities                      |                -                 | List of [Entity Model](#restentitymodel) |    **200**     |       -       |
| `/_api/rest/entities`       | `POST`   | Creates new entity                                    | [Entity Model](#restentitymodel) |     [Entity Model](#restentitymodel)     |    **201**     |       -       |
| `/_api/rest/entities`       | `DELETE` | Deletes all entities                                  |                -                 |                    -                     |    **204**     |       -       |
| `/_api/rest/entities/:name` | `GET`    | Returns entity by `name` or error if not found        |                -                 |     [Entity Model](#restentitymodel)     |    **200**     |    **404**    |
| `/_api/rest/entities/:name` | `PUT`    | Sets new entity by `name`, returns error if not found | [Entity Model](#restentitymodel) |     [Entity Model](#restentitymodel)     |    **200**     |    **404**    |
| `/_api/rest/entities/:name` | `DELETE` | Deletes entity by `name`, returns error if not found  |                -                 |                    -                     |    **204**     |    **404**    |

## Rest Entity API Examples

### GET `/_api/rest/entities`

**Response**
!INCLUDECODE "examples/api-rest/entities/api-get-all-response.json" (json)

### POST `/_api/rest/entities`

**Request:**

!INCLUDECODE "examples/api-rest/entities/api-post-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/entities/api-post-response.json" (json)

### GET `/_api/rest/entities/:id`

**Response**
!INCLUDECODE "examples/api-rest/entities/api-get-response.json" (json)

### PUT `/_api/rest/entities/:id`

**Request:**

!INCLUDECODE "examples/api-rest/entities/api-put-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/entities/api-put-response.json" (json)
