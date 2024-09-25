# Build stage
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Install Air for hot reload in development
RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

# Build the Go app for production
RUN go build -o main ./cmd

# Development stage
FROM golang:1.23.1-alpine AS development

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN mkdir -p tmp
ENV PATH="/root/go/bin:${PATH}"

# Command to run the app using Air for hot reload
CMD ["air", "-c", ".air.toml"]

# Production stage
FROM alpine:latest AS production

WORKDIR /root/

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/ui ./web/static

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
