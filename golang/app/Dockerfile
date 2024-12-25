# Step 1: Use a lightweight base image with Go installed
FROM golang:alpine AS builder

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy the Go module files to the working directory
COPY ./go.mod ./go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the rest of the application code into the working directory
COPY . .

# Step 6: Build the Go application as a statically linked binary
RUN go build -o auth-api .

# Step 7: Use a minimal Linux image to reduce final container size
FROM alpine:latest

# Step 8: Set the working directory for the runtime image
WORKDIR /app

# Step 9: Install postgresql-client in the Alpine image
RUN apk add --no-cache postgresql-client

# Step 9: Copy the binary from the builder stage to the runtime image
COPY --from=builder /app/auth-api .

# Step 10: Expose the port your API uses (e.g., 8000)
EXPOSE 8000

# Step 11: Set the default command to run the binary
CMD ["./auth-api"]
