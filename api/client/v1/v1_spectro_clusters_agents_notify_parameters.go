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

// NewV1SpectroClustersAgentsNotifyParams creates a new V1SpectroClustersAgentsNotifyParams object
// with the default values initialized.
func NewV1SpectroClustersAgentsNotifyParams() *V1SpectroClustersAgentsNotifyParams {
	var ()
	return &V1SpectroClustersAgentsNotifyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersAgentsNotifyParamsWithTimeout creates a new V1SpectroClustersAgentsNotifyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersAgentsNotifyParamsWithTimeout(timeout time.Duration) *V1SpectroClustersAgentsNotifyParams {
	var ()
	return &V1SpectroClustersAgentsNotifyParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersAgentsNotifyParamsWithContext creates a new V1SpectroClustersAgentsNotifyParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersAgentsNotifyParamsWithContext(ctx context.Context) *V1SpectroClustersAgentsNotifyParams {
	var ()
	return &V1SpectroClustersAgentsNotifyParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersAgentsNotifyParamsWithHTTPClient creates a new V1SpectroClustersAgentsNotifyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersAgentsNotifyParamsWithHTTPClient(client *http.Client) *V1SpectroClustersAgentsNotifyParams {
	var ()
	return &V1SpectroClustersAgentsNotifyParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersAgentsNotifyParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters agents notify operation typically these are written to a http.Request
*/
type V1SpectroClustersAgentsNotifyParams struct {

	/*Body*/
	Body *models.V1SpectroClustersAgentsNotifyEntity
	/*MessageKey*/
	MessageKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) WithTimeout(timeout time.Duration) *V1SpectroClustersAgentsNotifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) WithContext(ctx context.Context) *V1SpectroClustersAgentsNotifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) WithHTTPClient(client *http.Client) *V1SpectroClustersAgentsNotifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) WithBody(body *models.V1SpectroClustersAgentsNotifyEntity) *V1SpectroClustersAgentsNotifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) SetBody(body *models.V1SpectroClustersAgentsNotifyEntity) {
	o.Body = body
}

// WithMessageKey adds the messageKey to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) WithMessageKey(messageKey string) *V1SpectroClustersAgentsNotifyParams {
	o.SetMessageKey(messageKey)
	return o
}

// SetMessageKey adds the messageKey to the v1 spectro clusters agents notify params
func (o *V1SpectroClustersAgentsNotifyParams) SetMessageKey(messageKey string) {
	o.MessageKey = messageKey
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersAgentsNotifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param messageKey
	if err := r.SetPathParam("messageKey", o.MessageKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
