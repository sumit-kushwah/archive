# Use an official Go runtime as the base image
FROM golang:1.20-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the remaining source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the necessary port (if applicable)
EXPOSE 4000

# Define the command to run your application
CMD ["./main"]
