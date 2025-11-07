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

# Remove lines containing the kubebuilder:object:generate marker
find "$DIR" -name "*.go" ! -name "zz_generated*" -exec sed -i.bak '/^\/\/ +kubebuilder:object:generate=true$/d' {} \;

# Clean up backup files
find "$DIR" -name "*.go.bak" -delete

echo "DeepCopy markers removed successfully!"

