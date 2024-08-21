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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1AppProfilesUIDTiersPatchParams creates a new V1AppProfilesUIDTiersPatchParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersPatchParams() *V1AppProfilesUIDTiersPatchParams {
	var ()
	return &V1AppProfilesUIDTiersPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersPatchParamsWithTimeout creates a new V1AppProfilesUIDTiersPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersPatchParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersPatchParams {
	var ()
	return &V1AppProfilesUIDTiersPatchParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersPatchParamsWithContext creates a new V1AppProfilesUIDTiersPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersPatchParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersPatchParams {
	var ()
	return &V1AppProfilesUIDTiersPatchParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersPatchParamsWithHTTPClient creates a new V1AppProfilesUIDTiersPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersPatchParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersPatchParams {
	var ()
	return &V1AppProfilesUIDTiersPatchParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesUIDTiersPatchParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers patch operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersPatchParams struct {

	/*Body*/
	Body *models.V1AppTierPatchEntity
	/*UID
	  Application profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) WithBody(body *models.V1AppTierPatchEntity) *V1AppProfilesUIDTiersPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) SetBody(body *models.V1AppTierPatchEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) WithUID(uid string) *V1AppProfilesUIDTiersPatchParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers patch params
func (o *V1AppProfilesUIDTiersPatchParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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