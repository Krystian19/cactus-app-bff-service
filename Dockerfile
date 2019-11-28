FROM golang:1.13.2-alpine3.10

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .
EXPOSE 3000

RUN apk add make git gcc libc-dev curl

RUN go get github.com/unknwon/bra

CMD bra run