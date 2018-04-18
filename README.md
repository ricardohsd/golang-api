This is an example of http server written in Golang.

Build the docker image:

```
   $ docker build -t golang-http-example .
```

Run the container:
```
   $ docker run --rm -it -p 8081:8081 golang-http-example
```

In a new terminal window, execute (with httpie):

```
  $ http -a username:password http://localhost:8081/users
  HTTP/1.1 200 OK
  Content-Length: 3
  Content-Type: application/json;charset=UTF-8
  Date: Wed, 18 Apr 2018 21:14:12 GMT

  []
  $ http -v -a username:password -f POST http://localhost:8081/users Content-Type:application/json <<<'{"name":"john"}'
  HTTP/1.1 201 Created
  Content-Length: 16
  Content-Type: application/json;charset=UTF-8
  Date: Wed, 18 Apr 2018 21:12:54 GMT

  {
      "name": "john"
  }
  $ http -a username:password http://localhost:8081/users
  HTTP/1.1 200 OK
  Content-Length: 18
  Content-Type: application/json;charset=UTF-8
  Date: Wed, 18 Apr 2018 21:13:01 GMT

  [
      {
          "name": "john"
      }
  ]
```
