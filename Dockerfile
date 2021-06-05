FROM golang:1.10.5-alpine3.8

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o exec

EXPOSE 8000

ENTRYPOINT ["app/exec"]

