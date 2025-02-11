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

// NewV1ProjectsUIDAlertUpdateParams creates a new V1ProjectsUIDAlertUpdateParams object
// with the default values initialized.
func NewV1ProjectsUIDAlertUpdateParams() *V1ProjectsUIDAlertUpdateParams {
	var ()
	return &V1ProjectsUIDAlertUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsUIDAlertUpdateParamsWithTimeout creates a new V1ProjectsUIDAlertUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsUIDAlertUpdateParamsWithTimeout(timeout time.Duration) *V1ProjectsUIDAlertUpdateParams {
	var ()
	return &V1ProjectsUIDAlertUpdateParams{

		timeout: timeout,
	}
}

// NewV1ProjectsUIDAlertUpdateParamsWithContext creates a new V1ProjectsUIDAlertUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsUIDAlertUpdateParamsWithContext(ctx context.Context) *V1ProjectsUIDAlertUpdateParams {
	var ()
	return &V1ProjectsUIDAlertUpdateParams{

		Context: ctx,
	}
}

// NewV1ProjectsUIDAlertUpdateParamsWithHTTPClient creates a new V1ProjectsUIDAlertUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsUIDAlertUpdateParamsWithHTTPClient(client *http.Client) *V1ProjectsUIDAlertUpdateParams {
	var ()
	return &V1ProjectsUIDAlertUpdateParams{
		HTTPClient: client,
	}
}

/*V1ProjectsUIDAlertUpdateParams contains all the parameters to send to the API endpoint
for the v1 projects Uid alert update operation typically these are written to a http.Request
*/
type V1ProjectsUIDAlertUpdateParams struct {

	/*Body*/
	Body *models.V1AlertEntity
	/*Component*/
	Component string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) WithTimeout(timeout time.Duration) *V1ProjectsUIDAlertUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) WithContext(ctx context.Context) *V1ProjectsUIDAlertUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) WithHTTPClient(client *http.Client) *V1ProjectsUIDAlertUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) WithBody(body *models.V1AlertEntity) *V1ProjectsUIDAlertUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) SetBody(body *models.V1AlertEntity) {
	o.Body = body
}

// WithComponent adds the component to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) WithComponent(component string) *V1ProjectsUIDAlertUpdateParams {
	o.SetComponent(component)
	return o
}

// SetComponent adds the component to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) SetComponent(component string) {
	o.Component = component
}

// WithUID adds the uid to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) WithUID(uid string) *V1ProjectsUIDAlertUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 projects Uid alert update params
func (o *V1ProjectsUIDAlertUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsUIDAlertUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component
	if err := r.SetPathParam("component", o.Component); err != nil {
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
