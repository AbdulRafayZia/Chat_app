# Use the official Golang image as the base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Use the official PostgreSQL image as the base image for the second stage


# Set the working directory inside the container

EXPOSE 8080
# Command to run the application
CMD ["./main"]
