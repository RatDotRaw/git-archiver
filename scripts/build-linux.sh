#!/bin/bash

# Set environment variables for Linux 64-bit
export GOOS=linux
export GOARCH=amd64

# Build the project with optimizations
go build git-archiver cmd/git-archiver/main.go -o ./builds/git-archiver_linux_amd64 

# Verify the build
echo "Built ./builds/git-archiver_linux_amd64"