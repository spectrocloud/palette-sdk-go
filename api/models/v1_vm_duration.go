// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// V1VMDuration Duration is a wrapper around time.Duration which supports correct marshaling to YAML and JSON. In particular, it marshals into strings, which can be used as map keys in json.
//
// swagger:model v1VmDuration
type V1VMDuration string

// Validate validates this v1 Vm duration
func (m V1VMDuration) Validate(formats strfmt.Registry) error {
	return nil
}