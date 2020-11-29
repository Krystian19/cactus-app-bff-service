FROM golang:1.13.2-alpine3.10

ARG ACCESS_TOKEN_USR
ARG ACCESS_TOKEN_PWD

ENV GO111MODULE=on
ENV GOPRIVATE=github.com/Krystian19

RUN apk add make git gcc libc-dev curl

# Create a netrc file using the credentials specified using --build-arg
RUN printf "machine github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n\
    \n\
    machine api.github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

RUN go get github.com/unknwon/bra

WORKDIR /go/src/app
COPY . .
EXPOSE 3000

CMD bra run