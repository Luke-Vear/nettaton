#!/usr/bin/env bash

set -e

for t in $(go list ./... | grep -v -E 'vendor|sandbox'); do
  go test -cover ${t} 
done
