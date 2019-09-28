# Rest Endpoint Response Model

| Field Name | Type                  | Description                                                                                                                            |
| ---------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `body`     | `any`                 | Response body could be type including JSON objects. <br>**Default:** empty string                                                      |
| `bodyFile` | `string`              | Response body is the content of the file. If `bodyFile` value is not empty then `body` value is ignored. <br>**Default:** empty string |
| `status`   | `integer`             | Response HTTP status code. <br>**Default:** `0`                                                                                        |
| `headers`  | `map<string, string>` | Response Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                           |
