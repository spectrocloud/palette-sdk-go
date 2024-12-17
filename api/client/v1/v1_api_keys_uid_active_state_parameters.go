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

// NewV1APIKeysUIDActiveStateParams creates a new V1APIKeysUIDActiveStateParams object
// with the default values initialized.
func NewV1APIKeysUIDActiveStateParams() *V1APIKeysUIDActiveStateParams {
	var ()
	return &V1APIKeysUIDActiveStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1APIKeysUIDActiveStateParamsWithTimeout creates a new V1APIKeysUIDActiveStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1APIKeysUIDActiveStateParamsWithTimeout(timeout time.Duration) *V1APIKeysUIDActiveStateParams {
	var ()
	return &V1APIKeysUIDActiveStateParams{

		timeout: timeout,
	}
}

// NewV1APIKeysUIDActiveStateParamsWithContext creates a new V1APIKeysUIDActiveStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1APIKeysUIDActiveStateParamsWithContext(ctx context.Context) *V1APIKeysUIDActiveStateParams {
	var ()
	return &V1APIKeysUIDActiveStateParams{

		Context: ctx,
	}
}

// NewV1APIKeysUIDActiveStateParamsWithHTTPClient creates a new V1APIKeysUIDActiveStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1APIKeysUIDActiveStateParamsWithHTTPClient(client *http.Client) *V1APIKeysUIDActiveStateParams {
	var ()
	return &V1APIKeysUIDActiveStateParams{
		HTTPClient: client,
	}
}

/*V1APIKeysUIDActiveStateParams contains all the parameters to send to the API endpoint
for the v1 Api keys Uid active state operation typically these are written to a http.Request
*/
type V1APIKeysUIDActiveStateParams struct {

	/*Body*/
	Body *models.V1APIKeyActiveState
	/*UID
	  Specify API key uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) WithTimeout(timeout time.Duration) *V1APIKeysUIDActiveStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) WithContext(ctx context.Context) *V1APIKeysUIDActiveStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) WithHTTPClient(client *http.Client) *V1APIKeysUIDActiveStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) WithBody(body *models.V1APIKeyActiveState) *V1APIKeysUIDActiveStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) SetBody(body *models.V1APIKeyActiveState) {
	o.Body = body
}

// WithUID adds the uid to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) WithUID(uid string) *V1APIKeysUIDActiveStateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 Api keys Uid active state params
func (o *V1APIKeysUIDActiveStateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1APIKeysUIDActiveStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
