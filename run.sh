#!/bin/bash
DAY="$1"
INPUT="$2"
go build -o "build/$DAY" "src/$DAY.go" "src/util.go" "src/counter.go" && "build/$DAY" "inputs/$DAY.$INPUT"
