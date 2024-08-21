// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMFieldsV1 FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.
//
// Each key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:\<index>', where \<index> is position of a item in a list 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.
//
// The exact format is defined in sigs.k8s.io/structured-merge-diff
//
// swagger:model v1VmFieldsV1
type V1VMFieldsV1 struct {

	// raw
	Raw []strfmt.Base64 `json:"Raw"`
}

// Validate validates this v1 Vm fields v1
func (m *V1VMFieldsV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMFieldsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMFieldsV1) UnmarshalBinary(b []byte) error {
	var res V1VMFieldsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}