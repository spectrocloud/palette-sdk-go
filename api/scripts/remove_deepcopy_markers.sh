#!/bin/bash
# Remove deepcopy generation markers from model files
# These were temporarily added for controller-gen but should be cleaned up

set -e

if [ -z "$1" ]; then
    echo "Usage: remove_deepcopy_markers.sh <directory>"
    exit 1
fi

DIR="$1"

echo "Removing DeepCopy markers from $DIR..."

# Remove lines containing the kubebuilder:object:generate=true marker
# Skip v1_time.go as it needs to keep its generate=false marker for external projects
find "$DIR" -name "*.go" ! -name "zz_generated*" ! -name "v1_time.go" -exec sed -i.bak '/^\/\/ +kubebuilder:object:generate=true$/d' {} \;

# Clean up backup files
find "$DIR" -name "*.go.bak" -delete

echo "DeepCopy markers removed successfully!"

