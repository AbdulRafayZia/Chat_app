# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
