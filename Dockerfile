# Stage 1: build
FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /rma ./cmd/rma

# Stage 2: runtime
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /rma /app/rma
COPY config/default.yaml /app/config/default.yaml

EXPOSE 9000

ENTRYPOINT ["/app/rma", "/app/config/default.yaml"]
