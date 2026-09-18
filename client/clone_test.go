package client

// Clone() must preserve rootCAs and baseCtx — silent regression before this fix.
// Same-package test so we can read unexported fields directly without exposing accessors.

import (
	"context"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClone_PreservesRootCAs(t *testing.T) {
	pool := x509.NewCertPool()
	// Empty pool is fine — pointer identity is what we assert.

	orig := New(
		WithPaletteURI("api.example.com"),
		WithAPIKey("k"),
		WithRootCAs(pool),
	)
	require.Same(t, pool, orig.rootCAs, "sanity: original client holds the given pool")

	cloned := orig.Clone()
	assert.Same(t, pool, cloned.rootCAs,
		"Clone must propagate rootCAs — airgap / custom-CA callers regress otherwise")
}

func TestClone_PreservesBaseCtx(t *testing.T) {
	type ctxKey struct{}
	ctx := context.WithValue(context.Background(), ctxKey{}, "sentinel")

	orig := New(
		WithPaletteURI("api.example.com"),
		WithAPIKey("k"),
		WithContext(ctx),
	)
	require.Equal(t, "sentinel", orig.baseCtx.Value(ctxKey{}), "sanity: original baseCtx carries sentinel")

	cloned := orig.Clone()
	assert.Equal(t, "sentinel", cloned.baseCtx.Value(ctxKey{}),
		"Clone must propagate baseCtx — cancellation / deadlines / tracing regress otherwise")
}

func TestClone_PreservesProjectScopeAfterBaseCtx(t *testing.T) {
	// Ordering guard: WithContext runs before WithScopeProject in Clone's opts.
	// If order flipped, WithScopeProject would compute v.ctx off the OLD baseCtx
	// and WithContext would then overwrite v.ctx back to plain baseCtx,
	// erasing the ProjectUid header.
	orig := New(
		WithPaletteURI("api.example.com"),
		WithAPIKey("k"),
		WithScopeProject("proj-uid-123"),
	)
	require.Equal(t, "proj-uid-123", orig.projectUID)

	cloned := orig.Clone()
	assert.Equal(t, "proj-uid-123", cloned.projectUID, "Clone must preserve projectUID")
	assert.NotEqual(t, orig.baseCtx, cloned.ctx,
		"Clone's scope-project ctx must differ from baseCtx (carries ProjectUid header)")
}
