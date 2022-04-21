FROM golang:1.17.1 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main  /app/

CMD ["/app/main"]
