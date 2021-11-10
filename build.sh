#!/bin/sh

cd web
yarn
yarn build

cd ..
echo "Start building backend"
go mod download
mkdir -p dist
CGO_ENABLED=0 GOOS=linux go build --tags "release" -ldflags="-w -s" -o dist/urlshortener ./main.go
echo "Finish building backend"

# Start service by following command
# cd dist
# CONFIGFILE=../config/config.json WEBPATH=../web/dist ./urlshortener
