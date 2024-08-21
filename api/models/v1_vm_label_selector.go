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

// V1VMLabelSelector A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
//
// swagger:model v1VmLabelSelector
type V1VMLabelSelector struct {

	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []*V1VMLabelSelectorRequirement `json:"matchExpressions"`

	// matchLabels is a map of key-value pairs. A single key-value in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// Validate validates this v1 Vm label selector
func (m *V1VMLabelSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchExpressions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMLabelSelector) validateMatchExpressions(formats strfmt.Registry) error {

	if swag.IsZero(m.MatchExpressions) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchExpressions); i++ {
		if swag.IsZero(m.MatchExpressions[i]) { // not required
			continue
		}

		if m.MatchExpressions[i] != nil {
			if err := m.MatchExpressions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMLabelSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMLabelSelector) UnmarshalBinary(b []byte) error {
	var res V1VMLabelSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}