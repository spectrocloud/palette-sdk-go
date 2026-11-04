package apiutil

import (
	"encoding/json"
	"testing"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalObjectMetaInputPatch_omitsNilLabels(t *testing.T) {
	raw, err := MarshalObjectMetaInputPatch(&models.V1ObjectMetaInputEntity{
		Name: "profile-a",
	})
	require.NoError(t, err)
	assert.NotContains(t, string(raw), `"labels"`)
}

func TestMarshalObjectMetaInputPatch_includesEmptyLabels(t *testing.T) {
	raw, err := MarshalObjectMetaInputPatch(&models.V1ObjectMetaInputEntity{
		Name:   "profile-a",
		Labels: map[string]string{},
	})
	require.NoError(t, err)

	var payload map[string]interface{}
	require.NoError(t, json.Unmarshal(raw, &payload))
	labels, ok := payload["labels"].(map[string]interface{})
	require.True(t, ok)
	assert.Empty(t, labels)
}

func TestMarshalProfileMetaPatchBody_clearsTags(t *testing.T) {
	raw, err := MarshalProfileMetaPatchBody(&models.V1ProfileMetaEntity{
		Metadata: &models.V1ObjectMetaInputEntity{
			Name: "profile-a",
			Annotations: map[string]string{
				"description": "test",
			},
			Labels: map[string]string{},
		},
		Spec: &models.V1ClusterProfileSpecEntity{
			Version: "1.0.0",
		},
	})
	require.NoError(t, err)

	var payload map[string]interface{}
	require.NoError(t, json.Unmarshal(raw, &payload))

	meta, ok := payload["metadata"].(map[string]interface{})
	require.True(t, ok)
	labels, ok := meta["labels"].(map[string]interface{})
	require.True(t, ok)
	assert.Empty(t, labels)
	assert.Contains(t, meta, "annotations")
	assert.Equal(t, "1.0.0", payload["spec"].(map[string]interface{})["version"])
}

func TestMarshalObjectMetaInputPatch_standardJSONOmitsEmptyLabels(t *testing.T) {
	// Documents the bug PLT-2269: default encoding drops empty maps.
	raw, err := json.Marshal(&models.V1ObjectMetaInputEntity{
		Name:   "profile-a",
		Labels: map[string]string{},
	})
	require.NoError(t, err)
	assert.NotContains(t, string(raw), `"labels"`)
}
