FROM golang:1.16beta1-alpine3.12

ENV GO111MODULE=on

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 80

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

ENTRYPOINT ["/app/main"]