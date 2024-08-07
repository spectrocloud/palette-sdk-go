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

// V1ClusterProfileTemplateMeta Cluster profile template meta information
//
// swagger:model v1ClusterProfileTemplateMeta
type V1ClusterProfileTemplateMeta struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// Cluster profile name
	Name string `json:"name,omitempty"`

	// Cluster profile packs array
	Packs []*V1PackRef `json:"packs"`

	// scope or context(system, tenant or project)
	Scope string `json:"scope,omitempty"`

	// Cluster profile type [ "cluster", "infra", "add-on", "system" ]
	Type string `json:"type,omitempty"`

	// Cluster profile uid
	UID string `json:"uid,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this v1 cluster profile template meta
func (m *V1ClusterProfileTemplateMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterProfileTemplateMeta) validatePacks(formats strfmt.Registry) error {

	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterProfileTemplateMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterProfileTemplateMeta) UnmarshalBinary(b []byte) error {
	var res V1ClusterProfileTemplateMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
