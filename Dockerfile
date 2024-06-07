# Use the official Golang image as a base image
FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR /GO-01

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set the working directory to cmd
WORKDIR /GO-01/cmd

# Build the Go app
RUN go build -o /GO-01/main .

# Set the working directory back to /app
WORKDIR /GO-01

# Command to run the executable
CMD ["./main"]
