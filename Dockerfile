FROM golang:1.22-alpine AS builder

RUN apk add --update --no-cache vips-dev musl-dev git make build-base gcc g++ curl openssl

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux CGO_CFLAGS_ALLOW=-Xpreprocessor CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o service ./cmd/server

FROM alpine:latest

RUN apk add --update --no-cache musl-dev vips curl openssl
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/service /service

USER appuser
CMD ["/service"]
