#!/bin/bash
DAY="$1"
shift;
go build -o "build/$DAY" "src/$DAY.go" "src/util.go" && "build/$DAY" "$@"
