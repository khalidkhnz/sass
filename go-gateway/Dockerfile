
# Use an official Golang runtime as a parent image
FROM golang:1.23.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go files and set them to be used in the build process
COPY . .

# Download all dependencies, including tools(if any) needed for running your app or tests. 
RUN go mod download && apk add --no-cache git # This line installs Git if you need it for fetching packages from private repos, etc.

# Build your Go application using the makefile's build command as a template
RUN make build

FROM alpine:latest

COPY --from=builder /app/bin/backend /usr/local/bin/backend
# Assuming port is 8080, adjust according to your app.
EXPOSE 8080 

CMD ["backend"] # This assumes the executable's name from Makefile is 'backend'. Adjust accordingly.