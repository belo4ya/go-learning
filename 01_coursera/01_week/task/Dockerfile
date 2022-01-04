FROM golang:1.17.5

WORKDIR /usr/src/go

COPY . /usr/src/go

RUN go mod init task

RUN go test -v