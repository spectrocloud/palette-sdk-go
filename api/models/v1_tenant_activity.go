// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1TenantActivity Active tenant and clusters data
//
// swagger:model v1TenantActivity
type V1TenantActivity struct {

	// clusters info
	ClustersInfo *V1ClustersInfo `json:"clustersInfo,omitempty"`

	// org
	Org string `json:"org,omitempty"`

	// plan type
	PlanType string `json:"planType,omitempty"`

	// total projects
	TotalProjects int64 `json:"totalProjects,omitempty"`

	// total users
	TotalUsers int64 `json:"totalUsers,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// users
	// Unique: true
	Users []*V1UserActivityInfo `json:"users"`
}

// Validate validates this v1 tenant activity
func (m *V1TenantActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClustersInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TenantActivity) validateClustersInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ClustersInfo) { // not required
		return nil
	}

	if m.ClustersInfo != nil {
		if err := m.ClustersInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clustersInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V1TenantActivity) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	if err := validate.UniqueItems("users", "body", m.Users); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TenantActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TenantActivity) UnmarshalBinary(b []byte) error {
	var res V1TenantActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
