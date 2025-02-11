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

// NewV1ClusterFeatureLogFetcherLogDownloadParams creates a new V1ClusterFeatureLogFetcherLogDownloadParams object
// with the default values initialized.
func NewV1ClusterFeatureLogFetcherLogDownloadParams() *V1ClusterFeatureLogFetcherLogDownloadParams {
	var ()
	return &V1ClusterFeatureLogFetcherLogDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureLogFetcherLogDownloadParamsWithTimeout creates a new V1ClusterFeatureLogFetcherLogDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureLogFetcherLogDownloadParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureLogFetcherLogDownloadParams {
	var ()
	return &V1ClusterFeatureLogFetcherLogDownloadParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureLogFetcherLogDownloadParamsWithContext creates a new V1ClusterFeatureLogFetcherLogDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureLogFetcherLogDownloadParamsWithContext(ctx context.Context) *V1ClusterFeatureLogFetcherLogDownloadParams {
	var ()
	return &V1ClusterFeatureLogFetcherLogDownloadParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureLogFetcherLogDownloadParamsWithHTTPClient creates a new V1ClusterFeatureLogFetcherLogDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureLogFetcherLogDownloadParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureLogFetcherLogDownloadParams {
	var ()
	return &V1ClusterFeatureLogFetcherLogDownloadParams{
		HTTPClient: client,
	}
}

/*V1ClusterFeatureLogFetcherLogDownloadParams contains all the parameters to send to the API endpoint
for the v1 cluster feature log fetcher log download operation typically these are written to a http.Request
*/
type V1ClusterFeatureLogFetcherLogDownloadParams struct {

	/*FileName*/
	FileName *string
	/*UID
	  Log fetcher uid for which log is requested

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureLogFetcherLogDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) WithContext(ctx context.Context) *V1ClusterFeatureLogFetcherLogDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureLogFetcherLogDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileName adds the fileName to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) WithFileName(fileName *string) *V1ClusterFeatureLogFetcherLogDownloadParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) SetFileName(fileName *string) {
	o.FileName = fileName
}

// WithUID adds the uid to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) WithUID(uid string) *V1ClusterFeatureLogFetcherLogDownloadParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature log fetcher log download params
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureLogFetcherLogDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileName != nil {

		// query param fileName
		var qrFileName string
		if o.FileName != nil {
			qrFileName = *o.FileName
		}
		qFileName := qrFileName
		if qFileName != "" {
			if err := r.SetQueryParam("fileName", qFileName); err != nil {
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
