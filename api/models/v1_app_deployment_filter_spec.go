// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AppDeploymentFilterSpec Application deployment filter spec
//
// swagger:model v1AppDeploymentFilterSpec
type V1AppDeploymentFilterSpec struct {

	// app deployment name
	AppDeploymentName *V1FilterString `json:"appDeploymentName,omitempty"`

	// cluster uids
	ClusterUids *V1FilterArray `json:"clusterUids,omitempty"`

	// tags
	Tags *V1FilterArray `json:"tags,omitempty"`
}

// Validate validates this v1 app deployment filter spec
func (m *V1AppDeploymentFilterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppDeploymentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterUids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentFilterSpec) validateAppDeploymentName(formats strfmt.Registry) error {
	if swag.IsZero(m.AppDeploymentName) { // not required
		return nil
	}

	if m.AppDeploymentName != nil {
		if err := m.AppDeploymentName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appDeploymentName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appDeploymentName")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppDeploymentFilterSpec) validateClusterUids(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterUids) { // not required
		return nil
	}

	if m.ClusterUids != nil {
		if err := m.ClusterUids.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterUids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterUids")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppDeploymentFilterSpec) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 app deployment filter spec based on the context it is used
func (m *V1AppDeploymentFilterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppDeploymentName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterUids(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AppDeploymentFilterSpec) contextValidateAppDeploymentName(ctx context.Context, formats strfmt.Registry) error {

	if m.AppDeploymentName != nil {

		if swag.IsZero(m.AppDeploymentName) { // not required
			return nil
		}

		if err := m.AppDeploymentName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appDeploymentName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appDeploymentName")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppDeploymentFilterSpec) contextValidateClusterUids(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterUids != nil {

		if swag.IsZero(m.ClusterUids) { // not required
			return nil
		}

		if err := m.ClusterUids.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterUids")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterUids")
			}
			return err
		}
	}

	return nil
}

func (m *V1AppDeploymentFilterSpec) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	if m.Tags != nil {

		if swag.IsZero(m.Tags) { // not required
			return nil
		}

		if err := m.Tags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AppDeploymentFilterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AppDeploymentFilterSpec) UnmarshalBinary(b []byte) error {
	var res V1AppDeploymentFilterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
