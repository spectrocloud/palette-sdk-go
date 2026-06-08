// Package models provides API request and response types for the Spectro Cloud API.
package models

import "encoding/json"

// marshalObjectMetaFields encodes metadata for API requests. Non-nil empty maps for
// labels and annotations are included so callers can clear tags via PATCH; nil maps
// omit those keys (unchanged behavior for unset fields).
func marshalObjectMetaFields(name, uid string, annotations, labels map[string]string) ([]byte, error) {
	out := make(map[string]interface{})
	if name != "" {
		out["name"] = name
	}
	if uid != "" {
		out["uid"] = uid
	}
	if annotations != nil {
		out["annotations"] = annotations
	}
	if labels != nil {
		out["labels"] = labels
	}
	return json.Marshal(out)
}

// MarshalJSON implements encoding/json.Marshaler for V1ObjectMetaInputEntity.
func (m *V1ObjectMetaInputEntity) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return marshalObjectMetaFields(m.Name, "", m.Annotations, m.Labels)
}

// MarshalJSON implements encoding/json.Marshaler for V1ObjectMetaUpdateEntity.
func (m *V1ObjectMetaUpdateEntity) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return marshalObjectMetaFields(m.Name, m.UID, m.Annotations, m.Labels)
}

// MarshalJSON implements encoding/json.Marshaler for V1AppProfileMetaUpdateEntity.
func (m *V1AppProfileMetaUpdateEntity) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return marshalObjectMetaFields("", "", m.Annotations, m.Labels)
}
