// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ResourceCostSummaryOptions Resource cost summary options
//
// swagger:model v1ResourceCostSummaryOptions
type V1ResourceCostSummaryOptions struct {

	// enable summary view
	EnableSummaryView *bool `json:"enableSummaryView,omitempty"`

	// group by
	// Enum: [tenant project workspace cluster namespace deployment cloud]
	GroupBy *string `json:"groupBy,omitempty"`

	// period
	Period *int32 `json:"period,omitempty"`
}

// Validate validates this v1 resource cost summary options
func (m *V1ResourceCostSummaryOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1ResourceCostSummaryOptionsTypeGroupByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tenant","project","workspace","cluster","namespace","deployment","cloud"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ResourceCostSummaryOptionsTypeGroupByPropEnum = append(v1ResourceCostSummaryOptionsTypeGroupByPropEnum, v)
	}
}

const (

	// V1ResourceCostSummaryOptionsGroupByTenant captures enum value "tenant"
	V1ResourceCostSummaryOptionsGroupByTenant string = "tenant"

	// V1ResourceCostSummaryOptionsGroupByProject captures enum value "project"
	V1ResourceCostSummaryOptionsGroupByProject string = "project"

	// V1ResourceCostSummaryOptionsGroupByWorkspace captures enum value "workspace"
	V1ResourceCostSummaryOptionsGroupByWorkspace string = "workspace"

	// V1ResourceCostSummaryOptionsGroupByCluster captures enum value "cluster"
	V1ResourceCostSummaryOptionsGroupByCluster string = "cluster"

	// V1ResourceCostSummaryOptionsGroupByNamespace captures enum value "namespace"
	V1ResourceCostSummaryOptionsGroupByNamespace string = "namespace"

	// V1ResourceCostSummaryOptionsGroupByDeployment captures enum value "deployment"
	V1ResourceCostSummaryOptionsGroupByDeployment string = "deployment"

	// V1ResourceCostSummaryOptionsGroupByCloud captures enum value "cloud"
	V1ResourceCostSummaryOptionsGroupByCloud string = "cloud"
)

// prop value enum
func (m *V1ResourceCostSummaryOptions) validateGroupByEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ResourceCostSummaryOptionsTypeGroupByPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ResourceCostSummaryOptions) validateGroupBy(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupBy) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupByEnum("groupBy", "body", *m.GroupBy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceCostSummaryOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceCostSummaryOptions) UnmarshalBinary(b []byte) error {
	var res V1ResourceCostSummaryOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}