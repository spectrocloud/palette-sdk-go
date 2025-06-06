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

// NewV1SpectroClustersUIDConfigRbacsUIDGetParams creates a new V1SpectroClustersUIDConfigRbacsUIDGetParams object
// with the default values initialized.
func NewV1SpectroClustersUIDConfigRbacsUIDGetParams() *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDGetParamsWithTimeout creates a new V1SpectroClustersUIDConfigRbacsUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDConfigRbacsUIDGetParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDGetParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDGetParamsWithContext creates a new V1SpectroClustersUIDConfigRbacsUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDConfigRbacsUIDGetParamsWithContext(ctx context.Context) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDGetParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDGetParamsWithHTTPClient creates a new V1SpectroClustersUIDConfigRbacsUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDConfigRbacsUIDGetParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDConfigRbacsUIDGetParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid config rbacs Uid get operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDConfigRbacsUIDGetParams struct {

	/*RbacUID
	  RBAC resource uid

	*/
	RbacUID string
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) WithContext(ctx context.Context) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRbacUID adds the rbacUID to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) WithRbacUID(rbacUID string) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	o.SetRbacUID(rbacUID)
	return o
}

// SetRbacUID adds the rbacUid to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) SetRbacUID(rbacUID string) {
	o.RbacUID = rbacUID
}

// WithUID adds the uid to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) WithUID(uid string) *V1SpectroClustersUIDConfigRbacsUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid config rbacs Uid get params
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDConfigRbacsUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rbacUid
	if err := r.SetPathParam("rbacUid", o.RbacUID); err != nil {
		return err
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
