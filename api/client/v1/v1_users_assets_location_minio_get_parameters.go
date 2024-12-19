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

// NewV1UsersAssetsLocationMinioGetParams creates a new V1UsersAssetsLocationMinioGetParams object
// with the default values initialized.
func NewV1UsersAssetsLocationMinioGetParams() *V1UsersAssetsLocationMinioGetParams {
	var ()
	return &V1UsersAssetsLocationMinioGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetsLocationMinioGetParamsWithTimeout creates a new V1UsersAssetsLocationMinioGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetsLocationMinioGetParamsWithTimeout(timeout time.Duration) *V1UsersAssetsLocationMinioGetParams {
	var ()
	return &V1UsersAssetsLocationMinioGetParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetsLocationMinioGetParamsWithContext creates a new V1UsersAssetsLocationMinioGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetsLocationMinioGetParamsWithContext(ctx context.Context) *V1UsersAssetsLocationMinioGetParams {
	var ()
	return &V1UsersAssetsLocationMinioGetParams{

		Context: ctx,
	}
}

// NewV1UsersAssetsLocationMinioGetParamsWithHTTPClient creates a new V1UsersAssetsLocationMinioGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetsLocationMinioGetParamsWithHTTPClient(client *http.Client) *V1UsersAssetsLocationMinioGetParams {
	var ()
	return &V1UsersAssetsLocationMinioGetParams{
		HTTPClient: client,
	}
}

/*
V1UsersAssetsLocationMinioGetParams contains all the parameters to send to the API endpoint
for the v1 users assets location minio get operation typically these are written to a http.Request
*/
type V1UsersAssetsLocationMinioGetParams struct {

	/*UID
	  Specify the MinIO location uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) WithTimeout(timeout time.Duration) *V1UsersAssetsLocationMinioGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) WithContext(ctx context.Context) *V1UsersAssetsLocationMinioGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) WithHTTPClient(client *http.Client) *V1UsersAssetsLocationMinioGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) WithUID(uid string) *V1UsersAssetsLocationMinioGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users assets location minio get params
func (o *V1UsersAssetsLocationMinioGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetsLocationMinioGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
