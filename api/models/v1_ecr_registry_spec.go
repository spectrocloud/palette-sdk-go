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

// V1EcrRegistrySpec Ecr registry spec
//
// swagger:model v1EcrRegistrySpec
type V1EcrRegistrySpec struct {

	// OCI ecr registry content base path
	BaseContentPath string `json:"baseContentPath,omitempty"`

	// contains spectro manifest
	ContainsSpectroManifest bool `json:"containsSpectroManifest"`

	// credentials
	Credentials *V1AwsCloudAccount `json:"credentials,omitempty"`

	// default region
	DefaultRegion string `json:"defaultRegion,omitempty"`

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// is private
	// Required: true
	IsPrivate *bool `json:"isPrivate"`

	// is sync supported
	IsSyncSupported bool `json:"isSyncSupported,omitempty"`

	// provider type
	// Enum: ["helm","pack"]
	ProviderType *string `json:"providerType,omitempty"`

	// Ecr registry uid
	RegistryUID string `json:"registryUid,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// tls
	TLS *V1TLSConfiguration `json:"tls,omitempty"`
}

// Validate validates this v1 ecr registry spec
func (m *V1EcrRegistrySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPrivate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EcrRegistrySpec) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *V1EcrRegistrySpec) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *V1EcrRegistrySpec) validateIsPrivate(formats strfmt.Registry) error {

	if err := validate.Required("isPrivate", "body", m.IsPrivate); err != nil {
		return err
	}

	return nil
}

var v1EcrRegistrySpecTypeProviderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["helm","pack"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1EcrRegistrySpecTypeProviderTypePropEnum = append(v1EcrRegistrySpecTypeProviderTypePropEnum, v)
	}
}

const (

	// V1EcrRegistrySpecProviderTypeHelm captures enum value "helm"
	V1EcrRegistrySpecProviderTypeHelm string = "helm"

	// V1EcrRegistrySpecProviderTypePack captures enum value "pack"
	V1EcrRegistrySpecProviderTypePack string = "pack"
)

// prop value enum
func (m *V1EcrRegistrySpec) validateProviderTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1EcrRegistrySpecTypeProviderTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1EcrRegistrySpec) validateProviderType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderTypeEnum("providerType", "body", *m.ProviderType); err != nil {
		return err
	}

	return nil
}

func (m *V1EcrRegistrySpec) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 ecr registry spec based on the context it is used
func (m *V1EcrRegistrySpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EcrRegistrySpec) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {

		if swag.IsZero(m.Credentials) { // not required
			return nil
		}

		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *V1EcrRegistrySpec) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.TLS != nil {

		if swag.IsZero(m.TLS) { // not required
			return nil
		}

		if err := m.TLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EcrRegistrySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EcrRegistrySpec) UnmarshalBinary(b []byte) error {
	var res V1EcrRegistrySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
