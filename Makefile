# Makefile for your Go project

# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOINSTALL := $(GOCMD) install
BINARY_NAME := myapp
MAIN_FILE := ./cmd/main.go

# Targets
all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)
	./$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

install:
	$(GOINSTALL) ./...

.PHONY: all build clean run test install
