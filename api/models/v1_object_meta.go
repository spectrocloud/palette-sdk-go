// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ObjectMeta ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
//
// swagger:model v1ObjectMeta
type V1ObjectMeta struct {

	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// Format: date-time
	CreationTimestamp V1Time `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.
	//
	// Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// Format: date-time
	DeletionTimestamp V1Time `json:"deletionTimestamp,omitempty"`

	// Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	Labels map[string]string `json:"labels,omitempty"`

	// LastModifiedTimestamp is a timestamp representing the server time when this object was last modified. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// Format: date-time
	LastModifiedTimestamp V1Time `json:"lastModifiedTimestamp,omitempty"`

	// Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name,omitempty"`

	// UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.
	//
	// Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID string `json:"uid,omitempty"`
}

// Validate validates this v1 object meta
func (m *V1ObjectMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ObjectMeta) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := m.CreationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("creationTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ObjectMeta) validateDeletionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletionTimestamp) { // not required
		return nil
	}

	if err := m.DeletionTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deletionTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deletionTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ObjectMeta) validateLastModifiedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModifiedTimestamp) { // not required
		return nil
	}

	if err := m.LastModifiedTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModifiedTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastModifiedTimestamp")
		}
		return err
	}

	return nil
}

// ContextValidate validate this v1 object meta based on the context it is used
func (m *V1ObjectMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreationTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletionTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastModifiedTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ObjectMeta) contextValidateCreationTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := m.CreationTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("creationTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("creationTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ObjectMeta) contextValidateDeletionTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DeletionTimestamp) { // not required
		return nil
	}

	if err := m.DeletionTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deletionTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deletionTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1ObjectMeta) contextValidateLastModifiedTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedTimestamp) { // not required
		return nil
	}

	if err := m.LastModifiedTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModifiedTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastModifiedTimestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ObjectMeta) UnmarshalBinary(b []byte) error {
	var res V1ObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
