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

// NewV1ProjectsUIDAlertsUIDDeleteParams creates a new V1ProjectsUIDAlertsUIDDeleteParams object
// with the default values initialized.
func NewV1ProjectsUIDAlertsUIDDeleteParams() *V1ProjectsUIDAlertsUIDDeleteParams {
	var ()
	return &V1ProjectsUIDAlertsUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsUIDAlertsUIDDeleteParamsWithTimeout creates a new V1ProjectsUIDAlertsUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsUIDAlertsUIDDeleteParamsWithTimeout(timeout time.Duration) *V1ProjectsUIDAlertsUIDDeleteParams {
	var ()
	return &V1ProjectsUIDAlertsUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1ProjectsUIDAlertsUIDDeleteParamsWithContext creates a new V1ProjectsUIDAlertsUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsUIDAlertsUIDDeleteParamsWithContext(ctx context.Context) *V1ProjectsUIDAlertsUIDDeleteParams {
	var ()
	return &V1ProjectsUIDAlertsUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1ProjectsUIDAlertsUIDDeleteParamsWithHTTPClient creates a new V1ProjectsUIDAlertsUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsUIDAlertsUIDDeleteParamsWithHTTPClient(client *http.Client) *V1ProjectsUIDAlertsUIDDeleteParams {
	var ()
	return &V1ProjectsUIDAlertsUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1ProjectsUIDAlertsUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 projects Uid alerts Uid delete operation typically these are written to a http.Request
*/
type V1ProjectsUIDAlertsUIDDeleteParams struct {

	/*AlertUID*/
	AlertUID string
	/*Component*/
	Component string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WithTimeout(timeout time.Duration) *V1ProjectsUIDAlertsUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WithContext(ctx context.Context) *V1ProjectsUIDAlertsUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WithHTTPClient(client *http.Client) *V1ProjectsUIDAlertsUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertUID adds the alertUID to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WithAlertUID(alertUID string) *V1ProjectsUIDAlertsUIDDeleteParams {
	o.SetAlertUID(alertUID)
	return o
}

// SetAlertUID adds the alertUid to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) SetAlertUID(alertUID string) {
	o.AlertUID = alertUID
}

// WithComponent adds the component to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WithComponent(component string) *V1ProjectsUIDAlertsUIDDeleteParams {
	o.SetComponent(component)
	return o
}

// SetComponent adds the component to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) SetComponent(component string) {
	o.Component = component
}

// WithUID adds the uid to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WithUID(uid string) *V1ProjectsUIDAlertsUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 projects Uid alerts Uid delete params
func (o *V1ProjectsUIDAlertsUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsUIDAlertsUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertUid
	if err := r.SetPathParam("alertUid", o.AlertUID); err != nil {
		return err
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
