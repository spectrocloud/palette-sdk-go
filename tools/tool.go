// Package tools provides a registry for exposing workflows as agent-callable
// tools. Each tool has a name, description, and a Run function that accepts
// JSON input and returns a string result.
package tools

import (
	"encoding/json"
	"fmt"
)

// Tool represents a single callable tool for an agentic framework.
type Tool struct {
	// Name is the unique identifier agents use to invoke this tool.
	Name string

	// Description tells the agent when and why to use this tool.
	Description string

	// Run executes the tool with JSON input and returns a string result.
	Run func(input json.RawMessage) string
}

// Registry holds all registered tools.
type Registry struct {
	tools map[string]Tool
}

// NewRegistry creates an empty tool registry.
func NewRegistry() *Registry {
	return &Registry{tools: make(map[string]Tool)}
}

// Register adds a tool to the registry. Panics on duplicate names.
func (r *Registry) Register(t Tool) {
	if _, exists := r.tools[t.Name]; exists {
		panic(fmt.Sprintf("tools: duplicate tool name %q", t.Name))
	}
	r.tools[t.Name] = t
}

// Get returns a tool by name, or nil if not found.
func (r *Registry) Get(name string) *Tool {
	t, ok := r.tools[name]
	if !ok {
		return nil
	}
	return &t
}

// List returns all registered tools.
func (r *Registry) List() []Tool {
	out := make([]Tool, 0, len(r.tools))
	for _, t := range r.tools {
		out = append(out, t)
	}
	return out
}

// Invoke calls a tool by name with the given JSON input.
func (r *Registry) Invoke(name string, input json.RawMessage) (string, error) {
	t := r.Get(name)
	if t == nil {
		return "", fmt.Errorf("tool %q not found", name)
	}
	return t.Run(input), nil
}
