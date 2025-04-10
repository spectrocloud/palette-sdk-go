// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ClusterLogFetcherRequest Cluster Log Fetcher Request
//
// swagger:model v1ClusterLogFetcherRequest
type V1ClusterLogFetcherRequest struct {

	// Duration for which log is requested
	Duration *int64 `json:"duration,omitempty"`

	// k8s
	K8s *V1ClusterLogFetcherK8sRequest `json:"k8s,omitempty"`

	// Accepted Values - ["cluster", "app"]. if "app" then logs will be fetched from the virtual cluster
	// Enum: ["cluster","app"]
	Mode *string `json:"mode,omitempty"`

	// No of lines of logs requested
	NoOfLines *int64 `json:"noOfLines,omitempty"`

	// node
	Node *V1ClusterLogFetcherNodeRequest `json:"node,omitempty"`
}

// Validate validates this v1 cluster log fetcher request
func (m *V1ClusterLogFetcherRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateK8s(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterLogFetcherRequest) validateK8s(formats strfmt.Registry) error {
	if swag.IsZero(m.K8s) { // not required
		return nil
	}

	if m.K8s != nil {
		if err := m.K8s.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("k8s")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("k8s")
			}
			return err
		}
	}

	return nil
}

var v1ClusterLogFetcherRequestTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster","app"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ClusterLogFetcherRequestTypeModePropEnum = append(v1ClusterLogFetcherRequestTypeModePropEnum, v)
	}
}

const (

	// V1ClusterLogFetcherRequestModeCluster captures enum value "cluster"
	V1ClusterLogFetcherRequestModeCluster string = "cluster"

	// V1ClusterLogFetcherRequestModeApp captures enum value "app"
	V1ClusterLogFetcherRequestModeApp string = "app"
)

// prop value enum
func (m *V1ClusterLogFetcherRequest) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1ClusterLogFetcherRequestTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1ClusterLogFetcherRequest) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterLogFetcherRequest) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 cluster log fetcher request based on the context it is used
func (m *V1ClusterLogFetcherRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateK8s(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterLogFetcherRequest) contextValidateK8s(ctx context.Context, formats strfmt.Registry) error {

	if m.K8s != nil {

		if swag.IsZero(m.K8s) { // not required
			return nil
		}

		if err := m.K8s.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("k8s")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("k8s")
			}
			return err
		}
	}

	return nil
}

func (m *V1ClusterLogFetcherRequest) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterLogFetcherRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterLogFetcherRequest) UnmarshalBinary(b []byte) error {
	var res V1ClusterLogFetcherRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
