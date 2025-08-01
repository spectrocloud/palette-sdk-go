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

// V1TotalResourceUsage Total Resource Usage
//
// swagger:model v1TotalResourceUsage
type V1TotalResourceUsage struct {

	// projects
	// Unique: true
	Projects []*V1ProjectResourceUsage `json:"projects"`

	// total alloy Cpu core hours
	TotalAlloyCPUCoreHours float64 `json:"totalAlloyCpuCoreHours"`

	// total pure Cpu core hours
	TotalPureCPUCoreHours float64 `json:"totalPureCpuCoreHours"`
}

// Validate validates this v1 total resource usage
func (m *V1TotalResourceUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TotalResourceUsage) validateProjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Projects) { // not required
		return nil
	}

	if err := validate.UniqueItems("projects", "body", m.Projects); err != nil {
		return err
	}

	for i := 0; i < len(m.Projects); i++ {
		if swag.IsZero(m.Projects[i]) { // not required
			continue
		}

		if m.Projects[i] != nil {
			if err := m.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TotalResourceUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TotalResourceUsage) UnmarshalBinary(b []byte) error {
	var res V1TotalResourceUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
