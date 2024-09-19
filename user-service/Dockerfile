FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN mkdir -p /app/logs

COPY . .

RUN go build -o user-service ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/user-service /app/
COPY --from=builder /app/migrations /app/migrations
COPY ./.env /app/.env

WORKDIR /app

EXPOSE 7070

ENTRYPOINT ["/app/user-service"]
