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

// NewV1SpectroClustersUIDCostSummaryParams creates a new V1SpectroClustersUIDCostSummaryParams object
// with the default values initialized.
func NewV1SpectroClustersUIDCostSummaryParams() *V1SpectroClustersUIDCostSummaryParams {
	var ()
	return &V1SpectroClustersUIDCostSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDCostSummaryParamsWithTimeout creates a new V1SpectroClustersUIDCostSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDCostSummaryParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDCostSummaryParams {
	var ()
	return &V1SpectroClustersUIDCostSummaryParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDCostSummaryParamsWithContext creates a new V1SpectroClustersUIDCostSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDCostSummaryParamsWithContext(ctx context.Context) *V1SpectroClustersUIDCostSummaryParams {
	var ()
	return &V1SpectroClustersUIDCostSummaryParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDCostSummaryParamsWithHTTPClient creates a new V1SpectroClustersUIDCostSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDCostSummaryParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDCostSummaryParams {
	var ()
	return &V1SpectroClustersUIDCostSummaryParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersUIDCostSummaryParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid cost summary operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDCostSummaryParams struct {

	/*EndTime
	  Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.

	*/
	EndTime *strfmt.DateTime
	/*Period
	  period in minutes, group the data point by the specified period

	*/
	Period *int32
	/*StartTime
	  Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.

	*/
	StartTime *strfmt.DateTime
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDCostSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithContext(ctx context.Context) *V1SpectroClustersUIDCostSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDCostSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithEndTime(endTime *strfmt.DateTime) *V1SpectroClustersUIDCostSummaryParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithPeriod adds the period to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithPeriod(period *int32) *V1SpectroClustersUIDCostSummaryParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetPeriod(period *int32) {
	o.Period = period
}

// WithStartTime adds the startTime to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithStartTime(startTime *strfmt.DateTime) *V1SpectroClustersUIDCostSummaryParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WithUID adds the uid to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) WithUID(uid string) *V1SpectroClustersUIDCostSummaryParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid cost summary params
func (o *V1SpectroClustersUIDCostSummaryParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDCostSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime strfmt.DateTime
		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime.String()
		if qEndTime != "" {
			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}

	}

	if o.Period != nil {

		// query param period
		var qrPeriod int32
		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := swag.FormatInt32(qrPeriod)
		if qPeriod != "" {
			if err := r.SetQueryParam("period", qPeriod); err != nil {
				return err
			}
		}

	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime strfmt.DateTime
		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime.String()
		if qStartTime != "" {
			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
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
