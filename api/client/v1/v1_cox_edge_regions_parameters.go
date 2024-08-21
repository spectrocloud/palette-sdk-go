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

// NewV1CoxEdgeRegionsParams creates a new V1CoxEdgeRegionsParams object
// with the default values initialized.
func NewV1CoxEdgeRegionsParams() *V1CoxEdgeRegionsParams {
	var ()
	return &V1CoxEdgeRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CoxEdgeRegionsParamsWithTimeout creates a new V1CoxEdgeRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CoxEdgeRegionsParamsWithTimeout(timeout time.Duration) *V1CoxEdgeRegionsParams {
	var ()
	return &V1CoxEdgeRegionsParams{

		timeout: timeout,
	}
}

// NewV1CoxEdgeRegionsParamsWithContext creates a new V1CoxEdgeRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CoxEdgeRegionsParamsWithContext(ctx context.Context) *V1CoxEdgeRegionsParams {
	var ()
	return &V1CoxEdgeRegionsParams{

		Context: ctx,
	}
}

// NewV1CoxEdgeRegionsParamsWithHTTPClient creates a new V1CoxEdgeRegionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CoxEdgeRegionsParamsWithHTTPClient(client *http.Client) *V1CoxEdgeRegionsParams {
	var ()
	return &V1CoxEdgeRegionsParams{
		HTTPClient: client,
	}
}

/*
V1CoxEdgeRegionsParams contains all the parameters to send to the API endpoint
for the v1 cox edge regions operation typically these are written to a http.Request
*/
type V1CoxEdgeRegionsParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID *string
	/*Environment
	  CoxEdge environment name

	*/
	Environment *string
	/*OrganizationID
	  CoxEdge organization id

	*/
	OrganizationID *string
	/*Service
	  CoxEdge service name

	*/
	Service *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithTimeout(timeout time.Duration) *V1CoxEdgeRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithContext(ctx context.Context) *V1CoxEdgeRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithHTTPClient(client *http.Client) *V1CoxEdgeRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithCloudAccountUID(cloudAccountUID *string) *V1CoxEdgeRegionsParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithEnvironment adds the environment to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithEnvironment(environment *string) *V1CoxEdgeRegionsParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithOrganizationID adds the organizationID to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithOrganizationID(organizationID *string) *V1CoxEdgeRegionsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithService adds the service to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) WithService(service *string) *V1CoxEdgeRegionsParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the v1 cox edge regions params
func (o *V1CoxEdgeRegionsParams) SetService(service *string) {
	o.Service = service
}

// WriteToRequest writes these params to a swagger request
func (o *V1CoxEdgeRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudAccountUID != nil {

		// query param cloudAccountUid
		var qrCloudAccountUID string
		if o.CloudAccountUID != nil {
			qrCloudAccountUID = *o.CloudAccountUID
		}
		qCloudAccountUID := qrCloudAccountUID
		if qCloudAccountUID != "" {
			if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
				return err
			}
		}

	}

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string
		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {
			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
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

	if o.Service != nil {

		// query param service
		var qrService string
		if o.Service != nil {
			qrService = *o.Service
		}
		qService := qrService
		if qService != "" {
			if err := r.SetQueryParam("service", qService); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}