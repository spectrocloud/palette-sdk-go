// Package workflows provides high-level, agentic workflow tools that
// orchestrate multiple service calls into coherent operations. Each workflow
// follows a two-phase pattern (plan/apply) suitable for LLM-driven agents.
package workflows

import (
	"encoding/json"
	"fmt"
)

// Phase represents the execution phase of a workflow.
type Phase string

const (
	// PhasePlan returns the schema and requirements for the operation
	// without making any changes.
	PhasePlan Phase = "plan"

	// PhaseApply validates input and executes the operation.
	PhaseApply Phase = "apply"
)

// Result is the structured response from a workflow execution.
type Result struct {
	// Status is "ok" or "error".
	Status string `json:"status"`

	// Message is a human/agent-readable summary of what happened.
	Message string `json:"message"`

	// Data holds structured output (UIDs, schemas, resource details).
	Data map[string]any `json:"data,omitempty"`
}

// OK creates a successful result.
func OK(message string, data map[string]any) Result {
	return Result{Status: "ok", Message: message, Data: data}
}

// Errorf creates an error result.
func Errorf(format string, args ...any) Result {
	return Result{Status: "error", Message: fmt.Sprintf(format, args...)}
}

// String returns the JSON representation of the result for agent consumption.
func (r Result) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

const defaultWorkerPoolName = "worker-pool"

// FieldSpec describes a single input field for the plan phase.
type FieldSpec struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Default     any    `json:"default,omitempty"`
	Enum        []any  `json:"enum,omitempty"`
}
