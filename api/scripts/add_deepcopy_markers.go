// Package main adds deepcopy generation markers to viable struct types in Go files.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Add deepcopy generation markers to struct types in Go files
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: add_deepcopy_markers <directory>")
		os.Exit(1)
	}

	root, err := os.OpenRoot(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if closeErr := root.Close(); closeErr != nil {
			fmt.Printf("Error closing root: %v\n", closeErr)
		}
	}()

	rootPath := root.Name()

	err = filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		rel, err := filepath.Rel(rootPath, path)
		if err != nil || strings.HasPrefix(rel, "..") {
			return fmt.Errorf("path escapes root: %s", path)
		}

		// Skip generated deepcopy files
		if !d.IsDir() && strings.HasSuffix(path, ".go") &&
			!strings.Contains(path, "zz_generated") {
			return processFile(root, rel)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully added DeepCopy markers to eligible types")
}

func processFile(root *os.Root, relPath string) error {
	content, err := root.ReadFile(relPath)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, relPath, content, parser.ParseComments)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	modifications := make(map[int]string) // line number -> marker to add

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Only process struct types (skip interface{} aliases, primitive type aliases, etc.)
			structType, isStruct := typeSpec.Type.(*ast.StructType)
			if !isStruct || structType == nil {
				continue
			}

			// Skip empty structs
			if structType.Fields == nil || len(structType.Fields.List) == 0 {
				continue
			}

			// Check if marker already exists
			pos := fset.Position(typeSpec.Pos())
			lineNum := pos.Line - 1 // 0-indexed

			hasMarker := false
			// Check a few lines above for existing marker
			for i := max(0, lineNum-5); i < lineNum; i++ {
				if strings.Contains(lines[i], "+kubebuilder:object:generate=true") {
					hasMarker = true
					break
				}
			}

			if !hasMarker {
				// Find the line with "type TypeName struct {"
				typeLineNum := lineNum
				for i := lineNum; i >= 0 && i > lineNum-10; i-- {
					if strings.Contains(lines[i], "type "+typeSpec.Name.Name) {
						typeLineNum = i
						break
					}
				}

				modifications[typeLineNum] = "// +kubebuilder:object:generate=true"
			}
		}
	}

	// Apply modifications (insert markers)
	if len(modifications) > 0 {
		newLines := make([]string, 0, len(lines)+len(modifications))

		for i, line := range lines {
			if marker, ok := modifications[i]; ok {
				newLines = append(newLines, marker)
			}
			newLines = append(newLines, line)
		}

		newContent := strings.Join(newLines, "\n")
		if err := root.WriteFile(relPath, []byte(newContent), 0o600); err != nil {
			return fmt.Errorf("failed to write %s: %w", relPath, err)
		}
	}

	return nil
}
