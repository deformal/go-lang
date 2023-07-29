#!/bin/bash
rm -rf bin
#fileType="*.go"
env GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/main main.go
wait
echo Done
# find . -type f -name "$fileType" | while read -r file; do
#     pf=$(dirname "$file")
#     df=$(basename "$pf")
#     echo $df
#     env GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/$df $file
# done