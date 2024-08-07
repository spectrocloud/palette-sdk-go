// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMDataVolumeSource DataVolumeSource represents the source for our Data Volume, this can be HTTP, Imageio, S3, Registry or an existing PVC
//
// swagger:model v1VmDataVolumeSource
type V1VMDataVolumeSource struct {

	// blank
	Blank V1VMDataVolumeBlankImage `json:"blank,omitempty"`

	// http
	HTTP *V1VMDataVolumeSourceHTTP `json:"http,omitempty"`

	// imageio
	Imageio *V1VMDataVolumeSourceImageIO `json:"imageio,omitempty"`

	// pvc
	Pvc *V1VMDataVolumeSourcePVC `json:"pvc,omitempty"`

	// registry
	Registry *V1VMDataVolumeSourceRegistry `json:"registry,omitempty"`

	// s3
	S3 *V1VMDataVolumeSourceS3 `json:"s3,omitempty"`

	// upload
	Upload V1VMDataVolumeSourceUpload `json:"upload,omitempty"`

	// vddk
	Vddk *V1VMDataVolumeSourceVDDK `json:"vddk,omitempty"`
}

// Validate validates this v1 Vm data volume source
func (m *V1VMDataVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePvc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVddk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMDataVolumeSource) validateHTTP(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMDataVolumeSource) validateImageio(formats strfmt.Registry) error {

	if swag.IsZero(m.Imageio) { // not required
		return nil
	}

	if m.Imageio != nil {
		if err := m.Imageio.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageio")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMDataVolumeSource) validatePvc(formats strfmt.Registry) error {

	if swag.IsZero(m.Pvc) { // not required
		return nil
	}

	if m.Pvc != nil {
		if err := m.Pvc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pvc")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMDataVolumeSource) validateRegistry(formats strfmt.Registry) error {

	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMDataVolumeSource) validateS3(formats strfmt.Registry) error {

	if swag.IsZero(m.S3) { // not required
		return nil
	}

	if m.S3 != nil {
		if err := m.S3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMDataVolumeSource) validateVddk(formats strfmt.Registry) error {

	if swag.IsZero(m.Vddk) { // not required
		return nil
	}

	if m.Vddk != nil {
		if err := m.Vddk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vddk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMDataVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMDataVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1VMDataVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
