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

// NewV1OverlordsUIDOpenStackClusterProfileParams creates a new V1OverlordsUIDOpenStackClusterProfileParams object
// with the default values initialized.
func NewV1OverlordsUIDOpenStackClusterProfileParams() *V1OverlordsUIDOpenStackClusterProfileParams {
	var ()
	return &V1OverlordsUIDOpenStackClusterProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDOpenStackClusterProfileParamsWithTimeout creates a new V1OverlordsUIDOpenStackClusterProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDOpenStackClusterProfileParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackClusterProfileParams {
	var ()
	return &V1OverlordsUIDOpenStackClusterProfileParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDOpenStackClusterProfileParamsWithContext creates a new V1OverlordsUIDOpenStackClusterProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDOpenStackClusterProfileParamsWithContext(ctx context.Context) *V1OverlordsUIDOpenStackClusterProfileParams {
	var ()
	return &V1OverlordsUIDOpenStackClusterProfileParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDOpenStackClusterProfileParamsWithHTTPClient creates a new V1OverlordsUIDOpenStackClusterProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDOpenStackClusterProfileParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackClusterProfileParams {
	var ()
	return &V1OverlordsUIDOpenStackClusterProfileParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDOpenStackClusterProfileParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid open stack cluster profile operation typically these are written to a http.Request
*/
type V1OverlordsUIDOpenStackClusterProfileParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackClusterProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) WithContext(ctx context.Context) *V1OverlordsUIDOpenStackClusterProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackClusterProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) WithUID(uid string) *V1OverlordsUIDOpenStackClusterProfileParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid open stack cluster profile params
func (o *V1OverlordsUIDOpenStackClusterProfileParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDOpenStackClusterProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
