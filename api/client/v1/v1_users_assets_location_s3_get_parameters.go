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

// NewV1UsersAssetsLocationS3GetParams creates a new V1UsersAssetsLocationS3GetParams object
// with the default values initialized.
func NewV1UsersAssetsLocationS3GetParams() *V1UsersAssetsLocationS3GetParams {
	var ()
	return &V1UsersAssetsLocationS3GetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetsLocationS3GetParamsWithTimeout creates a new V1UsersAssetsLocationS3GetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetsLocationS3GetParamsWithTimeout(timeout time.Duration) *V1UsersAssetsLocationS3GetParams {
	var ()
	return &V1UsersAssetsLocationS3GetParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetsLocationS3GetParamsWithContext creates a new V1UsersAssetsLocationS3GetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetsLocationS3GetParamsWithContext(ctx context.Context) *V1UsersAssetsLocationS3GetParams {
	var ()
	return &V1UsersAssetsLocationS3GetParams{

		Context: ctx,
	}
}

// NewV1UsersAssetsLocationS3GetParamsWithHTTPClient creates a new V1UsersAssetsLocationS3GetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetsLocationS3GetParamsWithHTTPClient(client *http.Client) *V1UsersAssetsLocationS3GetParams {
	var ()
	return &V1UsersAssetsLocationS3GetParams{
		HTTPClient: client,
	}
}

/*
V1UsersAssetsLocationS3GetParams contains all the parameters to send to the API endpoint
for the v1 users assets location s3 get operation typically these are written to a http.Request
*/
type V1UsersAssetsLocationS3GetParams struct {

	/*UID
	  Specify the S3 location uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) WithTimeout(timeout time.Duration) *V1UsersAssetsLocationS3GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) WithContext(ctx context.Context) *V1UsersAssetsLocationS3GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) WithHTTPClient(client *http.Client) *V1UsersAssetsLocationS3GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) WithUID(uid string) *V1UsersAssetsLocationS3GetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users assets location s3 get params
func (o *V1UsersAssetsLocationS3GetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetsLocationS3GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
