#!/bin/bash

set -e

if [ -z "$1" ]; then
    echo "Usage: generate.sh <current_dir_path>"
    exit
fi

currentDir=$1
service="palette-api"
serviceDir="${currentDir}"
apiDir="${serviceDir}"
commonDir="${apiDir}/common"
templateDir="${apiDir}swagger/templates"

# Fetch the latest hapi spec
rm -rf hapi && git clone https://github.com/spectrocloud/hapi
(
    cd hapi
    bash generate_hubble_spec.sh
    go run api/main.go
    cp gen/docs-spec/palette-apis-spec.json ../spec/
    rm -rf hapi
)

# Flatten it
for i in {1..10}; do
    swagger flatten spec/palette-apis-spec.json --with-flatten=remove-unused --format=json -o spec/palette-apis-spec.json
done

# Generate swagger client & models
swagger mixin spec/palette-apis-spec.json spec/error_v1.json -o spec/palette.json
swagger generate model -t $apiDir -f spec/palette.json
swagger generate client \
    --skip-models \
    --existing-models=github.com/spectrocloud/palette-api-go/models \
    --template-dir="$templateDir" \
    -A "$service" -t "$serviceDir" \
    -f spec/palette.json
