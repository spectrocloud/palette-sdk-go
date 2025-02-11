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

// NewV1EventsComponentsListParams creates a new V1EventsComponentsListParams object
// with the default values initialized.
func NewV1EventsComponentsListParams() *V1EventsComponentsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1EventsComponentsListParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EventsComponentsListParamsWithTimeout creates a new V1EventsComponentsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EventsComponentsListParamsWithTimeout(timeout time.Duration) *V1EventsComponentsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1EventsComponentsListParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1EventsComponentsListParamsWithContext creates a new V1EventsComponentsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EventsComponentsListParamsWithContext(ctx context.Context) *V1EventsComponentsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1EventsComponentsListParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1EventsComponentsListParamsWithHTTPClient creates a new V1EventsComponentsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EventsComponentsListParamsWithHTTPClient(client *http.Client) *V1EventsComponentsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1EventsComponentsListParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*V1EventsComponentsListParams contains all the parameters to send to the API endpoint
for the v1 events components list operation typically these are written to a http.Request
*/
type V1EventsComponentsListParams struct {

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

// WithTimeout adds the timeout to the v1 events components list params
func (o *V1EventsComponentsListParams) WithTimeout(timeout time.Duration) *V1EventsComponentsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 events components list params
func (o *V1EventsComponentsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 events components list params
func (o *V1EventsComponentsListParams) WithContext(ctx context.Context) *V1EventsComponentsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 events components list params
func (o *V1EventsComponentsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 events components list params
func (o *V1EventsComponentsListParams) WithHTTPClient(client *http.Client) *V1EventsComponentsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 events components list params
func (o *V1EventsComponentsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinue adds the continueVar to the v1 events components list params
func (o *V1EventsComponentsListParams) WithContinue(continueVar *string) *V1EventsComponentsListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 events components list params
func (o *V1EventsComponentsListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithFields adds the fields to the v1 events components list params
func (o *V1EventsComponentsListParams) WithFields(fields *string) *V1EventsComponentsListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the v1 events components list params
func (o *V1EventsComponentsListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilters adds the filters to the v1 events components list params
func (o *V1EventsComponentsListParams) WithFilters(filters *string) *V1EventsComponentsListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the v1 events components list params
func (o *V1EventsComponentsListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the v1 events components list params
func (o *V1EventsComponentsListParams) WithLimit(limit *int64) *V1EventsComponentsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 events components list params
func (o *V1EventsComponentsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the v1 events components list params
func (o *V1EventsComponentsListParams) WithOffset(offset *int64) *V1EventsComponentsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the v1 events components list params
func (o *V1EventsComponentsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the v1 events components list params
func (o *V1EventsComponentsListParams) WithOrderBy(orderBy *string) *V1EventsComponentsListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the v1 events components list params
func (o *V1EventsComponentsListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *V1EventsComponentsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
