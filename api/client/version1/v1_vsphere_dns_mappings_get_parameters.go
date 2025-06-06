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

// NewV1VsphereDNSMappingsGetParams creates a new V1VsphereDNSMappingsGetParams object
// with the default values initialized.
func NewV1VsphereDNSMappingsGetParams() *V1VsphereDNSMappingsGetParams {
	var ()
	return &V1VsphereDNSMappingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VsphereDNSMappingsGetParamsWithTimeout creates a new V1VsphereDNSMappingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VsphereDNSMappingsGetParamsWithTimeout(timeout time.Duration) *V1VsphereDNSMappingsGetParams {
	var ()
	return &V1VsphereDNSMappingsGetParams{

		timeout: timeout,
	}
}

// NewV1VsphereDNSMappingsGetParamsWithContext creates a new V1VsphereDNSMappingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VsphereDNSMappingsGetParamsWithContext(ctx context.Context) *V1VsphereDNSMappingsGetParams {
	var ()
	return &V1VsphereDNSMappingsGetParams{

		Context: ctx,
	}
}

// NewV1VsphereDNSMappingsGetParamsWithHTTPClient creates a new V1VsphereDNSMappingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VsphereDNSMappingsGetParamsWithHTTPClient(client *http.Client) *V1VsphereDNSMappingsGetParams {
	var ()
	return &V1VsphereDNSMappingsGetParams{
		HTTPClient: client,
	}
}

/*
V1VsphereDNSMappingsGetParams contains all the parameters to send to the API endpoint
for the v1 vsphere Dns mappings get operation typically these are written to a http.Request
*/
type V1VsphereDNSMappingsGetParams struct {

	/*Filters
	  Filters can be combined with AND, OR operators with field path name. Eg: metadata.name=TestServiceANDspec.cloudType=aws

	Server will be restricted to certain fields based on the indexed data for each resource.

	*/
	Filters *string
	/*OrderBy
	  Specify the fields with sort order. 1 indicates ascending and -1 for descending. Eg: orderBy=metadata.name=1,metadata.uid=-1

	*/
	OrderBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) WithTimeout(timeout time.Duration) *V1VsphereDNSMappingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) WithContext(ctx context.Context) *V1VsphereDNSMappingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) WithHTTPClient(client *http.Client) *V1VsphereDNSMappingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) WithFilters(filters *string) *V1VsphereDNSMappingsGetParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithOrderBy adds the orderBy to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) WithOrderBy(orderBy *string) *V1VsphereDNSMappingsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the v1 vsphere Dns mappings get params
func (o *V1VsphereDNSMappingsGetParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *V1VsphereDNSMappingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
