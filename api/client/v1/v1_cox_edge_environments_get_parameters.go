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

// NewV1CoxEdgeEnvironmentsGetParams creates a new V1CoxEdgeEnvironmentsGetParams object
// with the default values initialized.
func NewV1CoxEdgeEnvironmentsGetParams() *V1CoxEdgeEnvironmentsGetParams {
	var ()
	return &V1CoxEdgeEnvironmentsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CoxEdgeEnvironmentsGetParamsWithTimeout creates a new V1CoxEdgeEnvironmentsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CoxEdgeEnvironmentsGetParamsWithTimeout(timeout time.Duration) *V1CoxEdgeEnvironmentsGetParams {
	var ()
	return &V1CoxEdgeEnvironmentsGetParams{

		timeout: timeout,
	}
}

// NewV1CoxEdgeEnvironmentsGetParamsWithContext creates a new V1CoxEdgeEnvironmentsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CoxEdgeEnvironmentsGetParamsWithContext(ctx context.Context) *V1CoxEdgeEnvironmentsGetParams {
	var ()
	return &V1CoxEdgeEnvironmentsGetParams{

		Context: ctx,
	}
}

// NewV1CoxEdgeEnvironmentsGetParamsWithHTTPClient creates a new V1CoxEdgeEnvironmentsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CoxEdgeEnvironmentsGetParamsWithHTTPClient(client *http.Client) *V1CoxEdgeEnvironmentsGetParams {
	var ()
	return &V1CoxEdgeEnvironmentsGetParams{
		HTTPClient: client,
	}
}

/*
V1CoxEdgeEnvironmentsGetParams contains all the parameters to send to the API endpoint
for the v1 cox edge environments get operation typically these are written to a http.Request
*/
type V1CoxEdgeEnvironmentsGetParams struct {

	/*CloudAccountUID
	  Uid for the specific CoxEdge cloud account

	*/
	CloudAccountUID string
	/*OrganizationID
	  OrganizationId for the specific CoxEdge account

	*/
	OrganizationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) WithTimeout(timeout time.Duration) *V1CoxEdgeEnvironmentsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) WithContext(ctx context.Context) *V1CoxEdgeEnvironmentsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) WithHTTPClient(client *http.Client) *V1CoxEdgeEnvironmentsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) WithCloudAccountUID(cloudAccountUID string) *V1CoxEdgeEnvironmentsGetParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithOrganizationID adds the organizationID to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) WithOrganizationID(organizationID *string) *V1CoxEdgeEnvironmentsGetParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the v1 cox edge environments get params
func (o *V1CoxEdgeEnvironmentsGetParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CoxEdgeEnvironmentsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloudAccountUid
	qrCloudAccountUID := o.CloudAccountUID
	qCloudAccountUID := qrCloudAccountUID
	if qCloudAccountUID != "" {
		if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
			return err
		}
	}

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID string
		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {
			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
