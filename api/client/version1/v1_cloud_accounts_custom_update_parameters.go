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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1CloudAccountsCustomUpdateParams creates a new V1CloudAccountsCustomUpdateParams object
// with the default values initialized.
func NewV1CloudAccountsCustomUpdateParams() *V1CloudAccountsCustomUpdateParams {
	var ()
	return &V1CloudAccountsCustomUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsCustomUpdateParamsWithTimeout creates a new V1CloudAccountsCustomUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsCustomUpdateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsCustomUpdateParams {
	var ()
	return &V1CloudAccountsCustomUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsCustomUpdateParamsWithContext creates a new V1CloudAccountsCustomUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsCustomUpdateParamsWithContext(ctx context.Context) *V1CloudAccountsCustomUpdateParams {
	var ()
	return &V1CloudAccountsCustomUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsCustomUpdateParamsWithHTTPClient creates a new V1CloudAccountsCustomUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsCustomUpdateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsCustomUpdateParams {
	var ()
	return &V1CloudAccountsCustomUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsCustomUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts custom update operation typically these are written to a http.Request
*/
type V1CloudAccountsCustomUpdateParams struct {

	/*Body*/
	Body *models.V1CustomAccountEntity
	/*CloudType
	  Custom cloud type

	*/
	CloudType string
	/*UID
	  Custom cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsCustomUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) WithContext(ctx context.Context) *V1CloudAccountsCustomUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsCustomUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) WithBody(body *models.V1CustomAccountEntity) *V1CloudAccountsCustomUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) SetBody(body *models.V1CustomAccountEntity) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) WithCloudType(cloudType string) *V1CloudAccountsCustomUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithUID adds the uid to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) WithUID(uid string) *V1CloudAccountsCustomUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts custom update params
func (o *V1CloudAccountsCustomUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsCustomUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
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
