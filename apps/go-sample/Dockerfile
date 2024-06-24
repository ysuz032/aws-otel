# Build stage
FROM golang:1.22.4
WORKDIR /app
RUN go install github.com/air-verse/air@latest

COPY src /app

# Expose port 8080 to the outside world
EXPOSE 8080

# Expose port 6060 for pprof
EXPOSE 6060

# Command to run the executable
CMD ["air"]