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

// NewV1CloudAccountsOpenStackUpdateParams creates a new V1CloudAccountsOpenStackUpdateParams object
// with the default values initialized.
func NewV1CloudAccountsOpenStackUpdateParams() *V1CloudAccountsOpenStackUpdateParams {
	var ()
	return &V1CloudAccountsOpenStackUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsOpenStackUpdateParamsWithTimeout creates a new V1CloudAccountsOpenStackUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsOpenStackUpdateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsOpenStackUpdateParams {
	var ()
	return &V1CloudAccountsOpenStackUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsOpenStackUpdateParamsWithContext creates a new V1CloudAccountsOpenStackUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsOpenStackUpdateParamsWithContext(ctx context.Context) *V1CloudAccountsOpenStackUpdateParams {
	var ()
	return &V1CloudAccountsOpenStackUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsOpenStackUpdateParamsWithHTTPClient creates a new V1CloudAccountsOpenStackUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsOpenStackUpdateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsOpenStackUpdateParams {
	var ()
	return &V1CloudAccountsOpenStackUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsOpenStackUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts open stack update operation typically these are written to a http.Request
*/
type V1CloudAccountsOpenStackUpdateParams struct {

	/*Body*/
	Body *models.V1OpenStackAccount
	/*UID
	  OpenStack cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsOpenStackUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) WithContext(ctx context.Context) *V1CloudAccountsOpenStackUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsOpenStackUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) WithBody(body *models.V1OpenStackAccount) *V1CloudAccountsOpenStackUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) SetBody(body *models.V1OpenStackAccount) {
	o.Body = body
}

// WithUID adds the uid to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) WithUID(uid string) *V1CloudAccountsOpenStackUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts open stack update params
func (o *V1CloudAccountsOpenStackUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsOpenStackUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
