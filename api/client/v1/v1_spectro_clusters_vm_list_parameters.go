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
	"github.com/go-openapi/swag"
)

// NewV1SpectroClustersVMListParams creates a new V1SpectroClustersVMListParams object
// with the default values initialized.
func NewV1SpectroClustersVMListParams() *V1SpectroClustersVMListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1SpectroClustersVMListParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVMListParamsWithTimeout creates a new V1SpectroClustersVMListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVMListParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVMListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1SpectroClustersVMListParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersVMListParamsWithContext creates a new V1SpectroClustersVMListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVMListParamsWithContext(ctx context.Context) *V1SpectroClustersVMListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1SpectroClustersVMListParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersVMListParamsWithHTTPClient creates a new V1SpectroClustersVMListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVMListParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVMListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1SpectroClustersVMListParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*
V1SpectroClustersVMListParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters VM list operation typically these are written to a http.Request
*/
type V1SpectroClustersVMListParams struct {

	/*Continue
	  continue token to paginate the subsequent data items

	*/
	Continue *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. Default and maximum value of the limit is 50.
	If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results.

	*/
	Limit *int64
	/*Namespace
	  Namespace names, comma separated value (ex: dev,test). If namespace is empty it returns the specific resource under all namespace

	*/
	Namespace []string
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVMListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithContext(ctx context.Context) *V1SpectroClustersVMListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVMListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinue adds the continueVar to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithContinue(continueVar *string) *V1SpectroClustersVMListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithLimit adds the limit to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithLimit(limit *int64) *V1SpectroClustersVMListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithNamespace(namespace []string) *V1SpectroClustersVMListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetNamespace(namespace []string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) WithUID(uid string) *V1SpectroClustersVMListParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters VM list params
func (o *V1SpectroClustersVMListParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVMListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Continue != nil {

		// query param continue
		var qrContinue string
		if o.Continue != nil {
			qrContinue = *o.Continue
		}
		qContinue := qrContinue
		if qContinue != "" {
			if err := r.SetQueryParam("continue", qContinue); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesNamespace := o.Namespace

	joinedNamespace := swag.JoinByFormat(valuesNamespace, "csv")
	// query array param namespace
	if err := r.SetQueryParam("namespace", joinedNamespace...); err != nil {
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
