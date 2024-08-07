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

// NewV1SpectroClustersClusterRbacParams creates a new V1SpectroClustersClusterRbacParams object
// with the default values initialized.
func NewV1SpectroClustersClusterRbacParams() *V1SpectroClustersClusterRbacParams {
	var ()
	return &V1SpectroClustersClusterRbacParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersClusterRbacParamsWithTimeout creates a new V1SpectroClustersClusterRbacParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersClusterRbacParamsWithTimeout(timeout time.Duration) *V1SpectroClustersClusterRbacParams {
	var ()
	return &V1SpectroClustersClusterRbacParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersClusterRbacParamsWithContext creates a new V1SpectroClustersClusterRbacParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersClusterRbacParamsWithContext(ctx context.Context) *V1SpectroClustersClusterRbacParams {
	var ()
	return &V1SpectroClustersClusterRbacParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersClusterRbacParamsWithHTTPClient creates a new V1SpectroClustersClusterRbacParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersClusterRbacParamsWithHTTPClient(client *http.Client) *V1SpectroClustersClusterRbacParams {
	var ()
	return &V1SpectroClustersClusterRbacParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersClusterRbacParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters cluster rbac operation typically these are written to a http.Request
*/
type V1SpectroClustersClusterRbacParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) WithTimeout(timeout time.Duration) *V1SpectroClustersClusterRbacParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) WithContext(ctx context.Context) *V1SpectroClustersClusterRbacParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) WithHTTPClient(client *http.Client) *V1SpectroClustersClusterRbacParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) WithUID(uid string) *V1SpectroClustersClusterRbacParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters cluster rbac params
func (o *V1SpectroClustersClusterRbacParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersClusterRbacParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
