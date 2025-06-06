// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SpectroClusterMetaSummary Spectro cluster meta summary
//
// swagger:model v1SpectroClusterMetaSummary
type V1SpectroClusterMetaSummary struct {

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// spec summary
	SpecSummary *V1SpectroClusterMetaSummarySpecSummary `json:"specSummary,omitempty"`

	// status
	Status *V1SpectroClusterMetaSummaryStatus `json:"status,omitempty"`
}

// Validate validates this v1 spectro cluster meta summary
func (m *V1SpectroClusterMetaSummary) Validate(formats strfmt.Registry) error {
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

func (m *V1SpectroClusterMetaSummary) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummary) validateSpecSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.SpecSummary) { // not required
		return nil
	}

	if m.SpecSummary != nil {
		if err := m.SpecSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("specSummary")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummary) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this v1 spectro cluster meta summary based on the context it is used
func (m *V1SpectroClusterMetaSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpecSummary(ctx, formats); err != nil {
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

func (m *V1SpectroClusterMetaSummary) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummary) contextValidateSpecSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.SpecSummary != nil {

		if swag.IsZero(m.SpecSummary) { // not required
			return nil
		}

		if err := m.SpecSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("specSummary")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummary) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1SpectroClusterMetaSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterMetaSummary) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterMetaSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroClusterMetaSummarySpecSummary Spectro cluster meta summary
//
// swagger:model V1SpectroClusterMetaSummarySpecSummary
type V1SpectroClusterMetaSummarySpecSummary struct {

	// Architecture type of the cluster
	ArchType []*string `json:"archType"`

	// cloud account Uid
	CloudAccountUID string `json:"cloudAccountUid,omitempty"`

	// cloud region
	CloudRegion string `json:"cloudRegion,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// cluster type
	ClusterType string `json:"clusterType,omitempty"`

	// import mode
	ImportMode string `json:"importMode,omitempty"`

	// location
	Location *V1ClusterMetaSpecLocation `json:"location,omitempty"`

	// project meta
	ProjectMeta *V1ProjectMeta `json:"projectMeta,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this v1 spectro cluster meta summary spec summary
func (m *V1SpectroClusterMetaSummarySpecSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
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

var v1SpectroClusterMetaSummarySpecSummaryArchTypeItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["arm64","amd64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SpectroClusterMetaSummarySpecSummaryArchTypeItemsEnum = append(v1SpectroClusterMetaSummarySpecSummaryArchTypeItemsEnum, v)
	}
}

func (m *V1SpectroClusterMetaSummarySpecSummary) validateArchTypeItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1SpectroClusterMetaSummarySpecSummaryArchTypeItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1SpectroClusterMetaSummarySpecSummary) validateArchType(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchType) { // not required
		return nil
	}

	for i := 0; i < len(m.ArchType); i++ {
		if swag.IsZero(m.ArchType[i]) { // not required
			continue
		}

		// value enum
		if err := m.validateArchTypeItemsEnum("specSummary"+"."+"archType"+"."+strconv.Itoa(i), "body", *m.ArchType[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1SpectroClusterMetaSummarySpecSummary) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("specSummary" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummarySpecSummary) validateProjectMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectMeta) { // not required
		return nil
	}

	if m.ProjectMeta != nil {
		if err := m.ProjectMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "projectMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("specSummary" + "." + "projectMeta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 spectro cluster meta summary spec summary based on the context it is used
func (m *V1SpectroClusterMetaSummarySpecSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterMetaSummarySpecSummary) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("specSummary" + "." + "location")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummarySpecSummary) contextValidateProjectMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.ProjectMeta != nil {

		if swag.IsZero(m.ProjectMeta) { // not required
			return nil
		}

		if err := m.ProjectMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("specSummary" + "." + "projectMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("specSummary" + "." + "projectMeta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterMetaSummarySpecSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterMetaSummarySpecSummary) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterMetaSummarySpecSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1SpectroClusterMetaSummaryStatus Spectro cluster meta status summary
//
// swagger:model V1SpectroClusterMetaSummaryStatus
type V1SpectroClusterMetaSummaryStatus struct {

	// cost
	Cost *V1ClusterMetaStatusCost `json:"cost,omitempty"`

	// fips
	Fips *V1ClusterFips `json:"fips,omitempty"`

	// health
	Health *V1ClusterMetaStatusHealth `json:"health,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// updates
	Updates *V1ClusterMetaStatusUpdates `json:"updates,omitempty"`
}

// Validate validates this v1 spectro cluster meta summary status
func (m *V1SpectroClusterMetaSummaryStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFips(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) validateCost(formats strfmt.Registry) error {
	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "cost")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) validateFips(formats strfmt.Registry) error {
	if swag.IsZero(m.Fips) { // not required
		return nil
	}

	if m.Fips != nil {
		if err := m.Fips.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "fips")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "health")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) validateUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.Updates) { // not required
		return nil
	}

	if m.Updates != nil {
		if err := m.Updates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "updates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 spectro cluster meta summary status based on the context it is used
func (m *V1SpectroClusterMetaSummaryStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) contextValidateCost(ctx context.Context, formats strfmt.Registry) error {

	if m.Cost != nil {

		if swag.IsZero(m.Cost) { // not required
			return nil
		}

		if err := m.Cost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "cost")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) contextValidateFips(ctx context.Context, formats strfmt.Registry) error {

	if m.Fips != nil {

		if swag.IsZero(m.Fips) { // not required
			return nil
		}

		if err := m.Fips.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "fips")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {

		if swag.IsZero(m.Health) { // not required
			return nil
		}

		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "health")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterMetaSummaryStatus) contextValidateUpdates(ctx context.Context, formats strfmt.Registry) error {

	if m.Updates != nil {

		if swag.IsZero(m.Updates) { // not required
			return nil
		}

		if err := m.Updates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + "updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + "updates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterMetaSummaryStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterMetaSummaryStatus) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterMetaSummaryStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
