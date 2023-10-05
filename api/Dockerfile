# Use the official Golang image to create a build artifact.
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Set necessary environment variables needed for our build.
ENV GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -o main server.go

# Use a lightweight alpine image.
FROM alpine:latest

# Copy the binary from builder to this new stage.
COPY --from=builder /app/main /main

# Command to run the application when the docker container starts.
ENTRYPOINT ["/main"]
