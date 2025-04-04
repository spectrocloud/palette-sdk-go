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

// NewV1ClusterGroupsUIDProfilesGetParams creates a new V1ClusterGroupsUIDProfilesGetParams object
// with the default values initialized.
func NewV1ClusterGroupsUIDProfilesGetParams() *V1ClusterGroupsUIDProfilesGetParams {
	var ()
	return &V1ClusterGroupsUIDProfilesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterGroupsUIDProfilesGetParamsWithTimeout creates a new V1ClusterGroupsUIDProfilesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterGroupsUIDProfilesGetParamsWithTimeout(timeout time.Duration) *V1ClusterGroupsUIDProfilesGetParams {
	var ()
	return &V1ClusterGroupsUIDProfilesGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterGroupsUIDProfilesGetParamsWithContext creates a new V1ClusterGroupsUIDProfilesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterGroupsUIDProfilesGetParamsWithContext(ctx context.Context) *V1ClusterGroupsUIDProfilesGetParams {
	var ()
	return &V1ClusterGroupsUIDProfilesGetParams{

		Context: ctx,
	}
}

// NewV1ClusterGroupsUIDProfilesGetParamsWithHTTPClient creates a new V1ClusterGroupsUIDProfilesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterGroupsUIDProfilesGetParamsWithHTTPClient(client *http.Client) *V1ClusterGroupsUIDProfilesGetParams {
	var ()
	return &V1ClusterGroupsUIDProfilesGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterGroupsUIDProfilesGetParams contains all the parameters to send to the API endpoint
for the v1 cluster groups Uid profiles get operation typically these are written to a http.Request
*/
type V1ClusterGroupsUIDProfilesGetParams struct {

	/*IncludePackMeta
	  includes pack meta such as schema, presets

	*/
	IncludePackMeta *string
	/*UID
	  ClusterGroup uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) WithTimeout(timeout time.Duration) *V1ClusterGroupsUIDProfilesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) WithContext(ctx context.Context) *V1ClusterGroupsUIDProfilesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) WithHTTPClient(client *http.Client) *V1ClusterGroupsUIDProfilesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludePackMeta adds the includePackMeta to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) WithIncludePackMeta(includePackMeta *string) *V1ClusterGroupsUIDProfilesGetParams {
	o.SetIncludePackMeta(includePackMeta)
	return o
}

// SetIncludePackMeta adds the includePackMeta to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) SetIncludePackMeta(includePackMeta *string) {
	o.IncludePackMeta = includePackMeta
}

// WithUID adds the uid to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) WithUID(uid string) *V1ClusterGroupsUIDProfilesGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster groups Uid profiles get params
func (o *V1ClusterGroupsUIDProfilesGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterGroupsUIDProfilesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludePackMeta != nil {

		// query param includePackMeta
		var qrIncludePackMeta string
		if o.IncludePackMeta != nil {
			qrIncludePackMeta = *o.IncludePackMeta
		}
		qIncludePackMeta := qrIncludePackMeta
		if qIncludePackMeta != "" {
			if err := r.SetQueryParam("includePackMeta", qIncludePackMeta); err != nil {
				return err
			}
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
