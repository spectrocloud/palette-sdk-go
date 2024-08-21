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

// NewV1ClusterProfilesUIDVariablesDeleteParams creates a new V1ClusterProfilesUIDVariablesDeleteParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDVariablesDeleteParams() *V1ClusterProfilesUIDVariablesDeleteParams {
	var ()
	return &V1ClusterProfilesUIDVariablesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDVariablesDeleteParamsWithTimeout creates a new V1ClusterProfilesUIDVariablesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDVariablesDeleteParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDVariablesDeleteParams {
	var ()
	return &V1ClusterProfilesUIDVariablesDeleteParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDVariablesDeleteParamsWithContext creates a new V1ClusterProfilesUIDVariablesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDVariablesDeleteParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDVariablesDeleteParams {
	var ()
	return &V1ClusterProfilesUIDVariablesDeleteParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDVariablesDeleteParamsWithHTTPClient creates a new V1ClusterProfilesUIDVariablesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDVariablesDeleteParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDVariablesDeleteParams {
	var ()
	return &V1ClusterProfilesUIDVariablesDeleteParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDVariablesDeleteParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid variables delete operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDVariablesDeleteParams struct {

	/*Body*/
	Body *models.V1VariableNames
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDVariablesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDVariablesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDVariablesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) WithBody(body *models.V1VariableNames) *V1ClusterProfilesUIDVariablesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) SetBody(body *models.V1VariableNames) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) WithUID(uid string) *V1ClusterProfilesUIDVariablesDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid variables delete params
func (o *V1ClusterProfilesUIDVariablesDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDVariablesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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