#!/bin/bash

declare output="./dist/${CODEBUILD_WEBHOOK_TRIGGER}/${CODEBUILD_BUILD_NUMBER}"

echo ${output}

# -tags lambda.norpc
#   reduces binary size by excluding the remote procedure call component
# -ldflags="-s -w"
#   reduces binary size by removing debug libraries
env GOOS="linux" CGO_ENABLED="0" GOARCH="arm64" GOTOOLCHAIN="auto" \
    go build \
    -tags lambda.norpc \
    -ldflags="-s -w" \
    -o ${output}/bootstrap main.go

# copy the config if it exists
FILE="ie2-config.yml"
echo "checking if $FILE exists"

if [ -f "$FILE" ]; then
    echo "found $FILE"
    echo "copying $FILE to output directory: $output"
    cp "$FILE" "$output/$FILE"
else
    echo "file $FILE does NOT exist..."
fi

cd ${output}

chmod +x bootstrap

zip bootstrap.zip bootstrap