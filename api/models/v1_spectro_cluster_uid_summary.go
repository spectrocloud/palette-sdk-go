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

// V1SpectroClusterUIDSummary Spectro cluster summary
//
// swagger:model v1SpectroClusterUidSummary
type V1SpectroClusterUIDSummary struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *V1SpectroClusterUIDSummarySpec `json:"spec,omitempty"`

	// status
	Status *V1SpectroClusterUIDStatusSummary `json:"status,omitempty"`
}

// Validate validates this v1 spectro cluster Uid summary
func (m *V1SpectroClusterUIDSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterUIDSummary) validateMetadata(formats strfmt.Registry) error {

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

func (m *V1SpectroClusterUIDSummary) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDSummary) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterUIDSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterUIDSummary) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterUIDSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroClusterUIDSummarySpec Spectro cluster spec summary
//
// swagger:model V1SpectroClusterUIDSummarySpec
type V1SpectroClusterUIDSummarySpec struct {

	// Architecture types of the cluster
	ArchTypes []V1ArchType `json:"archTypes"`

	// cloud config
	CloudConfig *V1CloudConfigMeta `json:"cloudConfig,omitempty"`

	// cloudaccount
	Cloudaccount *V1CloudAccountMeta `json:"cloudaccount,omitempty"`

	// cluster profile template
	ClusterProfileTemplate *V1ClusterProfileTemplateMeta `json:"clusterProfileTemplate,omitempty"`

	// cluster profile templates
	ClusterProfileTemplates []*V1ClusterProfileTemplateMeta `json:"clusterProfileTemplates"`
}

// Validate validates this v1 spectro cluster UID summary spec
func (m *V1SpectroClusterUIDSummarySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudaccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfileTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfileTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterUIDSummarySpec) validateArchTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ArchTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ArchTypes); i++ {

		if err := m.ArchTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "archTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDSummarySpec) validateCloudConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConfig) { // not required
		return nil
	}

	if m.CloudConfig != nil {
		if err := m.CloudConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDSummarySpec) validateCloudaccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Cloudaccount) { // not required
		return nil
	}

	if m.Cloudaccount != nil {
		if err := m.Cloudaccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "cloudaccount")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDSummarySpec) validateClusterProfileTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfileTemplate) { // not required
		return nil
	}

	if m.ClusterProfileTemplate != nil {
		if err := m.ClusterProfileTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec" + "." + "clusterProfileTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDSummarySpec) validateClusterProfileTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfileTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterProfileTemplates); i++ {
		if swag.IsZero(m.ClusterProfileTemplates[i]) { // not required
			continue
		}

		if m.ClusterProfileTemplates[i] != nil {
			if err := m.ClusterProfileTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spec" + "." + "clusterProfileTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterUIDSummarySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterUIDSummarySpec) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterUIDSummarySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
