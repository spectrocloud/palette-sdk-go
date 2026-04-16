// Package shared provides cross-cutting types used by all service packages.
package shared

import "github.com/spectrocloud/palette-sdk-go/client"

// Scope represents the operational scope for API operations.
type Scope string

const (
	// ScopeTenant indicates tenant-level scope.
	ScopeTenant Scope = "tenant"
	// ScopeProject indicates project-level scope.
	ScopeProject Scope = "project"
)

// String returns the string representation of the scope.
func (s Scope) String() string {
	return string(s)
}

// ServiceConfig holds common dependencies for all services.
type ServiceConfig struct {
	Client *client.V1Client
}
