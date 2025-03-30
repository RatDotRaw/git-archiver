#!/bin/bash

# Set environment variables for Windows 64-bit
export GOOS=windows
export GOARCH=amd64

# Build the project with optimizations and the .exe extension
go build git-archiver cmd/git-archiver/main.go -o ./builds/git-archiver_linux_amd64 

# Verify the build
echo "Built ./builds/git-archiver_windows_amd64.exe"