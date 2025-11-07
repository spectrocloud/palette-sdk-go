#!/bin/bash

set -e

if [ -z "$1" ] || [ -z "$2" ]; then
    echo "Usage: generate.sh <swagger_binary_path> <controller_gen_binary_path> <current_dir_path>"
    exit
fi

swagger=$1
controllerGen=$2
currentDir=$3

service="palette-api"
serviceDir="${currentDir}"
apiDir="${serviceDir}"
templateDir="${apiDir}/swagger/templates"

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

# Generate DeepCopy methods for all models
echo "Generating DeepCopy methods..."

# Step 1: Add markers to struct types (not interface{} aliases)
echo "  Adding generation markers to struct types..."
go run scripts/add_deepcopy_markers.go models/

# Step 2: Run controller-gen (it will only generate for marked types)
echo "  Running controller-gen..."
$controllerGen object paths=./models/... 2>&1 | grep -v "invalid field type: interface{}" || true

# Step 3: Fix empty loops in generated code
echo "  Fixing empty loops in generated code..."
bash scripts/fix_empty_loops.sh

# Step 4: Clean up markers from source files (keep generated file)
echo "  Cleaning up markers..."
bash scripts/remove_deepcopy_markers.sh models/
