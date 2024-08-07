package herr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spectrocloud/palette-sdk-go/api/apiutil/transport"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name: "ResourceNotFound",
			err: &transport.TransportError{
				Payload: &models.V1Error{
					Code: "ResourceNotFound",
				},
			},
			expected: true,
		},
		{
			name: "UnknownError",
			err: &transport.TransportError{
				Payload: &models.V1Error{
					Code: "UnknownError",
				},
			},
			expected: false,
		},
		{
			name:     "Error",
			err:      fmt.Errorf("error"),
			expected: false,
		},
		{
			name:     "Nil",
			err:      nil,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, IsNotFound(test.err))
		})
	}
}
