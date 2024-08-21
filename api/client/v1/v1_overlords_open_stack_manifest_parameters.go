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

// NewV1OverlordsOpenStackManifestParams creates a new V1OverlordsOpenStackManifestParams object
// with the default values initialized.
func NewV1OverlordsOpenStackManifestParams() *V1OverlordsOpenStackManifestParams {
	var ()
	return &V1OverlordsOpenStackManifestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsOpenStackManifestParamsWithTimeout creates a new V1OverlordsOpenStackManifestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsOpenStackManifestParamsWithTimeout(timeout time.Duration) *V1OverlordsOpenStackManifestParams {
	var ()
	return &V1OverlordsOpenStackManifestParams{

		timeout: timeout,
	}
}

// NewV1OverlordsOpenStackManifestParamsWithContext creates a new V1OverlordsOpenStackManifestParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsOpenStackManifestParamsWithContext(ctx context.Context) *V1OverlordsOpenStackManifestParams {
	var ()
	return &V1OverlordsOpenStackManifestParams{

		Context: ctx,
	}
}

// NewV1OverlordsOpenStackManifestParamsWithHTTPClient creates a new V1OverlordsOpenStackManifestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsOpenStackManifestParamsWithHTTPClient(client *http.Client) *V1OverlordsOpenStackManifestParams {
	var ()
	return &V1OverlordsOpenStackManifestParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsOpenStackManifestParams contains all the parameters to send to the API endpoint
for the v1 overlords open stack manifest operation typically these are written to a http.Request
*/
type V1OverlordsOpenStackManifestParams struct {

	/*PairingCode*/
	PairingCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) WithTimeout(timeout time.Duration) *V1OverlordsOpenStackManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) WithContext(ctx context.Context) *V1OverlordsOpenStackManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) WithHTTPClient(client *http.Client) *V1OverlordsOpenStackManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPairingCode adds the pairingCode to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) WithPairingCode(pairingCode string) *V1OverlordsOpenStackManifestParams {
	o.SetPairingCode(pairingCode)
	return o
}

// SetPairingCode adds the pairingCode to the v1 overlords open stack manifest params
func (o *V1OverlordsOpenStackManifestParams) SetPairingCode(pairingCode string) {
	o.PairingCode = pairingCode
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsOpenStackManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param pairingCode
	qrPairingCode := o.PairingCode
	qPairingCode := qrPairingCode
	if qPairingCode != "" {
		if err := r.SetQueryParam("pairingCode", qPairingCode); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}