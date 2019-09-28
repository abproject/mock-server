# Rest Endpoint Model

| Field Name | Type                                                       | Description                                      |
| ---------- | ---------------------------------------------------------- | ------------------------------------------------ |
| `id`       | `string`                                                   | Unique Endpoint ID. **Generates by mock-server** |
| `request`  | [Rest Endpoint Request Model](#restendpointrequestmodel)   | Request configuration model                      |
| `response` | [Rest Endpoint Response Model](#restendpointresponsemodel) | Response configuration model                     |

!INCLUDE "rest-endpoint-request.md", 1

!INCLUDE "rest-endpoint-response.md", 1
