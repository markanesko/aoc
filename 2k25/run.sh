#!/bin/bash

day=$1

if [ -z "$day" ]; then
  echo "Usage: $0 <day>"
  exit 1
fi

dir="day-$day"

if [ ! -d "$dir" ]; then
  echo "Directory '$dir' does not exist."
  exit 1
fi

go run "$dir/main.go" "$dir/input"

