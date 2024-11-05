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

// NewV1MaasAccountsUIDAzsParams creates a new V1MaasAccountsUIDAzsParams object
// with the default values initialized.
func NewV1MaasAccountsUIDAzsParams() *V1MaasAccountsUIDAzsParams {
	var ()
	return &V1MaasAccountsUIDAzsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MaasAccountsUIDAzsParamsWithTimeout creates a new V1MaasAccountsUIDAzsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MaasAccountsUIDAzsParamsWithTimeout(timeout time.Duration) *V1MaasAccountsUIDAzsParams {
	var ()
	return &V1MaasAccountsUIDAzsParams{

		timeout: timeout,
	}
}

// NewV1MaasAccountsUIDAzsParamsWithContext creates a new V1MaasAccountsUIDAzsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MaasAccountsUIDAzsParamsWithContext(ctx context.Context) *V1MaasAccountsUIDAzsParams {
	var ()
	return &V1MaasAccountsUIDAzsParams{

		Context: ctx,
	}
}

// NewV1MaasAccountsUIDAzsParamsWithHTTPClient creates a new V1MaasAccountsUIDAzsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MaasAccountsUIDAzsParamsWithHTTPClient(client *http.Client) *V1MaasAccountsUIDAzsParams {
	var ()
	return &V1MaasAccountsUIDAzsParams{
		HTTPClient: client,
	}
}

/*V1MaasAccountsUIDAzsParams contains all the parameters to send to the API endpoint
for the v1 maas accounts Uid azs operation typically these are written to a http.Request
*/
type V1MaasAccountsUIDAzsParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) WithTimeout(timeout time.Duration) *V1MaasAccountsUIDAzsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) WithContext(ctx context.Context) *V1MaasAccountsUIDAzsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) WithHTTPClient(client *http.Client) *V1MaasAccountsUIDAzsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) WithUID(uid string) *V1MaasAccountsUIDAzsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 maas accounts Uid azs params
func (o *V1MaasAccountsUIDAzsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1MaasAccountsUIDAzsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
