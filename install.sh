#!/bin/bash

# Run go build command
go build -o boiler-go $(pwd)/cmd

# Move generated boiler-go builded binary to GOPATH/bin
mv boiler-go $GOPATH/bin