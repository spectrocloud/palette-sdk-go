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
	"github.com/go-openapi/swag"
)

// NewV1UsersListParams creates a new V1UsersListParams object
// with the default values initialized.
func NewV1UsersListParams() *V1UsersListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1UsersListParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersListParamsWithTimeout creates a new V1UsersListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersListParamsWithTimeout(timeout time.Duration) *V1UsersListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1UsersListParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1UsersListParamsWithContext creates a new V1UsersListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersListParamsWithContext(ctx context.Context) *V1UsersListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1UsersListParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1UsersListParamsWithHTTPClient creates a new V1UsersListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersListParamsWithHTTPClient(client *http.Client) *V1UsersListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1UsersListParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*
V1UsersListParams contains all the parameters to send to the API endpoint
for the v1 users list operation typically these are written to a http.Request
*/
type V1UsersListParams struct {

	/*Continue
	  continue token to paginate the subsequent data items

	*/
	Continue *string
	/*Fields
	  Set of fields to be presented in the response with values. The fields are comma separated. Eg: metadata.uid,metadata.name

	*/
	Fields *string
	/*Filters
	  Filters can be combined with AND, OR operators with field path name. Eg: metadata.name=TestServiceANDspec.cloudType=aws

	Server will be restricted to certain fields based on the indexed data for each resource.

	*/
	Filters *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. Default and maximum value of the limit is 50.
	If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results.

	*/
	Limit *int64
	/*Offset
	  offset is the next index number from which the response will start. The response offset value can be used along with continue token for the pagination.

	*/
	Offset *int64
	/*OrderBy
	  Specify the fields with sort order. 1 indicates ascending and -1 for descending. Eg: orderBy=metadata.name=1,metadata.uid=-1

	*/
	OrderBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users list params
func (o *V1UsersListParams) WithTimeout(timeout time.Duration) *V1UsersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users list params
func (o *V1UsersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users list params
func (o *V1UsersListParams) WithContext(ctx context.Context) *V1UsersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users list params
func (o *V1UsersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users list params
func (o *V1UsersListParams) WithHTTPClient(client *http.Client) *V1UsersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users list params
func (o *V1UsersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinue adds the continueVar to the v1 users list params
func (o *V1UsersListParams) WithContinue(continueVar *string) *V1UsersListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 users list params
func (o *V1UsersListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithFields adds the fields to the v1 users list params
func (o *V1UsersListParams) WithFields(fields *string) *V1UsersListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the v1 users list params
func (o *V1UsersListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilters adds the filters to the v1 users list params
func (o *V1UsersListParams) WithFilters(filters *string) *V1UsersListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the v1 users list params
func (o *V1UsersListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the v1 users list params
func (o *V1UsersListParams) WithLimit(limit *int64) *V1UsersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 users list params
func (o *V1UsersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the v1 users list params
func (o *V1UsersListParams) WithOffset(offset *int64) *V1UsersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the v1 users list params
func (o *V1UsersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the v1 users list params
func (o *V1UsersListParams) WithOrderBy(orderBy *string) *V1UsersListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the v1 users list params
func (o *V1UsersListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Continue != nil {

		// query param continue
		var qrContinue string
		if o.Continue != nil {
			qrContinue = *o.Continue
		}
		qContinue := qrContinue
		if qContinue != "" {
			if err := r.SetQueryParam("continue", qContinue); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
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