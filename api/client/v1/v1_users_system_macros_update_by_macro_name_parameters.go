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

// NewV1UsersSystemMacrosUpdateByMacroNameParams creates a new V1UsersSystemMacrosUpdateByMacroNameParams object
// with the default values initialized.
func NewV1UsersSystemMacrosUpdateByMacroNameParams() *V1UsersSystemMacrosUpdateByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosUpdateByMacroNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersSystemMacrosUpdateByMacroNameParamsWithTimeout creates a new V1UsersSystemMacrosUpdateByMacroNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersSystemMacrosUpdateByMacroNameParamsWithTimeout(timeout time.Duration) *V1UsersSystemMacrosUpdateByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosUpdateByMacroNameParams{

		timeout: timeout,
	}
}

// NewV1UsersSystemMacrosUpdateByMacroNameParamsWithContext creates a new V1UsersSystemMacrosUpdateByMacroNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersSystemMacrosUpdateByMacroNameParamsWithContext(ctx context.Context) *V1UsersSystemMacrosUpdateByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosUpdateByMacroNameParams{

		Context: ctx,
	}
}

// NewV1UsersSystemMacrosUpdateByMacroNameParamsWithHTTPClient creates a new V1UsersSystemMacrosUpdateByMacroNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersSystemMacrosUpdateByMacroNameParamsWithHTTPClient(client *http.Client) *V1UsersSystemMacrosUpdateByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosUpdateByMacroNameParams{
		HTTPClient: client,
	}
}

/*V1UsersSystemMacrosUpdateByMacroNameParams contains all the parameters to send to the API endpoint
for the v1 users system macros update by macro name operation typically these are written to a http.Request
*/
type V1UsersSystemMacrosUpdateByMacroNameParams struct {

	/*Body*/
	Body *models.V1Macros

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) WithTimeout(timeout time.Duration) *V1UsersSystemMacrosUpdateByMacroNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) WithContext(ctx context.Context) *V1UsersSystemMacrosUpdateByMacroNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) WithHTTPClient(client *http.Client) *V1UsersSystemMacrosUpdateByMacroNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) WithBody(body *models.V1Macros) *V1UsersSystemMacrosUpdateByMacroNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users system macros update by macro name params
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) SetBody(body *models.V1Macros) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersSystemMacrosUpdateByMacroNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
