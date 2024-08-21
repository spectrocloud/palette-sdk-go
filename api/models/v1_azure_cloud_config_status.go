// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AzureCloudConfigStatus AzureCloudConfigStatus defines the observed state of AzureCloudConfig The cloudimage info built by Mold is stored here image should be mapped to a specific machinepool
//
// swagger:model v1AzureCloudConfigStatus
type V1AzureCloudConfigStatus struct {

	// For mold controller to identify if is there any changes in Pack
	AnsibleRoleDigest string `json:"ansibleRoleDigest,omitempty"`

	// spectroAnsibleProvisioner: should be added only once, subsequent recocile will use the same provisioner SpectroAnsiblePacker bool `json:"spectroAnsiblePacker,omitempty"`
	Conditions []*V1ClusterCondition `json:"conditions"`

	// Images array items should be 1-to-1 mapping to Spec.MachinePoolConfig
	Images *V1AzureImage `json:"images,omitempty"`

	// addon layers present in spc
	IsAddonLayer bool `json:"isAddonLayer,omitempty"`

	// this map will be for ansible roles present in eack pack
	RoleDigest map[string]string `json:"roleDigest,omitempty"`

	// sourceImageId, it can be from packref's annotations or from pack.json
	SourceImageID string `json:"sourceImageId,omitempty"`

	// PackerVariableDigest string `json:"packerDigest,omitempty"` If no ansible roles found in Packs then Mold should tell Drive to use capi image and not create custom image, because there is nothing to add
	UseCapiImage bool `json:"useCapiImage,omitempty"`

	// vhd image
	VhdImage *V1AzureVHDImage `json:"vhdImage,omitempty"`
}

// Validate validates this v1 azure cloud config status
func (m *V1AzureCloudConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVhdImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1AzureCloudConfigStatus) validateConditions(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1AzureCloudConfigStatus) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

func (m *V1AzureCloudConfigStatus) validateVhdImage(formats strfmt.Registry) error {

	if swag.IsZero(m.VhdImage) { // not required
		return nil
	}

	if m.VhdImage != nil {
		if err := m.VhdImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vhdImage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureCloudConfigStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureCloudConfigStatus) UnmarshalBinary(b []byte) error {
	var res V1AzureCloudConfigStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}