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

// NewV1ClusterProfilesUIDSpcDownloadParams creates a new V1ClusterProfilesUIDSpcDownloadParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDSpcDownloadParams() *V1ClusterProfilesUIDSpcDownloadParams {
	var ()
	return &V1ClusterProfilesUIDSpcDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDSpcDownloadParamsWithTimeout creates a new V1ClusterProfilesUIDSpcDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDSpcDownloadParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDSpcDownloadParams {
	var ()
	return &V1ClusterProfilesUIDSpcDownloadParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDSpcDownloadParamsWithContext creates a new V1ClusterProfilesUIDSpcDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDSpcDownloadParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDSpcDownloadParams {
	var ()
	return &V1ClusterProfilesUIDSpcDownloadParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDSpcDownloadParamsWithHTTPClient creates a new V1ClusterProfilesUIDSpcDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDSpcDownloadParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDSpcDownloadParams {
	var ()
	return &V1ClusterProfilesUIDSpcDownloadParams{
		HTTPClient: client,
	}
}

/*V1ClusterProfilesUIDSpcDownloadParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid spc download operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDSpcDownloadParams struct {

	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDSpcDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDSpcDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDSpcDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) WithUID(uid string) *V1ClusterProfilesUIDSpcDownloadParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid spc download params
func (o *V1ClusterProfilesUIDSpcDownloadParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDSpcDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
