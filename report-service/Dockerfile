FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o report-service ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/report-service /app/

WORKDIR /app

COPY ./config.yaml /app/config.yaml

EXPOSE 9090

ENTRYPOINT ["/app/report-service"]
