# base go image
FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailerApp ./cmd/api

RUN chmod +x /app/mailerApp

#build a tiny dokcer image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/mailerApp /app
COPY templates /templates

CMD [ "/app/mailerApp" ]