package models

import (
	"encoding/json"
	"testing"
)

func TestV1ObjectMetaInputEntity_MarshalJSON_emptyLabels(t *testing.T) {
	m := &V1ObjectMetaInputEntity{
		Name:        "profile-a",
		Annotations: map[string]string{"description": "test"},
		Labels:      map[string]string{},
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	labels, ok := got["labels"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected labels object in JSON, got %s", string(b))
	}
	if len(labels) != 0 {
		t.Fatalf("expected empty labels object, got %v", labels)
	}
	if _, ok := got["name"]; !ok {
		t.Fatalf("expected name in JSON, got %s", string(b))
	}
}

func TestV1ObjectMetaInputEntity_MarshalJSON_nilLabelsOmitted(t *testing.T) {
	m := &V1ObjectMetaInputEntity{
		Name:   "profile-a",
		Labels: nil,
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if _, ok := got["labels"]; ok {
		t.Fatalf("nil labels should be omitted, got %s", string(b))
	}
}

func TestV1ProfileMetaEntity_MarshalJSON_emptyLabels(t *testing.T) {
	body := &V1ProfileMetaEntity{
		Metadata: &V1ObjectMetaInputEntity{
			Name:   "profile-a",
			Labels: map[string]string{},
		},
	}
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !jsonContainsKeyPath(t, b, "metadata", "labels") {
		t.Fatalf("expected metadata.labels in patch body, got %s", string(b))
	}
}

func jsonContainsKeyPath(t *testing.T, b []byte, keys ...string) bool {
	t.Helper()
	var cur interface{}
	if err := json.Unmarshal(b, &cur); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, key := range keys {
		m, ok := cur.(map[string]interface{})
		if !ok {
			return false
		}
		cur, ok = m[key]
		if !ok {
			return false
		}
	}
	return true
}
