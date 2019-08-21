FROM golang

ADD . /go/src/github.com/bassaer/go-user-api
WORKDIR /go/src/gihub.com/bassaer/go-user-api

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/google/uuid
RUN go install /go/src/github.com/bassaer/go-user-api/cmd/userapi

ENTRYPOINT /go/bin/userapi

EXPOSE 8080
