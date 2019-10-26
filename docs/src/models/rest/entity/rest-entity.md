# Rest Entity Model

| Field Name | Type     | Description                                                                               |
| ---------- | -------- | ----------------------------------------------------------------------------------------- |
| `name`     | `string` | Entity name, will be used as path e.g., `/some-name`. <br>**Required, Not Empty, Unique** |
| `dataAll`     | `string` | Name of file with database (array of entity objects)              
| `dataNew`     | `string` | Name of file with response for creating new entity (POST request)
| `id`       | `string` | Property identifier name (`@id`) of entity                             |
