# [mock-server](../README.md): REST JSON

1\.  [Endpoint](#endpoint)  
1.1\.  [Endpoint Config Example](#endpointconfigexample)  
1.2\.  [Endpoint Object](#endpointobject)  
1.3\.  [Endpoint Request Object](#endpointrequestobject)  
1.4\.  [Endpoint Response Object](#endpointresponseobject)  

<a name="endpoint"></a>

## 1\. Endpoint

> Configures endpoints of mock-server

<a name="endpointconfigexample"></a>

### 1.1\. Endpoint Config Example

```json
{
  "rest": {
    "endpoints": [
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
    ]
  }
}

```

<a name="endpointobject"></a>

### 1.2\. Endpoint Object

| Field Name | Type                                                | Description                                      |
| ---------- | --------------------------------------------------- | ------------------------------------------------ |
| `id`       | `string`                                            | Unique Endpoint ID. **Generates by mock-server** |
| `request`  | [Endpoint Request Object](#endpointrequestobject)   | Request configuration object                     |
| `response` | [Endpoint Response Object](#endpointresponseobject) | Response configuration object                    |

<a name="endpointrequestobject"></a>

### 1.3\. Endpoint Request Object

| Field Name | Type                  | Description                                                                                                                              |
| ---------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `method`   | `string`              | Method name, e.g., `GET`, `POST`, `DELETE`, `PUT`. <br>Empty string means all type of requests. <br>**Default:** empty string            |
| `path`     | `string`              | Request Endpoint path, e.g., `/my-path`. <br>**Default:** empty string                                                                   |
| `pathReg`  | `string`              | Request Endpoint path as regular expression. If `pathReg` value is not empty then `path` value is ignored. <br>**Default:** empty string |
| `headers`  | `map<string, string>` | Request Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                              |

<a name="endpointresponseobject"></a>

### 1.4\. Endpoint Response Object

| Field Name | Type                  | Description                                                                                                                            |
| ---------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `body`     | `any`                 | Response body could be type including JSON objects. <br>**Default:** empty string                                                      |
| `bodyFile` | `string`              | Response body is the content of the file. If `bodyFile` value is not empty then `body` value is ignored. <br>**Default:** empty string |
| `status`   | `integer`             | Response HTTP status code. <br>**Default:** `0`                                                                                        |
| `headers`  | `map<string, string>` | Response Key-Value pairs of headers, e.g., <br> `"Content-Type": "application/json"` <br>**Default:** `null`                           |
