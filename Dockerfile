FROM golang:1.8 as build-stage
WORKDIR /go/src/github.com/ricardohsd/golang-http-example
COPY *.go /go/src/github.com/ricardohsd/golang-http-example/
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-server .

FROM alpine:latest
COPY --from=build-stage /go/src/github.com/ricardohsd/golang-http-example/go-server /bin
CMD go-server
