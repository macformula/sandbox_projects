#!/bin/bash

# Set the target architecture and ARM version
GOARCH="arm"
GOARM="7"
GOOS="linux"

# Set the output binary name
OUTPUT_BINARY="testLibSpi1"  # Replace with your desired binary name

cd ../

# Run the Go build command
GO111MODULE=on GOARCH="$GOARCH" GOARM="$GOARM" GOOS="$GOOS" go build -o "$OUTPUT_BINARY" .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build completed successfully. Binary: $OUTPUT_BINARY"
else
  echo "Build failed."
fi

scp $OUTPUT_BINARY user@172.26.187.66:/opt/macfe/bin
