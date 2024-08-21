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

// V1Notification Describes event notification and action definition
//
// swagger:model v1Notification
type V1Notification struct {

	// Describes actions for the notification
	Action *V1NotificationAction `json:"action,omitempty"`

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// related object
	RelatedObject *V1RelatedObject `json:"relatedObject,omitempty"`

	// Describes origin info for the notification
	Source *V1NotificationSource `json:"source,omitempty"`

	// Describes type of notification. Possible values [NotificationPackUpdate, NotificationPackRegistryUpdate, NotificationNone]
	// Enum: [NotificationPackUpdate NotificationPackRegistryUpdate NotificationNone]
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 notification
func (m *V1Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Notification) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *V1Notification) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1Notification) validateRelatedObject(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedObject) { // not required
		return nil
	}

	if m.RelatedObject != nil {
		if err := m.RelatedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relatedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V1Notification) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

var v1NotificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotificationPackUpdate","NotificationPackRegistryUpdate","NotificationNone"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1NotificationTypeTypePropEnum = append(v1NotificationTypeTypePropEnum, v)
	}
}

const (

	// V1NotificationTypeNotificationPackUpdate captures enum value "NotificationPackUpdate"
	V1NotificationTypeNotificationPackUpdate string = "NotificationPackUpdate"

	// V1NotificationTypeNotificationPackRegistryUpdate captures enum value "NotificationPackRegistryUpdate"
	V1NotificationTypeNotificationPackRegistryUpdate string = "NotificationPackRegistryUpdate"

	// V1NotificationTypeNotificationNone captures enum value "NotificationNone"
	V1NotificationTypeNotificationNone string = "NotificationNone"
)

// prop value enum
func (m *V1Notification) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1NotificationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1Notification) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Notification) UnmarshalBinary(b []byte) error {
	var res V1Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}