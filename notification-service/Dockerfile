FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o notifaction-service ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/notifaction-service /app/

WORKDIR /app

COPY ./config.yaml /app/config.yaml

EXPOSE 8080
EXPOSE 2020

ENTRYPOINT ["/app/notifaction-service"]