#!/bin/sh

export GOARCH=amd64

GOOS=linux go build -o blinks-installer-linux-amd64
GOOS=darwin go build -o blinks-installer-darwin-amd64
GOOS=windows go build -o blinks-installer-windows-amd64.exe

