// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1HybridPoolClusterCloudType Flag to indicate whether the pool is deployed in any cloud or an edge environment
//
// swagger:model v1HybridPoolClusterCloudType
type V1HybridPoolClusterCloudType string

const (

	// V1HybridPoolClusterCloudTypeEdgeNative captures enum value "edge-native"
	V1HybridPoolClusterCloudTypeEdgeNative V1HybridPoolClusterCloudType = "edge-native"
)

// for schema
var v1HybridPoolClusterCloudTypeEnum []interface{}

func init() {
	var res []V1HybridPoolClusterCloudType
	if err := json.Unmarshal([]byte(`["edge-native"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1HybridPoolClusterCloudTypeEnum = append(v1HybridPoolClusterCloudTypeEnum, v)
	}
}

func (m V1HybridPoolClusterCloudType) validateV1HybridPoolClusterCloudTypeEnum(path, location string, value V1HybridPoolClusterCloudType) error {
	if err := validate.EnumCase(path, location, value, v1HybridPoolClusterCloudTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 hybrid pool cluster cloud type
func (m V1HybridPoolClusterCloudType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1HybridPoolClusterCloudTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}