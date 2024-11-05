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

// NewV1SpectroClustersCertificatesRenewParams creates a new V1SpectroClustersCertificatesRenewParams object
// with the default values initialized.
func NewV1SpectroClustersCertificatesRenewParams() *V1SpectroClustersCertificatesRenewParams {
	var ()
	return &V1SpectroClustersCertificatesRenewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersCertificatesRenewParamsWithTimeout creates a new V1SpectroClustersCertificatesRenewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersCertificatesRenewParamsWithTimeout(timeout time.Duration) *V1SpectroClustersCertificatesRenewParams {
	var ()
	return &V1SpectroClustersCertificatesRenewParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersCertificatesRenewParamsWithContext creates a new V1SpectroClustersCertificatesRenewParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersCertificatesRenewParamsWithContext(ctx context.Context) *V1SpectroClustersCertificatesRenewParams {
	var ()
	return &V1SpectroClustersCertificatesRenewParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersCertificatesRenewParamsWithHTTPClient creates a new V1SpectroClustersCertificatesRenewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersCertificatesRenewParamsWithHTTPClient(client *http.Client) *V1SpectroClustersCertificatesRenewParams {
	var ()
	return &V1SpectroClustersCertificatesRenewParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersCertificatesRenewParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters certificates renew operation typically these are written to a http.Request
*/
type V1SpectroClustersCertificatesRenewParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) WithTimeout(timeout time.Duration) *V1SpectroClustersCertificatesRenewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) WithContext(ctx context.Context) *V1SpectroClustersCertificatesRenewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) WithHTTPClient(client *http.Client) *V1SpectroClustersCertificatesRenewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) WithUID(uid string) *V1SpectroClustersCertificatesRenewParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters certificates renew params
func (o *V1SpectroClustersCertificatesRenewParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersCertificatesRenewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
