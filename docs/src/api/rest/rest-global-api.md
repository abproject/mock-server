# Rest Global API

URL: `/_api/rest/global`

## Rest Global API Description

| Path                | Method   | Description                            |                Request Body                 |                Response Body                | Success Status | Failed Status |
| ------------------- | -------- | -------------------------------------- | :-----------------------------------------: | :-----------------------------------------: | :------------: | :-----------: |
| `/_api/rest/global` | `GET`    | Returns global endpoint configurations |                      -                      | [Endpoint Model](#restendpointrequestmodel) |    **200**     |       -       |
| `/_api/rest/global` | `POST`   | Creates new global endpoint entity     | [Endpoint Model](#restendpointrequestmodel) | [Endpoint Model](#restendpointrequestmodel) |    **201**     |       -       |
| `/_api/rest/global` | `DELETE` | Deletes global endpoint configuration  |                      -                      |                      -                      |    **204**     |       -       |

## Rest Global API Examples

### GET `/_api/rest/global`

**Response**
!INCLUDECODE "examples/api-rest/global/api-get-response.json" (json)

### POST `/_api/rest/global`

**Request:**

!INCLUDECODE "examples/api-rest/global/api-post-request.json" (json)

**Response:**

!INCLUDECODE "examples/api-rest/global/api-post-response.json" (json)
