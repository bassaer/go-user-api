FROM golang

ADD . /go/src/github.com/bassaer/go-user-api
WORKDIR /go/src/gihub.com/bassaer/go-user-api

RUN go install /go/src/github.com/bassaer/go-user-api

ENTRYPOINT /go/bin/go-user-api

EXPOSE 8080
