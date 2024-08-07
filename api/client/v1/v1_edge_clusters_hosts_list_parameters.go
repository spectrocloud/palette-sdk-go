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
)

// NewV1EdgeClustersHostsListParams creates a new V1EdgeClustersHostsListParams object
// with the default values initialized.
func NewV1EdgeClustersHostsListParams() *V1EdgeClustersHostsListParams {
	var ()
	return &V1EdgeClustersHostsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeClustersHostsListParamsWithTimeout creates a new V1EdgeClustersHostsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeClustersHostsListParamsWithTimeout(timeout time.Duration) *V1EdgeClustersHostsListParams {
	var ()
	return &V1EdgeClustersHostsListParams{

		timeout: timeout,
	}
}

// NewV1EdgeClustersHostsListParamsWithContext creates a new V1EdgeClustersHostsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeClustersHostsListParamsWithContext(ctx context.Context) *V1EdgeClustersHostsListParams {
	var ()
	return &V1EdgeClustersHostsListParams{

		Context: ctx,
	}
}

// NewV1EdgeClustersHostsListParamsWithHTTPClient creates a new V1EdgeClustersHostsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeClustersHostsListParamsWithHTTPClient(client *http.Client) *V1EdgeClustersHostsListParams {
	var ()
	return &V1EdgeClustersHostsListParams{
		HTTPClient: client,
	}
}

/*
V1EdgeClustersHostsListParams contains all the parameters to send to the API endpoint
for the v1 edge clusters hosts list operation typically these are written to a http.Request
*/
type V1EdgeClustersHostsListParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) WithTimeout(timeout time.Duration) *V1EdgeClustersHostsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) WithContext(ctx context.Context) *V1EdgeClustersHostsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) WithHTTPClient(client *http.Client) *V1EdgeClustersHostsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) WithUID(uid string) *V1EdgeClustersHostsListParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge clusters hosts list params
func (o *V1EdgeClustersHostsListParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeClustersHostsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
