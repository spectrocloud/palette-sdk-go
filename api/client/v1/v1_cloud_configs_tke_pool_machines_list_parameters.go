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

// NewV1CloudConfigsTkePoolMachinesListParams creates a new V1CloudConfigsTkePoolMachinesListParams object
// with the default values initialized.
func NewV1CloudConfigsTkePoolMachinesListParams() *V1CloudConfigsTkePoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsTkePoolMachinesListParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsTkePoolMachinesListParamsWithTimeout creates a new V1CloudConfigsTkePoolMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsTkePoolMachinesListParamsWithTimeout(timeout time.Duration) *V1CloudConfigsTkePoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsTkePoolMachinesListParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1CloudConfigsTkePoolMachinesListParamsWithContext creates a new V1CloudConfigsTkePoolMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsTkePoolMachinesListParamsWithContext(ctx context.Context) *V1CloudConfigsTkePoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsTkePoolMachinesListParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1CloudConfigsTkePoolMachinesListParamsWithHTTPClient creates a new V1CloudConfigsTkePoolMachinesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsTkePoolMachinesListParamsWithHTTPClient(client *http.Client) *V1CloudConfigsTkePoolMachinesListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1CloudConfigsTkePoolMachinesListParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*
V1CloudConfigsTkePoolMachinesListParams contains all the parameters to send to the API endpoint
for the v1 cloud configs tke pool machines list operation typically these are written to a http.Request
*/
type V1CloudConfigsTkePoolMachinesListParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithTimeout(timeout time.Duration) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithContext(ctx context.Context) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithHTTPClient(client *http.Client) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithConfigUID(configUID string) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithContinue adds the continueVar to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithContinue(continueVar *string) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithFields adds the fields to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithFields(fields *string) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilters adds the filters to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithFilters(filters *string) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithLimit(limit *int64) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithOffset adds the offset to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithOffset(offset *int64) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrderBy adds the orderBy to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) WithOrderBy(orderBy *string) *V1CloudConfigsTkePoolMachinesListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the v1 cloud configs tke pool machines list params
func (o *V1CloudConfigsTkePoolMachinesListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsTkePoolMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
