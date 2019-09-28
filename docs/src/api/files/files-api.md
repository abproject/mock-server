# Files API

URL: `/_api/files`

## Files API Description

| Path              | Method   | Description                                                     |                                                                        Request Body                                                                         |          Response Body           | Success Status | Failed Status |
| ----------------- | -------- | --------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------: | :------------------------------: | :------------: | :-----------: |
| `/_api/files`     | `GET`    | Returns the list of all files configurations                    |                                                                              -                                                                              | List of [File Model](#filemodel) |    **200**     |       -       |
| `/_api/files`     | `POST`   | Creates new file entity                                         | <div style="text-align: left">**Body** as `form-data`:<br>`file: <file content>`<br>**Headers**:<br>`Content-Type: application/x-www-form-urlencoded`</div> |     [File Model](#filemodel)     |    **201**     |       -       |
| `/_api/files`     | `DELETE` | Deletes all files configuration                                 |                                                                              -                                                                              |                -                 |    **204**     |       -       |
| `/_api/files/:id` | `GET`    | Returns file by `id` or error if not found                      |                                                                              -                                                                              |     [File Model](#filemodel)     |    **200**     |    **404**    |
| `/_api/files/:id` | `PUT`    | Sets new file configuration by `id`, returns error if not found | <div style="text-align: left">**Body** as `form-data`:<br>`file: <file content>`<br>**Headers**:<br>`Content-Type: application/x-www-form-urlencoded`</div> |     [File Model](#filemodel)     |    **200**     |    **404**    |
| `/_api/files/:id` | `DELETE` | Deletes file configuration by `id`, returns error if not found  |                                                                              -                                                                              |                -                 |    **204**     |    **404**    |

## Files API Examples

### GET `/_api/files`

**Response**
!INCLUDECODE "examples/api-files/api-get-all-response.json" (json)

### GET `/_api/files/:id`

**Response**
!INCLUDECODE "examples/api-files/api-get-response.json" (json)
