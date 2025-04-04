// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Channel v1 channel
//
// swagger:model v1Channel
type V1Channel struct {

	// alert all users
	AlertAllUsers bool `json:"alertAllUsers"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// http
	HTTP *V1ChannelHTTP `json:"http,omitempty"`

	// identifiers
	// Unique: true
	Identifiers []string `json:"identifiers"`

	// is active
	IsActive bool `json:"isActive"`

	// status
	Status *V1AlertNotificationStatus `json:"status,omitempty"`

	// type
	// Enum: ["email","app","http"]
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 channel
func (m *V1Channel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *V1Channel) validateHTTP(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *V1Channel) validateIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	if err := validate.UniqueItems("identifiers", "body", m.Identifiers); err != nil {
		return err
	}

	return nil
}

func (m *V1Channel) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

var v1ChannelTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["email","app","http"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ChannelTypeTypePropEnum = append(v1ChannelTypeTypePropEnum, v)
	}
}

const (

	// V1ChannelTypeEmail captures enum value "email"
	V1ChannelTypeEmail string = "email"

	// V1ChannelTypeApp captures enum value "app"
	V1ChannelTypeApp string = "app"

	// V1ChannelTypeHTTP captures enum value "http"
	V1ChannelTypeHTTP string = "http"
)

// prop value enum
func (m *V1Channel) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ChannelTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1Channel) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 channel based on the context it is used
func (m *V1Channel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Channel) contextValidateHTTP(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTP != nil {

		if swag.IsZero(m.HTTP) { // not required
			return nil
		}

		if err := m.HTTP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *V1Channel) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Channel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Channel) UnmarshalBinary(b []byte) error {
	var res V1Channel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1ChannelHTTP v1 channel HTTP
//
// swagger:model V1ChannelHTTP
type V1ChannelHTTP struct {

	// body
	Body string `json:"body,omitempty"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 channel HTTP
func (m *V1ChannelHTTP) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 channel HTTP based on context it is used
func (m *V1ChannelHTTP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ChannelHTTP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ChannelHTTP) UnmarshalBinary(b []byte) error {
	var res V1ChannelHTTP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
