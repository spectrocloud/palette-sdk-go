// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewV1BasicOciRegistriesUIDSyncParams creates a new V1BasicOciRegistriesUIDSyncParams object
// with the default values initialized.
func NewV1BasicOciRegistriesUIDSyncParams() *V1BasicOciRegistriesUIDSyncParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &V1BasicOciRegistriesUIDSyncParams{
		ForceSync: &forceSyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1BasicOciRegistriesUIDSyncParamsWithTimeout creates a new V1BasicOciRegistriesUIDSyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1BasicOciRegistriesUIDSyncParamsWithTimeout(timeout time.Duration) *V1BasicOciRegistriesUIDSyncParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &V1BasicOciRegistriesUIDSyncParams{
		ForceSync: &forceSyncDefault,

		timeout: timeout,
	}
}

// NewV1BasicOciRegistriesUIDSyncParamsWithContext creates a new V1BasicOciRegistriesUIDSyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1BasicOciRegistriesUIDSyncParamsWithContext(ctx context.Context) *V1BasicOciRegistriesUIDSyncParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &V1BasicOciRegistriesUIDSyncParams{
		ForceSync: &forceSyncDefault,

		Context: ctx,
	}
}

// NewV1BasicOciRegistriesUIDSyncParamsWithHTTPClient creates a new V1BasicOciRegistriesUIDSyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1BasicOciRegistriesUIDSyncParamsWithHTTPClient(client *http.Client) *V1BasicOciRegistriesUIDSyncParams {
	var (
		forceSyncDefault = bool(false)
	)
	return &V1BasicOciRegistriesUIDSyncParams{
		ForceSync:  &forceSyncDefault,
		HTTPClient: client,
	}
}

/*
V1BasicOciRegistriesUIDSyncParams contains all the parameters to send to the API endpoint
for the v1 basic oci registries Uid sync operation typically these are written to a http.Request
*/
type V1BasicOciRegistriesUIDSyncParams struct {

	/*ForceSync*/
	ForceSync *bool
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) WithTimeout(timeout time.Duration) *V1BasicOciRegistriesUIDSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) WithContext(ctx context.Context) *V1BasicOciRegistriesUIDSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) WithHTTPClient(client *http.Client) *V1BasicOciRegistriesUIDSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceSync adds the forceSync to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) WithForceSync(forceSync *bool) *V1BasicOciRegistriesUIDSyncParams {
	o.SetForceSync(forceSync)
	return o
}

// SetForceSync adds the forceSync to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) SetForceSync(forceSync *bool) {
	o.ForceSync = forceSync
}

// WithUID adds the uid to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) WithUID(uid string) *V1BasicOciRegistriesUIDSyncParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 basic oci registries Uid sync params
func (o *V1BasicOciRegistriesUIDSyncParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1BasicOciRegistriesUIDSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceSync != nil {

		// query param forceSync
		var qrForceSync bool
		if o.ForceSync != nil {
			qrForceSync = *o.ForceSync
		}
		qForceSync := swag.FormatBool(qrForceSync)
		if qForceSync != "" {
			if err := r.SetQueryParam("forceSync", qForceSync); err != nil {
				return err
			}
		}

	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}