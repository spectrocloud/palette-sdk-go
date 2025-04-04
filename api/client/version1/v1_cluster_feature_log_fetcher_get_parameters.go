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

// NewV1ClusterFeatureLogFetcherGetParams creates a new V1ClusterFeatureLogFetcherGetParams object
// with the default values initialized.
func NewV1ClusterFeatureLogFetcherGetParams() *V1ClusterFeatureLogFetcherGetParams {
	var ()
	return &V1ClusterFeatureLogFetcherGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureLogFetcherGetParamsWithTimeout creates a new V1ClusterFeatureLogFetcherGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureLogFetcherGetParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureLogFetcherGetParams {
	var ()
	return &V1ClusterFeatureLogFetcherGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureLogFetcherGetParamsWithContext creates a new V1ClusterFeatureLogFetcherGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureLogFetcherGetParamsWithContext(ctx context.Context) *V1ClusterFeatureLogFetcherGetParams {
	var ()
	return &V1ClusterFeatureLogFetcherGetParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureLogFetcherGetParamsWithHTTPClient creates a new V1ClusterFeatureLogFetcherGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureLogFetcherGetParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureLogFetcherGetParams {
	var ()
	return &V1ClusterFeatureLogFetcherGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureLogFetcherGetParams contains all the parameters to send to the API endpoint
for the v1 cluster feature log fetcher get operation typically these are written to a http.Request
*/
type V1ClusterFeatureLogFetcherGetParams struct {

	/*RequestID*/
	RequestID *string
	/*UID
	  Cluster uid for which log is requested

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureLogFetcherGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) WithContext(ctx context.Context) *V1ClusterFeatureLogFetcherGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureLogFetcherGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestID adds the requestID to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) WithRequestID(requestID *string) *V1ClusterFeatureLogFetcherGetParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) SetRequestID(requestID *string) {
	o.RequestID = requestID
}

// WithUID adds the uid to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) WithUID(uid string) *V1ClusterFeatureLogFetcherGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature log fetcher get params
func (o *V1ClusterFeatureLogFetcherGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureLogFetcherGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RequestID != nil {

		// query param requestId
		var qrRequestID string
		if o.RequestID != nil {
			qrRequestID = *o.RequestID
		}
		qRequestID := qrRequestID
		if qRequestID != "" {
			if err := r.SetQueryParam("requestId", qRequestID); err != nil {
				return err
			}
		}

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
