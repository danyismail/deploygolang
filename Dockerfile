FROM golang:1.10.5-alpine3.8

WORKDIR /app

COPY main.go .

RUN go build -o binary

EXPOSE 8000

CMD [“./main”]

