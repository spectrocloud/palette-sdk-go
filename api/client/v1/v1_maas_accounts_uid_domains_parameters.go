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

// NewV1MaasAccountsUIDDomainsParams creates a new V1MaasAccountsUIDDomainsParams object
// with the default values initialized.
func NewV1MaasAccountsUIDDomainsParams() *V1MaasAccountsUIDDomainsParams {
	var ()
	return &V1MaasAccountsUIDDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MaasAccountsUIDDomainsParamsWithTimeout creates a new V1MaasAccountsUIDDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MaasAccountsUIDDomainsParamsWithTimeout(timeout time.Duration) *V1MaasAccountsUIDDomainsParams {
	var ()
	return &V1MaasAccountsUIDDomainsParams{

		timeout: timeout,
	}
}

// NewV1MaasAccountsUIDDomainsParamsWithContext creates a new V1MaasAccountsUIDDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MaasAccountsUIDDomainsParamsWithContext(ctx context.Context) *V1MaasAccountsUIDDomainsParams {
	var ()
	return &V1MaasAccountsUIDDomainsParams{

		Context: ctx,
	}
}

// NewV1MaasAccountsUIDDomainsParamsWithHTTPClient creates a new V1MaasAccountsUIDDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MaasAccountsUIDDomainsParamsWithHTTPClient(client *http.Client) *V1MaasAccountsUIDDomainsParams {
	var ()
	return &V1MaasAccountsUIDDomainsParams{
		HTTPClient: client,
	}
}

/*
V1MaasAccountsUIDDomainsParams contains all the parameters to send to the API endpoint
for the v1 maas accounts Uid domains operation typically these are written to a http.Request
*/
type V1MaasAccountsUIDDomainsParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) WithTimeout(timeout time.Duration) *V1MaasAccountsUIDDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) WithContext(ctx context.Context) *V1MaasAccountsUIDDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) WithHTTPClient(client *http.Client) *V1MaasAccountsUIDDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) WithUID(uid string) *V1MaasAccountsUIDDomainsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 maas accounts Uid domains params
func (o *V1MaasAccountsUIDDomainsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1MaasAccountsUIDDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
