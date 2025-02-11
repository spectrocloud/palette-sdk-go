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

// NewV1UsersAssetSSHUpdateParams creates a new V1UsersAssetSSHUpdateParams object
// with the default values initialized.
func NewV1UsersAssetSSHUpdateParams() *V1UsersAssetSSHUpdateParams {
	var ()
	return &V1UsersAssetSSHUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetSSHUpdateParamsWithTimeout creates a new V1UsersAssetSSHUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetSSHUpdateParamsWithTimeout(timeout time.Duration) *V1UsersAssetSSHUpdateParams {
	var ()
	return &V1UsersAssetSSHUpdateParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetSSHUpdateParamsWithContext creates a new V1UsersAssetSSHUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetSSHUpdateParamsWithContext(ctx context.Context) *V1UsersAssetSSHUpdateParams {
	var ()
	return &V1UsersAssetSSHUpdateParams{

		Context: ctx,
	}
}

// NewV1UsersAssetSSHUpdateParamsWithHTTPClient creates a new V1UsersAssetSSHUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetSSHUpdateParamsWithHTTPClient(client *http.Client) *V1UsersAssetSSHUpdateParams {
	var ()
	return &V1UsersAssetSSHUpdateParams{
		HTTPClient: client,
	}
}

/*V1UsersAssetSSHUpdateParams contains all the parameters to send to the API endpoint
for the v1 users asset Ssh update operation typically these are written to a http.Request
*/
type V1UsersAssetSSHUpdateParams struct {

	/*Body*/
	Body *models.V1UserAssetSSH
	/*UID
	  Specify the SSH key uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) WithTimeout(timeout time.Duration) *V1UsersAssetSSHUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) WithContext(ctx context.Context) *V1UsersAssetSSHUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) WithHTTPClient(client *http.Client) *V1UsersAssetSSHUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) WithBody(body *models.V1UserAssetSSH) *V1UsersAssetSSHUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) SetBody(body *models.V1UserAssetSSH) {
	o.Body = body
}

// WithUID adds the uid to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) WithUID(uid string) *V1UsersAssetSSHUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users asset Ssh update params
func (o *V1UsersAssetSSHUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetSSHUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
