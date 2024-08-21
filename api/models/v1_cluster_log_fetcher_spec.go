// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ClusterLogFetcherSpec Cluster Log Fetcher Spec
//
// swagger:model v1ClusterLogFetcherSpec
type V1ClusterLogFetcherSpec struct {

	// cluster Uid
	ClusterUID string `json:"clusterUid,omitempty"`

	// log
	Log string `json:"log,omitempty"`
}

// Validate validates this v1 cluster log fetcher spec
func (m *V1ClusterLogFetcherSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterLogFetcherSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterLogFetcherSpec) UnmarshalBinary(b []byte) error {
	var res V1ClusterLogFetcherSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}