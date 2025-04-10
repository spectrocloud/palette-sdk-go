// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1PackLayer v1 pack layer
//
// swagger:model v1PackLayer
type V1PackLayer string

func NewV1PackLayer(value V1PackLayer) *V1PackLayer {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1PackLayer.
func (m V1PackLayer) Pointer() *V1PackLayer {
	return &m
}

const (

	// V1PackLayerKernel captures enum value "kernel"
	V1PackLayerKernel V1PackLayer = "kernel"

	// V1PackLayerOs captures enum value "os"
	V1PackLayerOs V1PackLayer = "os"

	// V1PackLayerK8s captures enum value "k8s"
	V1PackLayerK8s V1PackLayer = "k8s"

	// V1PackLayerCni captures enum value "cni"
	V1PackLayerCni V1PackLayer = "cni"

	// V1PackLayerCsi captures enum value "csi"
	V1PackLayerCsi V1PackLayer = "csi"

	// V1PackLayerAddon captures enum value "addon"
	V1PackLayerAddon V1PackLayer = "addon"
)

// for schema
var v1PackLayerEnum []interface{}

func init() {
	var res []V1PackLayer
	if err := json.Unmarshal([]byte(`["kernel","os","k8s","cni","csi","addon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PackLayerEnum = append(v1PackLayerEnum, v)
	}
}

func (m V1PackLayer) validateV1PackLayerEnum(path, location string, value V1PackLayer) error {
	if err := validate.EnumCase(path, location, value, v1PackLayerEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 pack layer
func (m V1PackLayer) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1PackLayerEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 pack layer based on context it is used
func (m V1PackLayer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
