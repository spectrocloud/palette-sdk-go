#!/bin/bash

set -e

if [ -z "$1" ]; then
    echo "Usage: generate.sh <swagger_binary_path> <current_dir_path>"
    exit
fi

swagger=$1
currentDir=$2

service="palette-api"
serviceDir="${currentDir}"
apiDir="${serviceDir}"
templateDir="${apiDir}/swagger/templates"

# Fetch the latest hapi spec
rm -rf hapi && git clone https://github.com/spectrocloud/hapi
(
    cd hapi
    git checkout cloudstack
    git pull
    bash generate_hubble_spec.sh
    go run api/main.go
    cp gen/docs-spec/palette-apis-spec.json ../spec/
    rm -rf hapi
)

# Flatten it
$swagger flatten spec/palette-apis-spec.json --with-flatten=remove-unused --format=json -o spec/palette-apis-spec.json

# Generate swagger client & models
$swagger mixin spec/palette-apis-spec.json spec/error_v1.json -o spec/palette.json
$swagger generate model -t "$apiDir" -f spec/palette.json
$swagger generate client \
    --skip-models \
    --existing-models=github.com/spectrocloud/palette-sdk-go/api/models \
    --template-dir="$templateDir" \
    -A "$service" -t "$serviceDir" \
    -f spec/palette.json
