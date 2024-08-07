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

// NewV1ClusterProfilesUIDVariablesPatchParams creates a new V1ClusterProfilesUIDVariablesPatchParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDVariablesPatchParams() *V1ClusterProfilesUIDVariablesPatchParams {
	var ()
	return &V1ClusterProfilesUIDVariablesPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDVariablesPatchParamsWithTimeout creates a new V1ClusterProfilesUIDVariablesPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDVariablesPatchParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDVariablesPatchParams {
	var ()
	return &V1ClusterProfilesUIDVariablesPatchParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDVariablesPatchParamsWithContext creates a new V1ClusterProfilesUIDVariablesPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDVariablesPatchParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDVariablesPatchParams {
	var ()
	return &V1ClusterProfilesUIDVariablesPatchParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDVariablesPatchParamsWithHTTPClient creates a new V1ClusterProfilesUIDVariablesPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDVariablesPatchParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDVariablesPatchParams {
	var ()
	return &V1ClusterProfilesUIDVariablesPatchParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDVariablesPatchParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid variables patch operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDVariablesPatchParams struct {

	/*Body*/
	Body *models.V1Variables
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDVariablesPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDVariablesPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDVariablesPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) WithBody(body *models.V1Variables) *V1ClusterProfilesUIDVariablesPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) SetBody(body *models.V1Variables) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) WithUID(uid string) *V1ClusterProfilesUIDVariablesPatchParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid variables patch params
func (o *V1ClusterProfilesUIDVariablesPatchParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDVariablesPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
