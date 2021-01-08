#!/bin/bash
DAY=`printf "day%02d" $1`

if [ ! -f "src/$DAY.go" ]; then
    cp day.go "src/$DAY.go"
fi

touch "inputs/$DAY.input"
touch "inputs/$DAY.example"

echo "./run.sh $DAY example"
