**Checks:**

<table>
  <tr>
    <th colspan="2">Request</th>
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
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center">List of <a href="#entitymodel">Models</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
   <tr>
    <td style="text-align:center"><code>GET</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>POST</code></td>
    <td>/planets</td>
    <td style="text-align:center"><code>201</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>200</code></td>
    <td style="text-align:center"><a href="#entitymodel">Model</a>
    </td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>PUT</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/3</td>
    <td style="text-align:center"><code>204</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
  <tr>
    <td style="text-align:center"><code>DELETE</code></td>
    <td>/planets/42</td>
    <td style="text-align:center"><code>404</code></td>
    <td style="text-align:center">-</td>
    <td><code>Content-Type: application/json</code></td>
  </tr>
</table>
