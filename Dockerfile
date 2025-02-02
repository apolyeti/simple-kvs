# Use the official Golang image as a base
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN apk --no-cache add curl

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
EXPOSE 50051

CMD ["./main"]
