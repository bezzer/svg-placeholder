#!/bin/bash
# Compile a static binary
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
# Re-build docker container
docker build -t bezzer/svg-placeholder .