# Use the latest Alpine image as the base
FROM alpine:latest


# Set the working directory
WORKDIR /app

# Copy the compiled binary to the image
COPY ./bin/backend ./backend

# Ensure the binary is executable
RUN chmod +x ./backend

# Expose the port your application runs on
EXPOSE 8083


# Run the application
CMD ["./backend"]
