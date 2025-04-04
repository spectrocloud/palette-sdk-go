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
)

// NewV1EventsComponentsObjTypeUIDDeleteParams creates a new V1EventsComponentsObjTypeUIDDeleteParams object
// with the default values initialized.
func NewV1EventsComponentsObjTypeUIDDeleteParams() *V1EventsComponentsObjTypeUIDDeleteParams {
	var ()
	return &V1EventsComponentsObjTypeUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EventsComponentsObjTypeUIDDeleteParamsWithTimeout creates a new V1EventsComponentsObjTypeUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EventsComponentsObjTypeUIDDeleteParamsWithTimeout(timeout time.Duration) *V1EventsComponentsObjTypeUIDDeleteParams {
	var ()
	return &V1EventsComponentsObjTypeUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1EventsComponentsObjTypeUIDDeleteParamsWithContext creates a new V1EventsComponentsObjTypeUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EventsComponentsObjTypeUIDDeleteParamsWithContext(ctx context.Context) *V1EventsComponentsObjTypeUIDDeleteParams {
	var ()
	return &V1EventsComponentsObjTypeUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1EventsComponentsObjTypeUIDDeleteParamsWithHTTPClient creates a new V1EventsComponentsObjTypeUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EventsComponentsObjTypeUIDDeleteParamsWithHTTPClient(client *http.Client) *V1EventsComponentsObjTypeUIDDeleteParams {
	var ()
	return &V1EventsComponentsObjTypeUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1EventsComponentsObjTypeUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 events components obj type Uid delete operation typically these are written to a http.Request
*/
type V1EventsComponentsObjTypeUIDDeleteParams struct {

	/*ObjectKind
	  Describes the related object uid for which events has to be fetched

	*/
	ObjectKind string
	/*ObjectUID
	  Describes the related object kind for which events has to be fetched

	*/
	ObjectUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) WithTimeout(timeout time.Duration) *V1EventsComponentsObjTypeUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) WithContext(ctx context.Context) *V1EventsComponentsObjTypeUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) WithHTTPClient(client *http.Client) *V1EventsComponentsObjTypeUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithObjectKind adds the objectKind to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) WithObjectKind(objectKind string) *V1EventsComponentsObjTypeUIDDeleteParams {
	o.SetObjectKind(objectKind)
	return o
}

// SetObjectKind adds the objectKind to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) SetObjectKind(objectKind string) {
	o.ObjectKind = objectKind
}

// WithObjectUID adds the objectUID to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) WithObjectUID(objectUID string) *V1EventsComponentsObjTypeUIDDeleteParams {
	o.SetObjectUID(objectUID)
	return o
}

// SetObjectUID adds the objectUid to the v1 events components obj type Uid delete params
func (o *V1EventsComponentsObjTypeUIDDeleteParams) SetObjectUID(objectUID string) {
	o.ObjectUID = objectUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1EventsComponentsObjTypeUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param objectKind
	if err := r.SetPathParam("objectKind", o.ObjectKind); err != nil {
		return err
	}

	// path param objectUid
	if err := r.SetPathParam("objectUid", o.ObjectUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
