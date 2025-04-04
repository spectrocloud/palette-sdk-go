// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectCleanup Project delete request payload
//
// swagger:model v1ProjectCleanup
type V1ProjectCleanup struct {

	// deleting cluster duration threshold in min
	DeletingClusterDurationThresholdInMin int32 `json:"deletingClusterDurationThresholdInMin,omitempty"`

	// provisioning cluster duration threshold in min
	ProvisioningClusterDurationThresholdInMin int32 `json:"provisioningClusterDurationThresholdInMin,omitempty"`
}

// Validate validates this v1 project cleanup
func (m *V1ProjectCleanup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 project cleanup based on context it is used
func (m *V1ProjectCleanup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectCleanup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectCleanup) UnmarshalBinary(b []byte) error {
	var res V1ProjectCleanup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
