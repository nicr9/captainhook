FROM golang:1.9
MAINTAINER Nic Roland <nicroland9@gmail.com>

WORKDIR /go/src/app
COPY *.go /go/src/app/
COPY templates /go/src/app/templates

RUN go get -d -v ./...
RUN go install -v

ENTRYPOINT app
