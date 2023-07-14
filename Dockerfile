FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nat_backend_go .

FROM alpine:3.18.2

WORKDIR /app

COPY --from=builder /app/nat_backend_go .

ENTRYPOINT ["./nat_backend_go"]