FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-gateway ./cmd

FROM gcr.io/distroless/base

COPY --from=builder /app/api-gateway /app/

WORKDIR /app

COPY config.yaml /app/config.yaml

COPY ./tls /app/tls

EXPOSE 4040

ENTRYPOINT ["/app/api-gateway"]