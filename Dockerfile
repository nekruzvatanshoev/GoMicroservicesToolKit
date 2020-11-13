FROM golang

ADD main.go /go/src/CloudNativeGo/main
WORKDIR /go/src/CloudNativeGo/main
RUN go build main.go

ENTRYPOINT /go/bin/CloudNativeGo/main

EXPOSE 8080