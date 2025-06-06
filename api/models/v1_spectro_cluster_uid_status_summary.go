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

// V1SpectroClusterUIDStatusSummary Spectro cluster status summary
//
// swagger:model v1SpectroClusterUidStatusSummary
type V1SpectroClusterUIDStatusSummary struct {

	// abort timestamp
	// Format: date-time
	AbortTimestamp V1Time `json:"abortTimestamp,omitempty"`

	// add on services
	AddOnServices []*V1SpectroClusterAddOnServiceSummary `json:"addOnServices"`

	// api endpoints
	APIEndpoints []*V1APIEndpoint `json:"apiEndpoints"`

	// cluster import
	ClusterImport *V1ClusterImport `json:"clusterImport,omitempty"`

	// conditions
	Conditions []*V1ClusterCondition `json:"conditions"`

	// cost
	Cost *V1ResourceCost `json:"cost,omitempty"`

	// fips
	Fips *V1ClusterFips `json:"fips,omitempty"`

	// health
	Health *V1SpectroClusterHealthStatus `json:"health,omitempty"`

	// hourly rate
	HourlyRate *V1ResourceCost `json:"hourlyRate,omitempty"`

	// kube meta
	KubeMeta *V1KubeMeta `json:"kubeMeta,omitempty"`

	// location
	Location *V1ClusterMetaSpecLocation `json:"location,omitempty"`

	// metrics
	Metrics *V1SpectroClusterMetrics `json:"metrics,omitempty"`

	// notifications
	Notifications *V1ClusterNotificationStatus `json:"notifications,omitempty"`

	// packs
	Packs []*V1ClusterPackStatus `json:"packs"`

	// services
	Services []*V1LoadBalancerService `json:"services"`

	// spc apply
	SpcApply *V1SpcApply `json:"spcApply,omitempty"`

	// current operational state
	State string `json:"state,omitempty"`

	// upgrades
	Upgrades []*V1Upgrades `json:"upgrades"`

	// virtual
	Virtual *V1Virtual `json:"virtual,omitempty"`

	// workspaces
	Workspaces []*V1ResourceReference `json:"workspaces"`
}

// Validate validates this v1 spectro cluster Uid status summary
func (m *V1SpectroClusterUIDStatusSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbortTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddOnServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterImport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
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

	if err := m.validateKubeMeta(formats); err != nil {
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

	if err := m.validatePacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpcApply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgrades(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateAbortTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.AbortTimestamp) { // not required
		return nil
	}

	if err := m.AbortTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("abortTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("abortTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateAddOnServices(formats strfmt.Registry) error {
	if swag.IsZero(m.AddOnServices) { // not required
		return nil
	}

	for i := 0; i < len(m.AddOnServices); i++ {
		if swag.IsZero(m.AddOnServices[i]) { // not required
			continue
		}

		if m.AddOnServices[i] != nil {
			if err := m.AddOnServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addOnServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addOnServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateAPIEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.APIEndpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.APIEndpoints); i++ {
		if swag.IsZero(m.APIEndpoints[i]) { // not required
			continue
		}

		if m.APIEndpoints[i] != nil {
			if err := m.APIEndpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiEndpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiEndpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateClusterImport(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterImport) { // not required
		return nil
	}

	if m.ClusterImport != nil {
		if err := m.ClusterImport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterImport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterImport")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateConditions(formats strfmt.Registry) error {
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

func (m *V1SpectroClusterUIDStatusSummary) validateCost(formats strfmt.Registry) error {
	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateFips(formats strfmt.Registry) error {
	if swag.IsZero(m.Fips) { // not required
		return nil
	}

	if m.Fips != nil {
		if err := m.Fips.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fips")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateHourlyRate(formats strfmt.Registry) error {
	if swag.IsZero(m.HourlyRate) { // not required
		return nil
	}

	if m.HourlyRate != nil {
		if err := m.HourlyRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourlyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hourlyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateKubeMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.KubeMeta) { // not required
		return nil
	}

	if m.KubeMeta != nil {
		if err := m.KubeMeta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubeMeta")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validatePacks(formats strfmt.Registry) error {
	if swag.IsZero(m.Packs) { // not required
		return nil
	}

	for i := 0; i < len(m.Packs); i++ {
		if swag.IsZero(m.Packs[i]) { // not required
			continue
		}

		if m.Packs[i] != nil {
			if err := m.Packs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateSpcApply(formats strfmt.Registry) error {
	if swag.IsZero(m.SpcApply) { // not required
		return nil
	}

	if m.SpcApply != nil {
		if err := m.SpcApply.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spcApply")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spcApply")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateUpgrades(formats strfmt.Registry) error {
	if swag.IsZero(m.Upgrades) { // not required
		return nil
	}

	for i := 0; i < len(m.Upgrades); i++ {
		if swag.IsZero(m.Upgrades[i]) { // not required
			continue
		}

		if m.Upgrades[i] != nil {
			if err := m.Upgrades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("upgrades" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("upgrades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateVirtual(formats strfmt.Registry) error {
	if swag.IsZero(m.Virtual) { // not required
		return nil
	}

	if m.Virtual != nil {
		if err := m.Virtual.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) validateWorkspaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Workspaces); i++ {
		if swag.IsZero(m.Workspaces[i]) { // not required
			continue
		}

		if m.Workspaces[i] != nil {
			if err := m.Workspaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 spectro cluster Uid status summary based on the context it is used
func (m *V1SpectroClusterUIDStatusSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAbortTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddOnServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterImport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFips(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHourlyRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubeMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpcApply(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpgrades(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtual(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateAbortTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AbortTimestamp) { // not required
		return nil
	}

	if err := m.AbortTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("abortTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("abortTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateAddOnServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AddOnServices); i++ {

		if m.AddOnServices[i] != nil {

			if swag.IsZero(m.AddOnServices[i]) { // not required
				return nil
			}

			if err := m.AddOnServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addOnServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addOnServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateAPIEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APIEndpoints); i++ {

		if m.APIEndpoints[i] != nil {

			if swag.IsZero(m.APIEndpoints[i]) { // not required
				return nil
			}

			if err := m.APIEndpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiEndpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiEndpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateClusterImport(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterImport != nil {

		if swag.IsZero(m.ClusterImport) { // not required
			return nil
		}

		if err := m.ClusterImport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterImport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterImport")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1SpectroClusterUIDStatusSummary) contextValidateCost(ctx context.Context, formats strfmt.Registry) error {

	if m.Cost != nil {

		if swag.IsZero(m.Cost) { // not required
			return nil
		}

		if err := m.Cost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateFips(ctx context.Context, formats strfmt.Registry) error {

	if m.Fips != nil {

		if swag.IsZero(m.Fips) { // not required
			return nil
		}

		if err := m.Fips.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fips")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fips")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {

		if swag.IsZero(m.Health) { // not required
			return nil
		}

		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateHourlyRate(ctx context.Context, formats strfmt.Registry) error {

	if m.HourlyRate != nil {

		if swag.IsZero(m.HourlyRate) { // not required
			return nil
		}

		if err := m.HourlyRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hourlyRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hourlyRate")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateKubeMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.KubeMeta != nil {

		if swag.IsZero(m.KubeMeta) { // not required
			return nil
		}

		if err := m.KubeMeta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeMeta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubeMeta")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	if m.Metrics != nil {

		if swag.IsZero(m.Metrics) { // not required
			return nil
		}

		if err := m.Metrics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {

		if swag.IsZero(m.Notifications) { // not required
			return nil
		}

		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidatePacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Packs); i++ {

		if m.Packs[i] != nil {

			if swag.IsZero(m.Packs[i]) { // not required
				return nil
			}

			if err := m.Packs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {

			if swag.IsZero(m.Services[i]) { // not required
				return nil
			}

			if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateSpcApply(ctx context.Context, formats strfmt.Registry) error {

	if m.SpcApply != nil {

		if swag.IsZero(m.SpcApply) { // not required
			return nil
		}

		if err := m.SpcApply.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spcApply")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spcApply")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateUpgrades(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Upgrades); i++ {

		if m.Upgrades[i] != nil {

			if swag.IsZero(m.Upgrades[i]) { // not required
				return nil
			}

			if err := m.Upgrades[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("upgrades" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("upgrades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateVirtual(ctx context.Context, formats strfmt.Registry) error {

	if m.Virtual != nil {

		if swag.IsZero(m.Virtual) { // not required
			return nil
		}

		if err := m.Virtual.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual")
			}
			return err
		}
	}

	return nil
}

func (m *V1SpectroClusterUIDStatusSummary) contextValidateWorkspaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Workspaces); i++ {

		if m.Workspaces[i] != nil {

			if swag.IsZero(m.Workspaces[i]) { // not required
				return nil
			}

			if err := m.Workspaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("workspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SpectroClusterUIDStatusSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SpectroClusterUIDStatusSummary) UnmarshalBinary(b []byte) error {
	var res V1SpectroClusterUIDStatusSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
