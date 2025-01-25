
FROM alpine:latest

COPY --from=builder /app/bin/backend /usr/local/bin/backend
# Assuming port is 8080, adjust according to your app.
EXPOSE 8080 

CMD ["backend"] # This assumes the executable's name from Makefile is 'backend'. Adjust accordingly.