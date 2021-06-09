FROM golang:1.13

ENV PORT=8000
ENV INSTANCE_ID=is_running

WORKDIR /app

COPY . .

RUN go build -o binary

EXPOSE 8080

ENTRYPOINT ["/app/binary"]

