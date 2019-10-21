FROM golang:1.13.2-alpine3.10

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN apk add make git gcc libc-dev

RUN go get github.com/unknwon/bra

CMD bra run