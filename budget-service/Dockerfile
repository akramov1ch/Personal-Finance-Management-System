FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o budget-service ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/budget-service /app/

WORKDIR /app

COPY ./.env /app/.env

EXPOSE 5050

ENTRYPOINT ["/app/budget-service"]