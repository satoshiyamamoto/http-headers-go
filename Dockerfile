FROM golang:latest

LABEL maintainer="Satoshi Yamamoto <satoshiyamamoto@me.com>"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
