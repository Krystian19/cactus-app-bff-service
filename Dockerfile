FROM golang:1.12.7-alpine3.10

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN apk add make git gcc libc-dev

RUN go get github.com/Unknwon/bra

CMD bra run