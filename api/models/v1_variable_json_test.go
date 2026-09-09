package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV1Variable_UnmarshalJSON_DefaultValueScalars(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "numeric defaultValue",
			data: `{"name":"gatewayHttpNodePort","defaultValue":30080}`,
			want: "30080",
		},
		{
			name: "boolean defaultValue",
			data: `{"name":"enabled","defaultValue":true}`,
			want: "true",
		},
		{
			name: "string defaultValue",
			data: `{"name":"replicas","defaultValue":"2"}`,
			want: "2",
		},
		{
			name: "decimal defaultValue",
			data: `{"name":"ratio","defaultValue":1.5}`,
			want: "1.5",
		},
		{
			name: "large integer defaultValue preserves precision",
			data: `{"name":"bigInt","defaultValue":9007199254740993}`,
			want: "9007199254740993",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var variable V1Variable
			require.NoError(t, json.Unmarshal([]byte(tt.data), &variable))
			assert.Equal(t, tt.want, variable.DefaultValue)
		})
	}
}

func TestV1ClusterProfileImportEntity_UnmarshalJSON_NumericVariableDefault(t *testing.T) {
	data := `{
		"metadata": {"name": "edge-airgap-base"},
		"spec": {
			"variables": [
				{"name": "gatewayHttpNodePort", "defaultValue": 30080},
				{"name": "gatewayHttpsNodePort", "defaultValue": 30443}
			]
		}
	}`

	var profile V1ClusterProfileImportEntity
	require.NoError(t, json.Unmarshal([]byte(data), &profile))
	require.NotNil(t, profile.Spec)
	require.Len(t, profile.Spec.Variables, 2)
	assert.Equal(t, "30080", profile.Spec.Variables[0].DefaultValue)
	assert.Equal(t, "30443", profile.Spec.Variables[1].DefaultValue)
}
