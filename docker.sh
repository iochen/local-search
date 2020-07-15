#!/usr/bin/env bash

# pass version as $1, cmd_path as $2

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags \"-static\"" -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -v -o ./build/localgen_amd64_linux_static $2
docker build -t localgen:$1 .