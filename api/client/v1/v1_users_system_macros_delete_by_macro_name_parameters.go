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

// NewV1UsersSystemMacrosDeleteByMacroNameParams creates a new V1UsersSystemMacrosDeleteByMacroNameParams object
// with the default values initialized.
func NewV1UsersSystemMacrosDeleteByMacroNameParams() *V1UsersSystemMacrosDeleteByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosDeleteByMacroNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersSystemMacrosDeleteByMacroNameParamsWithTimeout creates a new V1UsersSystemMacrosDeleteByMacroNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersSystemMacrosDeleteByMacroNameParamsWithTimeout(timeout time.Duration) *V1UsersSystemMacrosDeleteByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosDeleteByMacroNameParams{

		timeout: timeout,
	}
}

// NewV1UsersSystemMacrosDeleteByMacroNameParamsWithContext creates a new V1UsersSystemMacrosDeleteByMacroNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersSystemMacrosDeleteByMacroNameParamsWithContext(ctx context.Context) *V1UsersSystemMacrosDeleteByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosDeleteByMacroNameParams{

		Context: ctx,
	}
}

// NewV1UsersSystemMacrosDeleteByMacroNameParamsWithHTTPClient creates a new V1UsersSystemMacrosDeleteByMacroNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersSystemMacrosDeleteByMacroNameParamsWithHTTPClient(client *http.Client) *V1UsersSystemMacrosDeleteByMacroNameParams {
	var ()
	return &V1UsersSystemMacrosDeleteByMacroNameParams{
		HTTPClient: client,
	}
}

/*V1UsersSystemMacrosDeleteByMacroNameParams contains all the parameters to send to the API endpoint
for the v1 users system macros delete by macro name operation typically these are written to a http.Request
*/
type V1UsersSystemMacrosDeleteByMacroNameParams struct {

	/*Body*/
	Body *models.V1Macros

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) WithTimeout(timeout time.Duration) *V1UsersSystemMacrosDeleteByMacroNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) WithContext(ctx context.Context) *V1UsersSystemMacrosDeleteByMacroNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) WithHTTPClient(client *http.Client) *V1UsersSystemMacrosDeleteByMacroNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) WithBody(body *models.V1Macros) *V1UsersSystemMacrosDeleteByMacroNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users system macros delete by macro name params
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) SetBody(body *models.V1Macros) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersSystemMacrosDeleteByMacroNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
