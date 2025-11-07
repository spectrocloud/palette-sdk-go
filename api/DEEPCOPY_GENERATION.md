# DeepCopy Method Generation

This directory contains scripts to automatically generate `DeepCopy()` and `DeepCopyInto()` methods for all swagger-generated model structs using Kubernetes `controller-gen`.

## Overview

The DeepCopy methods are generated using the following workflow:

1. **Pre-process**: Add `// +kubebuilder:object:generate=true` markers to all struct types
2. **Generate**: Run `controller-gen` to generate DeepCopy methods 
3. **Post-process**: Fix edge cases (empty loops, unexported fields)
4. **Cleanup**: Remove temporary markers from source files

## Files

- `scripts/add_deepcopy_markers.go` - Go program that adds generation markers to struct types (not interface{} aliases)
- `scripts/remove_deepcopy_markers.sh` - Shell script that removes markers from source files after generation
- `scripts/fix_empty_loops.sh` - Shell script that fixes edge cases in generated code:
  - Empty loops for `map[string]interface{}` types
- `models/zz_generated.deepcopy.go` - Generated DeepCopy methods (auto-generated, do not edit)
- `models/zz_generated_time.deepcopy.go` - Manually maintained DeepCopy methods for V1Time (do not delete)

## Integration with generate.sh

The DeepCopy generation is integrated into the main `generate.sh` script and runs automatically after swagger code generation:

```bash
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
```

## Why This Approach?

### Problem

Swagger-generated models don't have DeepCopy methods, which are useful for:

- Creating independent copies of API objects
- Avoiding accidental mutations
- Testing with modified copies

### Why Not Package-Level Markers?

Adding `// +kubebuilder:object:generate=true` to `doc.go` would generate DeepCopy for ALL types, including:
- `interface{}` type aliases (like `V1PackSummaryStatus`) - causes "invalid receiver" errors
- Type aliases to external types with unexported fields - causes compilation errors

### Solution

Selectively mark only struct types that can have DeepCopy methods:

1. Pre-processing script identifies struct types (not interface{} aliases)
2. Temporarily adds markers to those types
3. controller-gen generates DeepCopy methods
4. Post-processing fixes edge cases
5. Cleanup removes temporary markers

## Edge Cases Handled

### 1. Empty Loops for map[string]interface{}

Controller-gen generates empty loops when it encounters `map[string]interface{}`:
```go
for key, val := range *in {
}  // key and val unused - compilation error
```

Fixed by adding the copy statement:
```go
for key, val := range *in {
    (*out)[key] = val
}
```

### 2. V1Time Type Alias

`V1Time` is a type alias (`type V1Time strfmt.DateTime`), not a struct, so controller-gen won't generate DeepCopy methods for it. However, many structs have `V1Time` fields and need to call `DeepCopyInto()` on them.

**Solution**: Manually maintained `zz_generated_time.deepcopy.go` provides the DeepCopy methods for V1Time. This file:

- Is not touched by `generate.sh`
- Should not be deleted
- Uses simple assignment (`*out = *in`) which is safe for time values

## Usage

### Automatic (Recommended)

Run the main generation script which includes DeepCopy generation:
```bash
make generate
```

### Manual

If you only want to regenerate DeepCopy methods:

```bash
cd api
go run scripts/add_deepcopy_markers.go models/
../bin/controller-gen-v0.19.0 object paths=./models/...
bash scripts/fix_empty_loops.sh
bash scripts/remove_deepcopy_markers.sh models/
```

## Maintenance

When swagger regenerates the models:

1. The `zz_generated.deepcopy.go` file should be regenerated
2. Run `make generate` to update everything
3. Markers are temporary and automatically cleaned up
4. No manual intervention needed unless new edge cases are discovered

## Notes

- Type aliases to `interface{}` (like `V1PackSummaryStatus`, `V1TeamStatus`, `V1Updated`) do NOT get DeepCopy methods - this is intentional and correct
- Controller-gen warnings about "invalid field type: interface{}" are expected and filtered out
- All generation markers are temporary and automatically removed after generation
- **Important**: `models/zz_generated_time.deepcopy.go` is manually maintained and should NOT be deleted or modified during generation

