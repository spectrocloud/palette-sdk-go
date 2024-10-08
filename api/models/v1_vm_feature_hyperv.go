// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VMFeatureHyperv Hyperv specific features.
//
// swagger:model v1VmFeatureHyperv
type V1VMFeatureHyperv struct {

	// evmcs
	Evmcs *V1VMFeatureState `json:"evmcs,omitempty"`

	// frequencies
	Frequencies *V1VMFeatureState `json:"frequencies,omitempty"`

	// ipi
	Ipi *V1VMFeatureState `json:"ipi,omitempty"`

	// reenlightenment
	Reenlightenment *V1VMFeatureState `json:"reenlightenment,omitempty"`

	// relaxed
	Relaxed *V1VMFeatureState `json:"relaxed,omitempty"`

	// reset
	Reset *V1VMFeatureState `json:"reset,omitempty"`

	// runtime
	Runtime *V1VMFeatureState `json:"runtime,omitempty"`

	// spinlocks
	Spinlocks *V1VMFeatureSpinlocks `json:"spinlocks,omitempty"`

	// synic
	Synic *V1VMFeatureState `json:"synic,omitempty"`

	// synictimer
	Synictimer *V1VMSyNICTimer `json:"synictimer,omitempty"`

	// tlbflush
	Tlbflush *V1VMFeatureState `json:"tlbflush,omitempty"`

	// vapic
	Vapic *V1VMFeatureState `json:"vapic,omitempty"`

	// vendorid
	Vendorid *V1VMFeatureVendorID `json:"vendorid,omitempty"`

	// vpindex
	Vpindex *V1VMFeatureState `json:"vpindex,omitempty"`
}

// Validate validates this v1 Vm feature hyperv
func (m *V1VMFeatureHyperv) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvmcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReenlightenment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelaxed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpinlocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSynic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSynictimer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTlbflush(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVapic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpindex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMFeatureHyperv) validateEvmcs(formats strfmt.Registry) error {

	if swag.IsZero(m.Evmcs) { // not required
		return nil
	}

	if m.Evmcs != nil {
		if err := m.Evmcs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("evmcs")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateFrequencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequencies) { // not required
		return nil
	}

	if m.Frequencies != nil {
		if err := m.Frequencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frequencies")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateIpi(formats strfmt.Registry) error {

	if swag.IsZero(m.Ipi) { // not required
		return nil
	}

	if m.Ipi != nil {
		if err := m.Ipi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipi")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateReenlightenment(formats strfmt.Registry) error {

	if swag.IsZero(m.Reenlightenment) { // not required
		return nil
	}

	if m.Reenlightenment != nil {
		if err := m.Reenlightenment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reenlightenment")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateRelaxed(formats strfmt.Registry) error {

	if swag.IsZero(m.Relaxed) { // not required
		return nil
	}

	if m.Relaxed != nil {
		if err := m.Relaxed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relaxed")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateReset(formats strfmt.Registry) error {

	if swag.IsZero(m.Reset) { // not required
		return nil
	}

	if m.Reset != nil {
		if err := m.Reset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reset")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateRuntime(formats strfmt.Registry) error {

	if swag.IsZero(m.Runtime) { // not required
		return nil
	}

	if m.Runtime != nil {
		if err := m.Runtime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateSpinlocks(formats strfmt.Registry) error {

	if swag.IsZero(m.Spinlocks) { // not required
		return nil
	}

	if m.Spinlocks != nil {
		if err := m.Spinlocks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spinlocks")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateSynic(formats strfmt.Registry) error {

	if swag.IsZero(m.Synic) { // not required
		return nil
	}

	if m.Synic != nil {
		if err := m.Synic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synic")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateSynictimer(formats strfmt.Registry) error {

	if swag.IsZero(m.Synictimer) { // not required
		return nil
	}

	if m.Synictimer != nil {
		if err := m.Synictimer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("synictimer")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateTlbflush(formats strfmt.Registry) error {

	if swag.IsZero(m.Tlbflush) { // not required
		return nil
	}

	if m.Tlbflush != nil {
		if err := m.Tlbflush.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tlbflush")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateVapic(formats strfmt.Registry) error {

	if swag.IsZero(m.Vapic) { // not required
		return nil
	}

	if m.Vapic != nil {
		if err := m.Vapic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vapic")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateVendorid(formats strfmt.Registry) error {

	if swag.IsZero(m.Vendorid) { // not required
		return nil
	}

	if m.Vendorid != nil {
		if err := m.Vendorid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vendorid")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMFeatureHyperv) validateVpindex(formats strfmt.Registry) error {

	if swag.IsZero(m.Vpindex) { // not required
		return nil
	}

	if m.Vpindex != nil {
		if err := m.Vpindex.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpindex")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMFeatureHyperv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMFeatureHyperv) UnmarshalBinary(b []byte) error {
	var res V1VMFeatureHyperv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
