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

// V1SpectroClusterSummary Spectro cluster summary
//
// swagger:model v1SpectroClusterSummary
type V1SpectroClusterSummary struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec summary
	SpecSummary *V1SpectroClusterSummarySpecSummary `json:"specSummary,omitempty"`

	// status
	Status *V1SpectroClusterSummaryStatus `json:"status,omitempty"`
}

// Validate validates this v1 spectro cluster summary
func (m *V1SpectroClusterSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecSummary(formats); err != nil {
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

func (m *V1SpectroClusterSummary) validateMetadata(formats strfmt.Registry) error {

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

func (m *V1SpectroClusterSummary) validateSpecSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.SpecSummary) { // not required
		return nil
	}

	if m.SpecSummary != nil {
		if err := m.SpecSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummary) validateStatus(formats strfmt.Registry) error {

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
func (m *V1SpectroClusterSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterSummary) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroClusterSummarySpecSummary Spectro cluster spec summary
//
// swagger:model V1SpectroClusterSummarySpecSummary
type V1SpectroClusterSummarySpecSummary struct {

	// Architecture type of the cluster
	ArchTypes []V1ArchType `json:"archTypes"`

	// cloud account meta
	CloudAccountMeta *V1CloudAccountMeta `json:"cloudAccountMeta,omitempty"`

	// cloud config
	CloudConfig *V1CloudConfigMeta `json:"cloudConfig,omitempty"`

	// cluster config
	ClusterConfig *V1ClusterConfigResponse `json:"clusterConfig,omitempty"`

	// cluster profile template
	ClusterProfileTemplate *V1ClusterProfileTemplateMeta `json:"clusterProfileTemplate,omitempty"`

	// cluster profile templates
	ClusterProfileTemplates []*V1ClusterProfileTemplateMeta `json:"clusterProfileTemplates"`

	// project meta
	ProjectMeta *V1ProjectMeta `json:"projectMeta,omitempty"`
}

// Validate validates this v1 spectro cluster summary spec summary
func (m *V1SpectroClusterSummarySpecSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfileTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterProfileTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateArchTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ArchTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ArchTypes); i++ {

		if err := m.ArchTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "archTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateCloudAccountMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudAccountMeta) { // not required
		return nil
	}

	if m.CloudAccountMeta != nil {
		if err := m.CloudAccountMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "cloudAccountMeta")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateCloudConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConfig) { // not required
		return nil
	}

	if m.CloudConfig != nil {
		if err := m.CloudConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "cloudConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateClusterConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterConfig) { // not required
		return nil
	}

	if m.ClusterConfig != nil {
		if err := m.ClusterConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "clusterConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateClusterProfileTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterProfileTemplate) { // not required
		return nil
	}

	if m.ClusterProfileTemplate != nil {
		if err := m.ClusterProfileTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "clusterProfileTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateClusterProfileTemplates(formats strfmt.Registry) error {

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
					return ve.ValidateName("specSummary" + "." + "clusterProfileTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterSummarySpecSummary) validateProjectMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectMeta) { // not required
		return nil
	}

	if m.ProjectMeta != nil {
		if err := m.ProjectMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "projectMeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterSummarySpecSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterSummarySpecSummary) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterSummarySpecSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroClusterSummaryStatus Spectro cluster status summary
//
// swagger:model V1SpectroClusterSummaryStatus
type V1SpectroClusterSummaryStatus struct {

	// cluster import
	ClusterImport *V1ClusterImport `json:"clusterImport,omitempty"`

	// cost
	Cost *V1ResourceCost `json:"cost,omitempty"`

	// fips
	Fips *V1ClusterFips `json:"fips,omitempty"`

	// health
	Health *V1SpectroClusterHealthStatus `json:"health,omitempty"`

	// hourly rate
	HourlyRate *V1ResourceCost `json:"hourlyRate,omitempty"`

	// location
	Location *V1ClusterMetaSpecLocation `json:"location,omitempty"`

	// metrics
	Metrics *V1SpectroClusterMetrics `json:"metrics,omitempty"`

	// notifications
	Notifications *V1ClusterNotificationStatus `json:"notifications,omitempty"`

	// repave
	Repave *V1ClusterRepaveStatus `json:"repave,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// virtual
	Virtual *V1Virtual `json:"virtual,omitempty"`
}

// Validate validates this v1 spectro cluster summary status
func (m *V1SpectroClusterSummaryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterImport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFips(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHourlyRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepave(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtual(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateClusterImport(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterImport) { // not required
		return nil
	}

	if m.ClusterImport != nil {
		if err := m.ClusterImport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "clusterImport")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateCost(formats strfmt.Registry) error {

	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "cost")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateFips(formats strfmt.Registry) error {

	if swag.IsZero(m.Fips) { // not required
		return nil
	}

	if m.Fips != nil {
		if err := m.Fips.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "fips")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "health")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateHourlyRate(formats strfmt.Registry) error {

	if swag.IsZero(m.HourlyRate) { // not required
		return nil
	}

	if m.HourlyRate != nil {
		if err := m.HourlyRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "hourlyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "metrics")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "notifications")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateRepave(formats strfmt.Registry) error {

	if swag.IsZero(m.Repave) { // not required
		return nil
	}

	if m.Repave != nil {
		if err := m.Repave.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "repave")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterSummaryStatus) validateVirtual(formats strfmt.Registry) error {

	if swag.IsZero(m.Virtual) { // not required
		return nil
	}

	if m.Virtual != nil {
		if err := m.Virtual.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "virtual")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterSummaryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterSummaryStatus) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterSummaryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
