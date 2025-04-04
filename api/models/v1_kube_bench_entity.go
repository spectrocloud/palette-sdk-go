// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1KubeBenchEntity KubeBench response
//
// swagger:model v1KubeBenchEntity
type V1KubeBenchEntity struct {

	// reports
	// Required: true
	Reports map[string]V1KubeBenchReportEntity `json:"reports"`

	// request Uid
	// Required: true
	RequestUID *string `json:"requestUid"`

	// status
	// Required: true
	// Enum: ["Completed","InProgress","Failed","Initiated"]
	Status *string `json:"status"`
}

// Validate validates this v1 kube bench entity
func (m *V1KubeBenchEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubeBenchEntity) validateReports(formats strfmt.Registry) error {

	if err := validate.Required("reports", "body", m.Reports); err != nil {
		return err
	}

	for k := range m.Reports {

		if err := validate.Required("reports"+"."+k, "body", m.Reports[k]); err != nil {
			return err
		}
		if val, ok := m.Reports[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reports" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reports" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1KubeBenchEntity) validateRequestUID(formats strfmt.Registry) error {

	if err := validate.Required("requestUid", "body", m.RequestUID); err != nil {
		return err
	}

	return nil
}

var v1KubeBenchEntityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Completed","InProgress","Failed","Initiated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1KubeBenchEntityTypeStatusPropEnum = append(v1KubeBenchEntityTypeStatusPropEnum, v)
	}
}

const (

	// V1KubeBenchEntityStatusCompleted captures enum value "Completed"
	V1KubeBenchEntityStatusCompleted string = "Completed"

	// V1KubeBenchEntityStatusInProgress captures enum value "InProgress"
	V1KubeBenchEntityStatusInProgress string = "InProgress"

	// V1KubeBenchEntityStatusFailed captures enum value "Failed"
	V1KubeBenchEntityStatusFailed string = "Failed"

	// V1KubeBenchEntityStatusInitiated captures enum value "Initiated"
	V1KubeBenchEntityStatusInitiated string = "Initiated"
)

// prop value enum
func (m *V1KubeBenchEntity) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1KubeBenchEntityTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1KubeBenchEntity) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 kube bench entity based on the context it is used
func (m *V1KubeBenchEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubeBenchEntity) contextValidateReports(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("reports", "body", m.Reports); err != nil {
		return err
	}

	for k := range m.Reports {

		if val, ok := m.Reports[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1KubeBenchEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KubeBenchEntity) UnmarshalBinary(b []byte) error {
	var res V1KubeBenchEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
