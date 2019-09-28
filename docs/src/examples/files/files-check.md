Check `hello.txt`:

```bash
curl -v http://localhost:4242/hello-txt

# Response
...
< HTTP/1.1 200 OK
< Content-Type: text/plain
...
Hello from file!
```

Check `hello.json`:

```bash
curl -v http://localhost:4242/hello-json

# Response
...
< HTTP/1.1 200 OK
< Content-Type: application/json
...
<
{
  "message": "Hello, World!"
}
```
