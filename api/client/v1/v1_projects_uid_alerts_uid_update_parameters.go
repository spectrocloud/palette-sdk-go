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

// NewV1ProjectsUIDAlertsUIDUpdateParams creates a new V1ProjectsUIDAlertsUIDUpdateParams object
// with the default values initialized.
func NewV1ProjectsUIDAlertsUIDUpdateParams() *V1ProjectsUIDAlertsUIDUpdateParams {
	var ()
	return &V1ProjectsUIDAlertsUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsUIDAlertsUIDUpdateParamsWithTimeout creates a new V1ProjectsUIDAlertsUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsUIDAlertsUIDUpdateParamsWithTimeout(timeout time.Duration) *V1ProjectsUIDAlertsUIDUpdateParams {
	var ()
	return &V1ProjectsUIDAlertsUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1ProjectsUIDAlertsUIDUpdateParamsWithContext creates a new V1ProjectsUIDAlertsUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsUIDAlertsUIDUpdateParamsWithContext(ctx context.Context) *V1ProjectsUIDAlertsUIDUpdateParams {
	var ()
	return &V1ProjectsUIDAlertsUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1ProjectsUIDAlertsUIDUpdateParamsWithHTTPClient creates a new V1ProjectsUIDAlertsUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsUIDAlertsUIDUpdateParamsWithHTTPClient(client *http.Client) *V1ProjectsUIDAlertsUIDUpdateParams {
	var ()
	return &V1ProjectsUIDAlertsUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1ProjectsUIDAlertsUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 projects Uid alerts Uid update operation typically these are written to a http.Request
*/
type V1ProjectsUIDAlertsUIDUpdateParams struct {

	/*AlertUID*/
	AlertUID string
	/*Body*/
	Body *models.V1Channel
	/*Component*/
	Component string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithTimeout(timeout time.Duration) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithContext(ctx context.Context) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithHTTPClient(client *http.Client) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertUID adds the alertUID to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithAlertUID(alertUID string) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetAlertUID(alertUID)
	return o
}

// SetAlertUID adds the alertUid to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetAlertUID(alertUID string) {
	o.AlertUID = alertUID
}

// WithBody adds the body to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithBody(body *models.V1Channel) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetBody(body *models.V1Channel) {
	o.Body = body
}

// WithComponent adds the component to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithComponent(component string) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetComponent(component)
	return o
}

// SetComponent adds the component to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetComponent(component string) {
	o.Component = component
}

// WithUID adds the uid to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WithUID(uid string) *V1ProjectsUIDAlertsUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 projects Uid alerts Uid update params
func (o *V1ProjectsUIDAlertsUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsUIDAlertsUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertUid
	if err := r.SetPathParam("alertUid", o.AlertUID); err != nil {
		return err
	}

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