// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CustomCloudTypeContentResponse Custom cloudType content response
//
// swagger:model v1CustomCloudTypeContentResponse
type V1CustomCloudTypeContentResponse struct {

	// custom cloud type content
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this v1 custom cloud type content response
func (m *V1CustomCloudTypeContentResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CustomCloudTypeContentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CustomCloudTypeContentResponse) UnmarshalBinary(b []byte) error {
	var res V1CustomCloudTypeContentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}