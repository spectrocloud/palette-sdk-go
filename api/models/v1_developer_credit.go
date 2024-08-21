// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DeveloperCredit Credits allocated for each tenant/user
//
// swagger:model v1DeveloperCredit
type V1DeveloperCredit struct {

	// cpu in cores
	CPU int32 `json:"cpu"`

	// memory in GiB
	MemoryGiB int32 `json:"memoryGiB"`

	// storage in GiB
	StorageGiB int32 `json:"storageGiB"`

	// number of active virtual clusters
	VirtualClustersLimit int32 `json:"virtualClustersLimit"`
}

// Validate validates this v1 developer credit
func (m *V1DeveloperCredit) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DeveloperCredit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeveloperCredit) UnmarshalBinary(b []byte) error {
	var res V1DeveloperCredit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}