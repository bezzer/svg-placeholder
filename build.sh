#!/bin/bash
# Compile a static binary
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

# Build front-end resources (requires gulp)
gulp 

# Re-build docker container
docker build --no-cache=true -t bezzer/svg-placeholder .