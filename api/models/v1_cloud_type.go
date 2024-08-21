// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1CloudType v1 cloud type
//
// swagger:model v1CloudType
type V1CloudType string

func NewV1CloudType(value V1CloudType) *V1CloudType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1CloudType.
func (m V1CloudType) Pointer() *V1CloudType {
	return &m
}

const (

	// V1CloudTypeAll captures enum value "all"
	V1CloudTypeAll V1CloudType = "all"

	// V1CloudTypeAws captures enum value "aws"
	V1CloudTypeAws V1CloudType = "aws"

	// V1CloudTypeAzure captures enum value "azure"
	V1CloudTypeAzure V1CloudType = "azure"

	// V1CloudTypeGcp captures enum value "gcp"
	V1CloudTypeGcp V1CloudType = "gcp"

	// V1CloudTypeVsphere captures enum value "vsphere"
	V1CloudTypeVsphere V1CloudType = "vsphere"

	// V1CloudTypeOpenstack captures enum value "openstack"
	V1CloudTypeOpenstack V1CloudType = "openstack"

	// V1CloudTypeMaas captures enum value "maas"
	V1CloudTypeMaas V1CloudType = "maas"

	// V1CloudTypeNested captures enum value "nested"
	V1CloudTypeNested V1CloudType = "nested"

	// V1CloudTypeBaremetal captures enum value "baremetal"
	V1CloudTypeBaremetal V1CloudType = "baremetal"

	// V1CloudTypeEks captures enum value "eks"
	V1CloudTypeEks V1CloudType = "eks"

	// V1CloudTypeAks captures enum value "aks"
	V1CloudTypeAks V1CloudType = "aks"

	// V1CloudTypeEdge captures enum value "edge"
	V1CloudTypeEdge V1CloudType = "edge"

	// V1CloudTypeEdgeDashNative captures enum value "edge-native"
	V1CloudTypeEdgeDashNative V1CloudType = "edge-native"

	// V1CloudTypeLibvirt captures enum value "libvirt"
	V1CloudTypeLibvirt V1CloudType = "libvirt"

	// V1CloudTypeTencent captures enum value "tencent"
	V1CloudTypeTencent V1CloudType = "tencent"

	// V1CloudTypeTke captures enum value "tke"
	V1CloudTypeTke V1CloudType = "tke"

	// V1CloudTypeCoxedge captures enum value "coxedge"
	V1CloudTypeCoxedge V1CloudType = "coxedge"

	// V1CloudTypeGeneric captures enum value "generic"
	V1CloudTypeGeneric V1CloudType = "generic"

	// V1CloudTypeGke captures enum value "gke"
	V1CloudTypeGke V1CloudType = "gke"
)

// for schema
var v1CloudTypeEnum []interface{}

func init() {
	var res []V1CloudType
	if err := json.Unmarshal([]byte(`["all","aws","azure","gcp","vsphere","openstack","maas","nested","baremetal","eks","aks","edge","edge-native","libvirt","tencent","tke","coxedge","generic","gke"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CloudTypeEnum = append(v1CloudTypeEnum, v)
	}
}

func (m V1CloudType) validateV1CloudTypeEnum(path, location string, value V1CloudType) error {
	if err := validate.EnumCase(path, location, value, v1CloudTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cloud type
func (m V1CloudType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1CloudTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 cloud type based on context it is used
func (m V1CloudType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}