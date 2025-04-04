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

// NewV1UsersAssetsSSHGetParams creates a new V1UsersAssetsSSHGetParams object
// with the default values initialized.
func NewV1UsersAssetsSSHGetParams() *V1UsersAssetsSSHGetParams {
	var ()
	return &V1UsersAssetsSSHGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetsSSHGetParamsWithTimeout creates a new V1UsersAssetsSSHGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetsSSHGetParamsWithTimeout(timeout time.Duration) *V1UsersAssetsSSHGetParams {
	var ()
	return &V1UsersAssetsSSHGetParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetsSSHGetParamsWithContext creates a new V1UsersAssetsSSHGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetsSSHGetParamsWithContext(ctx context.Context) *V1UsersAssetsSSHGetParams {
	var ()
	return &V1UsersAssetsSSHGetParams{

		Context: ctx,
	}
}

// NewV1UsersAssetsSSHGetParamsWithHTTPClient creates a new V1UsersAssetsSSHGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetsSSHGetParamsWithHTTPClient(client *http.Client) *V1UsersAssetsSSHGetParams {
	var ()
	return &V1UsersAssetsSSHGetParams{
		HTTPClient: client,
	}
}

/*
V1UsersAssetsSSHGetParams contains all the parameters to send to the API endpoint
for the v1 users assets Ssh get operation typically these are written to a http.Request
*/
type V1UsersAssetsSSHGetParams struct {

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

// WithTimeout adds the timeout to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) WithTimeout(timeout time.Duration) *V1UsersAssetsSSHGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) WithContext(ctx context.Context) *V1UsersAssetsSSHGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) WithHTTPClient(client *http.Client) *V1UsersAssetsSSHGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) WithFilters(filters *string) *V1UsersAssetsSSHGetParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithOrderBy adds the orderBy to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) WithOrderBy(orderBy *string) *V1UsersAssetsSSHGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the v1 users assets Ssh get params
func (o *V1UsersAssetsSSHGetParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetsSSHGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
