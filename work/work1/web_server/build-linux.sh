#!/bin/bash
export GOPATH=`pwd`
export GOOS=linux
export GOARCH=amd64

go build

read -p "Press any key to continue."


