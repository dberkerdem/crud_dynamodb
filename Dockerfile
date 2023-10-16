# Builder
FROM golang:1.21.1 AS builder

WORKDIR /app

COPY ./api .

# Download dependencies
RUN go mod tidy

# Build the Go app
ENV CGO_ENABLED 0
RUN go build -o service-exec main.go

# Certificate downloader
FROM alpine:latest AS certs

# Download Amazon Root CA certificates
RUN apk --no-cache add ca-certificates && \
    wget -O /etc/ssl/certs/AmazonRootCA1.pem https://www.amazontrust.com/repository/AmazonRootCA1.pem && \
    wget -O /etc/ssl/certs/AmazonRootCA2.pem https://www.amazontrust.com/repository/AmazonRootCA2.pem && \
    wget -O /etc/ssl/certs/AmazonRootCA3.pem https://www.amazontrust.com/repository/AmazonRootCA3.pem && \
    wget -O /etc/ssl/certs/AmazonRootCA4.pem https://www.amazontrust.com/repository/AmazonRootCA4.pem

# Final
FROM scratch

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/service-exec .
COPY --from=certs /etc/ssl/certs/ /etc/ssl/certs/
EXPOSE 8080

# Command to run the executable
CMD ["./service-exec"]
