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

// NewV1SpectroClustersSummaryUIDOverviewParams creates a new V1SpectroClustersSummaryUIDOverviewParams object
// with the default values initialized.
func NewV1SpectroClustersSummaryUIDOverviewParams() *V1SpectroClustersSummaryUIDOverviewParams {
	var ()
	return &V1SpectroClustersSummaryUIDOverviewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersSummaryUIDOverviewParamsWithTimeout creates a new V1SpectroClustersSummaryUIDOverviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersSummaryUIDOverviewParamsWithTimeout(timeout time.Duration) *V1SpectroClustersSummaryUIDOverviewParams {
	var ()
	return &V1SpectroClustersSummaryUIDOverviewParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersSummaryUIDOverviewParamsWithContext creates a new V1SpectroClustersSummaryUIDOverviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersSummaryUIDOverviewParamsWithContext(ctx context.Context) *V1SpectroClustersSummaryUIDOverviewParams {
	var ()
	return &V1SpectroClustersSummaryUIDOverviewParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersSummaryUIDOverviewParamsWithHTTPClient creates a new V1SpectroClustersSummaryUIDOverviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersSummaryUIDOverviewParamsWithHTTPClient(client *http.Client) *V1SpectroClustersSummaryUIDOverviewParams {
	var ()
	return &V1SpectroClustersSummaryUIDOverviewParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersSummaryUIDOverviewParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters summary Uid overview operation typically these are written to a http.Request
*/
type V1SpectroClustersSummaryUIDOverviewParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) WithTimeout(timeout time.Duration) *V1SpectroClustersSummaryUIDOverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) WithContext(ctx context.Context) *V1SpectroClustersSummaryUIDOverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) WithHTTPClient(client *http.Client) *V1SpectroClustersSummaryUIDOverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) WithUID(uid string) *V1SpectroClustersSummaryUIDOverviewParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters summary Uid overview params
func (o *V1SpectroClustersSummaryUIDOverviewParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersSummaryUIDOverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
