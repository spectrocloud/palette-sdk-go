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
	"github.com/go-openapi/swag"
)

// NewV1CloudConfigsGenericPoolMachinesListParams creates a new V1CloudConfigsGenericPoolMachinesListParams object
// with the default values initialized.
func NewV1CloudConfigsGenericPoolMachinesListParams() *V1CloudConfigsGenericPoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsGenericPoolMachinesListParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGenericPoolMachinesListParamsWithTimeout creates a new V1CloudConfigsGenericPoolMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGenericPoolMachinesListParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGenericPoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsGenericPoolMachinesListParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1CloudConfigsGenericPoolMachinesListParamsWithContext creates a new V1CloudConfigsGenericPoolMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGenericPoolMachinesListParamsWithContext(ctx context.Context) *V1CloudConfigsGenericPoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsGenericPoolMachinesListParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1CloudConfigsGenericPoolMachinesListParamsWithHTTPClient creates a new V1CloudConfigsGenericPoolMachinesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGenericPoolMachinesListParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGenericPoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsGenericPoolMachinesListParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGenericPoolMachinesListParams contains all the parameters to send to the API endpoint
for the v1 cloud configs generic pool machines list operation typically these are written to a http.Request
*/
type V1CloudConfigsGenericPoolMachinesListParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string
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
	/*MachinePoolName
	  Machine pool name

	*/
	MachinePoolName string
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

// WithTimeout adds the timeout to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithContext(ctx context.Context) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithConfigUID(configUID string) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithContinue adds the continueVar to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithContinue(continueVar *string) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithFields adds the fields to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithFields(fields *string) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilters adds the filters to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithFilters(filters *string) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithLimit(limit *int64) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithOffset adds the offset to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithOffset(offset *int64) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) WithOrderBy(orderBy *string) *V1CloudConfigsGenericPoolMachinesListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the v1 cloud configs generic pool machines list params
func (o *V1CloudConfigsGenericPoolMachinesListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGenericPoolMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configUid
	if err := r.SetPathParam("configUid", o.ConfigUID); err != nil {
		return err
	}

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

	// path param machinePoolName
	if err := r.SetPathParam("machinePoolName", o.MachinePoolName); err != nil {
		return err
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
