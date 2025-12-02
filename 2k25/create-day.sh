#!/bin/bash

day=$1

if [ -z "$day" ]; then
  echo "Usage: $0 <day>"
  exit 1
fi

dir="day-$day"

if [ -d "$dir" ]; then
  echo "Directory '$dir' already exist."
  exit 1
fi

mkdir $dir

cp ./template.go $dir/main.go


