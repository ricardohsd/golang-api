This is an example of http server written in Golang.

Run the project and in a new terminal window:

```
  $ curl http://127.0.0.1:8081/users
  []
  $ curl -H "Content-Type: application/json" -X POST -d '{"name":"John"}' http://127.0.0.1:8081/users
  {"name":"John"}
  $ curl http://127.0.0.1:8081/users
  [{"name":"John"}]
  $ curl -H "Content-Type: application/json" -X POST -d '{"name":"John"}' http://127.0.0.1:8081/users
  {"name":"John"}
  $ curl http://127.0.0.1:8081/users
  [{"name":"John"}]
  $ curl -H "Content-Type: application/json" -X POST -d '{"name":"Marc"}' http://127.0.0.1:8081/users
  {"name":"Marc"}
  $ curl http://127.0.0.1:8081/users
  [{"name":"John"},{"name":"Marc"}]
```
