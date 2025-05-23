# syntax=docker/dockerfile:1

# Stage 1: Run generator
FROM golang:1.22 AS generator

# Install make (should be included in golang images but just in case)
RUN apt-get update && apt-get install -y make

# Set working directory
WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY 004-improved-generation-for-annotations/ ./004-improved-generation-for-annotations/

# Generate routes
WORKDIR /app/004-improved-generation-for-annotations
RUN make generate-routes

# Build the binary statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/myapp ./cmd/main.go

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

# Stage 2: Runtime
FROM scratch
WORKDIR /app
COPY --from=generator /app/myapp .
COPY --from=generator /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


CMD ["./myapp"]
