#!/bin/bash
# Fix empty for loops in generated DeepCopy code
# controller-gen creates empty loops for map[string]interface{} which causes unused variable errors

set -e

DEEPCOPY_FILE="models/zz_generated.deepcopy.go"

if [ ! -f "$DEEPCOPY_FILE" ]; then
    echo "Error: $DEEPCOPY_FILE not found"
    exit 1
fi

echo "Fixing empty loop bodies for map[string]interface{}..."

# Fix: for key, val := range *in { } -> add (*out)[key] = val inside the loop
awk '
/for key, val := range \*in \{/ {
    print
    getline
    if ($0 ~ /^[[:space:]]*\}$/) {
        print "\t\t\t(*out)[key] = val"
    }
    print
    next
}
{ print }
' "$DEEPCOPY_FILE" > "${DEEPCOPY_FILE}.tmp"

mv "${DEEPCOPY_FILE}.tmp" "$DEEPCOPY_FILE"

echo "Fixed empty loops in $DEEPCOPY_FILE"

# Also fix V1Time DeepCopyInto which tries to access unexported 'loc' field
echo "Fixing V1Time DeepCopyInto to avoid unexported field..."
perl -i -0777 -pe 's/(func \(in \*V1Time\) DeepCopyInto\(out \*V1Time\) \{)\n\t\*out = \*in\n\tif in\.loc != nil \{[^\}]*\}/$1\n\t*out = *in/gs' "$DEEPCOPY_FILE"

# Remove unused timex import if V1Time no longer uses it
echo "Cleaning up unused imports..."
perl -i -0777 -pe 's/\n\ttimex "time"//' "$DEEPCOPY_FILE"

echo "All fixes applied to $DEEPCOPY_FILE"

