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

// V1AuditActor Audit actor object
//
// swagger:model v1AuditActor
type V1AuditActor struct {

	// actor type
	// Enum: [user system service]
	ActorType string `json:"actorType,omitempty"`

	// project
	Project *V1ProjectMeta `json:"project,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// user
	User *V1UserMeta `json:"user,omitempty"`
}

// Validate validates this v1 audit actor
func (m *V1AuditActor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1AuditActorTypeActorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","system","service"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1AuditActorTypeActorTypePropEnum = append(v1AuditActorTypeActorTypePropEnum, v)
	}
}

const (

	// V1AuditActorActorTypeUser captures enum value "user"
	V1AuditActorActorTypeUser string = "user"

	// V1AuditActorActorTypeSystem captures enum value "system"
	V1AuditActorActorTypeSystem string = "system"

	// V1AuditActorActorTypeService captures enum value "service"
	V1AuditActorActorTypeService string = "service"
)

// prop value enum
func (m *V1AuditActor) validateActorTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1AuditActorTypeActorTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1AuditActor) validateActorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ActorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateActorTypeEnum("actorType", "body", m.ActorType); err != nil {
		return err
	}

	return nil
}

func (m *V1AuditActor) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *V1AuditActor) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AuditActor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AuditActor) UnmarshalBinary(b []byte) error {
	var res V1AuditActor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
