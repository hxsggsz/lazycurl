BINARY_NAME=lazycurl

# Default target
all: generate build

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) ./cmd/main.go

# Forward args after 'run' to your binary
runargs := $(filter-out run,$(MAKECMDGOALS))
run:
	./bin/$(BINARY_NAME) $(runargs)
