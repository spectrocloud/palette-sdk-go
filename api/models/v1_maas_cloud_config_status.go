// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1MaasCloudConfigStatus MaasCloudConfigStatus defines the observed state of MaasCloudConfig The cloudimage info built by Mold is stored here image should be mapped to a specific machinepool
//
// swagger:model v1MaasCloudConfigStatus
type V1MaasCloudConfigStatus struct {

	// For mold controller to identify if is there any changes in Pack
	AnsibleRoleDigest string `json:"ansibleRoleDigest,omitempty"`

	// conditions
	Conditions []*V1ClusterCondition `json:"conditions"`

	// addon layers present in spc
	IsAddonLayer bool `json:"isAddonLayer,omitempty"`

	// node image
	NodeImage *V1MaasImage `json:"nodeImage,omitempty"`

	// this map will be for ansible roles present in eack pack
	RoleDigest map[string]string `json:"roleDigest,omitempty"`

	// sourceImageId, it can be from packref's annotations or from pack.json
	SourceImageID string `json:"sourceImageId,omitempty"`

	// PackerVariableDigest string `json:"packerDigest,omitempty"` If no ansible roles found in Packs then Mold should tell Drive to use capi image and not create custom image, because there is nothing to add
	UseCapiImage bool `json:"useCapiImage,omitempty"`
}

// Validate validates this v1 maas cloud config status
func (m *V1MaasCloudConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MaasCloudConfigStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MaasCloudConfigStatus) validateNodeImage(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeImage) { // not required
		return nil
	}

	if m.NodeImage != nil {
		if err := m.NodeImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeImage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 maas cloud config status based on the context it is used
func (m *V1MaasCloudConfigStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MaasCloudConfigStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {

			if swag.IsZero(m.Conditions[i]) { // not required
				return nil
			}

			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MaasCloudConfigStatus) contextValidateNodeImage(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeImage != nil {

		if swag.IsZero(m.NodeImage) { // not required
			return nil
		}

		if err := m.NodeImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeImage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MaasCloudConfigStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MaasCloudConfigStatus) UnmarshalBinary(b []byte) error {
	var res V1MaasCloudConfigStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
