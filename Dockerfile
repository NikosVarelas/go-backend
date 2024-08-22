# Stage 1: Build the Go application
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first, for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Gin application
RUN go build -o /app/main ./cmd/app

# Stage 2: Build Tailwind CSS
FROM node:20-alpine as tailwind-builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json files first, for dependency caching
COPY package.json package-lock.json ./

# Install Node.js dependencies
RUN npm install

# Copy the rest of the application (including Tailwind config and source files)
COPY . .

# Run the Tailwind CSS build process
RUN npx tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

# Stage 3: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install any necessary dependencies
RUN apk --no-cache add ca-certificates


# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the built Tailwind CSS from the tailwind-builder stage
COPY --from=tailwind-builder /app/static/css/style.min.css /app/static/css/style.css

# Copy the templates directory into the final image
COPY --from=builder /app/templates ./templates

# Expose the port that the application will run on
EXPOSE 3000

# Command to run the application
CMD ["./main"]
