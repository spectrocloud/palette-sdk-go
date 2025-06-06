// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GcpAccountValidateSpec Gcp cloud account entity which takes json credentials or reference to the file where credentials are stored
//
// swagger:model v1GcpAccountValidateSpec
type V1GcpAccountValidateSpec struct {

	// Gcp cloud account json credentials
	JSONCredentials string `json:"jsonCredentials,omitempty"`

	// Deprecated - Use jsonCredentials for Gcp account credentials; Reference of the credentials stored in the file
	JSONCredentialsFileUID string `json:"jsonCredentialsFileUid,omitempty"`
}

// Validate validates this v1 gcp account validate spec
func (m *V1GcpAccountValidateSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 gcp account validate spec based on context it is used
func (m *V1GcpAccountValidateSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GcpAccountValidateSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GcpAccountValidateSpec) UnmarshalBinary(b []byte) error {
	var res V1GcpAccountValidateSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
