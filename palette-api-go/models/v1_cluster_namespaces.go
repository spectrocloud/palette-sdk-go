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

// V1ClusterNamespaces Cluster's available namespaces
//
// swagger:model v1ClusterNamespaces
type V1ClusterNamespaces struct {

	// namespaces
	Namespaces []*V1ClusterNamespace `json:"namespaces"`
}

// Validate validates this v1 cluster namespaces
func (m *V1ClusterNamespaces) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterNamespaces) validateNamespaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Namespaces); i++ {
		if swag.IsZero(m.Namespaces[i]) { // not required
			continue
		}

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterNamespaces) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterNamespaces) UnmarshalBinary(b []byte) error {
	var res V1ClusterNamespaces
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
