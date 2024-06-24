.PHONY: all build run clean

# Build the Docker image
build:
	cd apps && docker-compose build

# Run the Docker container
run:
	cd apps && docker-compose up -d

# Clean up the Docker environment
clean:
	cd apps && docker-compose down
	cd apps && docker-compose rm -f
	cd apps && docker rmi apps-go-sample

# Default target
all: build run