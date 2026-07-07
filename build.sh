#!/bin/bash

SRC_DIR="src/meeting_tracker_rpb_go"
BINARY_NAME="meeting_tracker"

echo "Building for Raspberry Pi Zero (linux/arm6)..."
cd "$SRC_DIR" && GOOS=linux GOARCH=arm GOARM=6 go build -o "../../$BINARY_NAME" main.go && cd ../..

if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Build successful: $BINARY_NAME"
