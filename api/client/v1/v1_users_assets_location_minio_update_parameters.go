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

// NewV1UsersAssetsLocationMinioUpdateParams creates a new V1UsersAssetsLocationMinioUpdateParams object
// with the default values initialized.
func NewV1UsersAssetsLocationMinioUpdateParams() *V1UsersAssetsLocationMinioUpdateParams {
	var ()
	return &V1UsersAssetsLocationMinioUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetsLocationMinioUpdateParamsWithTimeout creates a new V1UsersAssetsLocationMinioUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetsLocationMinioUpdateParamsWithTimeout(timeout time.Duration) *V1UsersAssetsLocationMinioUpdateParams {
	var ()
	return &V1UsersAssetsLocationMinioUpdateParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetsLocationMinioUpdateParamsWithContext creates a new V1UsersAssetsLocationMinioUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetsLocationMinioUpdateParamsWithContext(ctx context.Context) *V1UsersAssetsLocationMinioUpdateParams {
	var ()
	return &V1UsersAssetsLocationMinioUpdateParams{

		Context: ctx,
	}
}

// NewV1UsersAssetsLocationMinioUpdateParamsWithHTTPClient creates a new V1UsersAssetsLocationMinioUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetsLocationMinioUpdateParamsWithHTTPClient(client *http.Client) *V1UsersAssetsLocationMinioUpdateParams {
	var ()
	return &V1UsersAssetsLocationMinioUpdateParams{
		HTTPClient: client,
	}
}

/*
V1UsersAssetsLocationMinioUpdateParams contains all the parameters to send to the API endpoint
for the v1 users assets location minio update operation typically these are written to a http.Request
*/
type V1UsersAssetsLocationMinioUpdateParams struct {

	/*Body*/
	Body *models.V1UserAssetsLocationS3
	/*UID
	  Specify the MinIO location uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) WithTimeout(timeout time.Duration) *V1UsersAssetsLocationMinioUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) WithContext(ctx context.Context) *V1UsersAssetsLocationMinioUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) WithHTTPClient(client *http.Client) *V1UsersAssetsLocationMinioUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) WithBody(body *models.V1UserAssetsLocationS3) *V1UsersAssetsLocationMinioUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) SetBody(body *models.V1UserAssetsLocationS3) {
	o.Body = body
}

// WithUID adds the uid to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) WithUID(uid string) *V1UsersAssetsLocationMinioUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users assets location minio update params
func (o *V1UsersAssetsLocationMinioUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetsLocationMinioUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
