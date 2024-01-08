FROM golang:1.21-alpine AS builder

WORKDIR /target

COPY . .

RUN go mod download

RUN GOOS=darwin GOARCH=arm64 go build -o quote_c ./cmd/client

FROM alpine:latest

COPY --from=builder /target/quote_c /

ENTRYPOINT ["/quote_c"]
