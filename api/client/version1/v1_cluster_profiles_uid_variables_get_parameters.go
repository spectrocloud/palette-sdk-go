// Code generated by go-swagger; DO NOT EDIT.

package version1

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
)

// NewV1ClusterProfilesUIDVariablesGetParams creates a new V1ClusterProfilesUIDVariablesGetParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDVariablesGetParams() *V1ClusterProfilesUIDVariablesGetParams {
	var ()
	return &V1ClusterProfilesUIDVariablesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDVariablesGetParamsWithTimeout creates a new V1ClusterProfilesUIDVariablesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDVariablesGetParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDVariablesGetParams {
	var ()
	return &V1ClusterProfilesUIDVariablesGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDVariablesGetParamsWithContext creates a new V1ClusterProfilesUIDVariablesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDVariablesGetParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDVariablesGetParams {
	var ()
	return &V1ClusterProfilesUIDVariablesGetParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDVariablesGetParamsWithHTTPClient creates a new V1ClusterProfilesUIDVariablesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDVariablesGetParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDVariablesGetParams {
	var ()
	return &V1ClusterProfilesUIDVariablesGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDVariablesGetParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid variables get operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDVariablesGetParams struct {

	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDVariablesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDVariablesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDVariablesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) WithUID(uid string) *V1ClusterProfilesUIDVariablesGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid variables get params
func (o *V1ClusterProfilesUIDVariablesGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDVariablesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
