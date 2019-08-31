# [mock-server](../README.md): REST API

1\.  [Endpoint](#endpoint)  
1.1\.  [Endpoint API Description](#endpointapidescription)  
1.2\.  [Endpoint Body Examples](#endpointbodyexamples)  
1.3\.  [Endpoint Object](#endpointobject)  
1.4\.  [Endpoint Response Object](#endpointresponseobject)  
1.5\.  [Endpoint Request Object](#endpointrequestobject)  

<a name="endpoint"></a>

## 1\. Endpoint

`/_api/rest/endpoint`

> Configures endpoints of mock-server

<a name="endpointapidescription"></a>

### 1.1\. Endpoint API Description

| Path                      | Method   | Description                                                         |            Request Body            |               Response Body                | Success Status | Failed Status |
| ------------------------- | -------- | ------------------------------------------------------------------- | :--------------------------------: | :----------------------------------------: | :------------: | :-----------: |
| `/_api/rest/endpoint`     | `GET`    | Returns the list of all endpoints configurations                    |                 -                  | List of [Endpoint Object](#endpointobject) |    **200**     |       -       |
| `/_api/rest/endpoint`     | `POST`   | Creates new endpoint entity                                         | [Endpoint Object](#endpointobject) |     [Endpoint Object](#endpointobject)     |    **201**     |       -       |
| `/_api/rest/endpoint`     | `DELETE` | Deletes all endpoints configuration                                 |                 -                  |                     -                      |    **204**     |       -       |
| `/_api/rest/endpoint/:id` | `GET`    | Returns endpoint by `id` or error if not found                      |                 -                  |     [Endpoint Object](#endpointobject)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `PUT`    | Sets new endpoint configuration by `id`, returns error if not found | [Endpoint Object](#endpointobject) |     [Endpoint Object](#endpointobject)     |    **200**     |    **404**    |
| `/_api/rest/endpoint/:id` | `DELETE` | Deletes endpoint configuration by `id`, returns error if not found  |                 -                  |                     -                      |    **204**     |    **404**    |

<a name="endpointbodyexamples"></a>

### 1.2\. Endpoint Body Examples

**Request:**

```json
{
  "request": {
    "method": "GET",
    "path": "hello"
  },
  "response": {
    "status": 200,
    "body": "Hello, World!",
    "headers": {
      "Content-Type": "text/html"
    }
  }
}

```

**Response:**

```json
{
  "id": ":id",
  "request": {
    "method": "GET",
    "path": "hello",
    "pathReg": "",
    "headers": null
  },
  "response": {
    "status": 200,
    "body": "Hello, World!",
    "bodyFile": "",
    "headers": {
      "Content-Type": "text/html"
    }
  }
}

```

<a name="endpointobject"></a>

### 1.3\. Endpoint Object

| Field Name | Type                                                | Description                           |
| ---------- | --------------------------------------------------- | ------------------------------------- |
| `id`       | `string`                                            | Unique Endpoint ID. **Response only** |
| `request`  | [Endpoint Request Object](#endpointrequestobject)   | Request configuration object          |
| `response` | [Endpoint Response Object](#endpointresponseobject) | Response configuration object         |

<a name="endpointresponseobject"></a>

### 1.4\. Endpoint Response Object

| Field Name | Type                  | Description                                                                                                                              |
| ---------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `method`   | `string`              | Method name, e.g., `GET`, `POST`, `DELETE`, `PUT`. <br>Empty string means all type of requests. <br>**Default:** empty string            |
| `path`     | `string`              | Request Endpoint path, e.g., `/my-path`. <br>**Default:** empty string                                                                   |
| `pathReg`  | `string`              | Request Endpoint path as regular expression. If `pathReg` value is not empty then `path` value is ignored. <br>**Default:** empty string |
| `headers`  | `map<string, string>` | Request Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                              |

<a name="endpointrequestobject"></a>

### 1.5\. Endpoint Request Object

| Field Name | Type                  | Description                                                                                                                            |
| ---------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `body`     | `any`                 | Response body could be type including JSON objects. <br>**Default:** empty string                                                      |
| `bodyFile` | `string`              | Response body is the content of the file. If `bodyFile` value is not empty then `body` value is ignored. <br>**Default:** empty string |
| `status`   | `integer`             | Response HTTP status code. <br>**Default:** `0`                                                                                        |
| `headers`  | `map<string, string>` | Response Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                           |
