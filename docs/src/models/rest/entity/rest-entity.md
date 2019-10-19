# Rest Entity Model

| Field Name | Type     | Description                                                                               |
| ---------- | -------- | ----------------------------------------------------------------------------------------- |
| `name`     | `string` | Entity name, will be used as path e.g., `/some-name`. <br>**Required, Not Empty, Unique** |
| `file`     | `string` | Response body is the content of the file (must be JSON format)                            |
| `id`       | `string` | Property identifier name (`@id`) of entity (`file` object)                                |
