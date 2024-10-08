// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OverloadSpec Overload spec
//
// swagger:model v1OverloadSpec
type V1OverloadSpec struct {

	// cloud account Uid
	CloudAccountUID string `json:"cloudAccountUid"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// ip pools
	IPPools []*V1IPPoolEntity `json:"ipPools"`

	// is self hosted
	IsSelfHosted bool `json:"isSelfHosted,omitempty"`

	// is system
	IsSystem bool `json:"isSystem,omitempty"`

	// spectro cluster Uid
	SpectroClusterUID string `json:"spectroClusterUid"`

	// tenant Uid
	TenantUID string `json:"tenantUid,omitempty"`
}

// Validate validates this v1 overload spec
func (m *V1OverloadSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPPools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OverloadSpec) validateIPPools(formats strfmt.Registry) error {

	if swag.IsZero(m.IPPools) { // not required
		return nil
	}

	for i := 0; i < len(m.IPPools); i++ {
		if swag.IsZero(m.IPPools[i]) { // not required
			continue
		}

		if m.IPPools[i] != nil {
			if err := m.IPPools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipPools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OverloadSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OverloadSpec) UnmarshalBinary(b []byte) error {
	var res V1OverloadSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
