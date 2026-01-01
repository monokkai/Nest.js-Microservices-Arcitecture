FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o metrics-service

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/metrics-service .

EXPOSE 9091
CMD ["./metrics-service"]
