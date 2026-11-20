package apiutil

import (
	"encoding/json"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// MarshalObjectMetaInputPatch serializes metadata for PATCH requests.
// Non-nil empty Labels/Annotations maps must appear as {} so the API can clear them.
// Nil maps are omitted (leave existing values unchanged).
func MarshalObjectMetaInputPatch(meta *models.V1ObjectMetaInputEntity) ([]byte, error) {
	if meta == nil {
		return []byte("null"), nil
	}

	out := make(map[string]interface{})
	if meta.Name != "" {
		out["name"] = meta.Name
	}
	if meta.Labels != nil {
		out["labels"] = meta.Labels
	}
	if meta.Annotations != nil {
		out["annotations"] = meta.Annotations
	}
	return json.Marshal(out)
}

// MarshalProfileMetaPatchBody serializes a cluster profile metadata patch body.
func MarshalProfileMetaPatchBody(body *models.V1ProfileMetaEntity) ([]byte, error) {
	if body == nil {
		return []byte("null"), nil
	}

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}

	if body.Metadata != nil {
		metaBytes, err := MarshalObjectMetaInputPatch(body.Metadata)
		if err != nil {
			return nil, err
		}
		payload["metadata"] = metaBytes
	}

	return json.Marshal(payload)
}
