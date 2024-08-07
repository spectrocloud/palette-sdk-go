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

// V1SyftReportEntity Syft report
//
// swagger:model v1SyftReportEntity
type V1SyftReportEntity struct {

	// batch no
	BatchNo int32 `json:"batchNo,omitempty"`

	// batch size
	BatchSize int32 `json:"batchSize,omitempty"`

	// dependencies
	Dependencies []*V1SyftDependencyEntity `json:"dependencies"`

	// image
	Image string `json:"image,omitempty"`

	// image contexts
	ImageContexts []*V1SyftImageContext `json:"imageContexts"`

	// sbom
	Sbom string `json:"sbom,omitempty"`

	// time
	// Format: date-time
	Time V1Time `json:"time,omitempty"`

	// vulnerabilities
	Vulnerabilities []*V1SyftVulnerabilityEntity `json:"vulnerabilities"`

	// vulnerability summary
	VulnerabilitySummary *V1SyftVulnerabilitySummaryEntity `json:"vulnerabilitySummary,omitempty"`
}

// Validate validates this v1 syft report entity
func (m *V1SyftReportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageContexts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilitySummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SyftReportEntity) validateDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SyftReportEntity) validateImageContexts(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageContexts) { // not required
		return nil
	}

	for i := 0; i < len(m.ImageContexts); i++ {
		if swag.IsZero(m.ImageContexts[i]) { // not required
			continue
		}

		if m.ImageContexts[i] != nil {
			if err := m.ImageContexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imageContexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SyftReportEntity) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		}
		return err
	}

	return nil
}

func (m *V1SyftReportEntity) validateVulnerabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Vulnerabilities); i++ {
		if swag.IsZero(m.Vulnerabilities[i]) { // not required
			continue
		}

		if m.Vulnerabilities[i] != nil {
			if err := m.Vulnerabilities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vulnerabilities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SyftReportEntity) validateVulnerabilitySummary(formats strfmt.Registry) error {

	if swag.IsZero(m.VulnerabilitySummary) { // not required
		return nil
	}

	if m.VulnerabilitySummary != nil {
		if err := m.VulnerabilitySummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerabilitySummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SyftReportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SyftReportEntity) UnmarshalBinary(b []byte) error {
	var res V1SyftReportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
