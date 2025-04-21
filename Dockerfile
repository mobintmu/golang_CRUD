FROM golang:1.23.2-alpine

WORKDIR /app

# Copy all files
COPY . .

COPY /config/config.yaml ./config/config.yaml

# Install build dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd

EXPOSE 8080

CMD ["./main"]