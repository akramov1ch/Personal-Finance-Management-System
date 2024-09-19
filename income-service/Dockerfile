FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o income-service ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/income-service /app/

WORKDIR /app

COPY ./.env /app/.env

EXPOSE 6060

ENTRYPOINT ["/app/income-service"]