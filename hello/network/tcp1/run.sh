#!/usr/bin/env bash

# from :
# https://github.com/aaronbieber/tcp-server-client-go

go run ./server/main.go -port 5555

go run ./client/main.go -host localhost -port 5555