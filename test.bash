#!/usr/bin/env bash

set -e

for t in $(go list ./... | grep -v vendor); do
  go test -cover ${t}
done
