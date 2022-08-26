FROM golang:1.18 as builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
#RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o api-afip-wsfe
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine
RUN apk add --no-cache \
        openssl \
        ca-certificates \
        bash
WORKDIR /app
COPY --from=builder /app/api-afip-wsfe /app/
CMD ["/app/api-afip-wsfe"]
