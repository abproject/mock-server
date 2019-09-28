**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
    <th rowspan="2">cURL</th>
    <th colspan="3">Response</th>
  </tr>
  <tr>
    <th>Type</th>
    <th>Endpoint</th>
    <th>Status</th>
    <th>Body</th>
    <th>Headers</th>
  </tr>
  <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelall">Body All</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td>
      <code>curl -v -X POST http://localhost:4242/planets</code>
    </td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X PUT http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td>
      <code>curl -v -X PUT  http://localhost:4242/planets/42</code>
    </td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#crudmodelone">Body One</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td>
      <code>curl -v -X DELETE http://localhost:4242/planets/3</code>
    </td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>
