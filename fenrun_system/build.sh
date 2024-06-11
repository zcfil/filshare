#!/usr/bin/env bash
/usr/local/go/bin/go mod tidy
git pull
/usr/local/go/bin/go build -o kdsystem