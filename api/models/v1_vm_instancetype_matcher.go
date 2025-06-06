// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMInstancetypeMatcher InstancetypeMatcher references a instancetype that is used to fill fields in the VMI template.
//
// swagger:model v1VmInstancetypeMatcher
type V1VMInstancetypeMatcher struct {

	// InferFromVolume lists the name of a volume that should be used to infer or discover the instancetype to be used through known annotations on the underlying resource. Once applied to the InstancetypeMatcher this field is removed.
	InferFromVolume string `json:"inferFromVolume,omitempty"`

	// Kind specifies which instancetype resource is referenced. Allowed values are: "VirtualMachineInstancetype" and "VirtualMachineClusterInstancetype". If not specified, "VirtualMachineClusterInstancetype" is used by default.
	Kind string `json:"kind,omitempty"`

	// Name is the name of the VirtualMachineInstancetype or VirtualMachineClusterInstancetype
	Name string `json:"name,omitempty"`

	// RevisionName specifies a ControllerRevision containing a specific copy of the VirtualMachineInstancetype or VirtualMachineClusterInstancetype to be used. This is initially captured the first time the instancetype is applied to the VirtualMachineInstance.
	RevisionName string `json:"revisionName,omitempty"`
}

// Validate validates this v1 Vm instancetype matcher
func (m *V1VMInstancetypeMatcher) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 Vm instancetype matcher based on context it is used
func (m *V1VMInstancetypeMatcher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VMInstancetypeMatcher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMInstancetypeMatcher) UnmarshalBinary(b []byte) error {
	var res V1VMInstancetypeMatcher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
