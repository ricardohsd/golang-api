This is an example of http server written in Golang.

Build the docker image:

```
   $ docker build -t golang-http-example .
```

Run the container:
```
   $ docker run --rm -it -p 8081:8081 golang-http-example
```

In a new terminal window, execute:

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
