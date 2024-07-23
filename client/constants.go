package client

const (
	authTokenInput = "header"
	authAPIKey     = "ApiKey"
	authJwt        = "Authorization"

	// OverlordUID is an annotation key used to link a cloud accounts with a private cloud gateway.
	OverlordUID = "overlordUID"

	// Scope is the key for the scope in the Palette authorization context
	// One of: "project", "system", "tenant"
	Scope = "scope"
)
