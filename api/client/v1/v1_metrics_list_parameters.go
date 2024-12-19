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

// NewV1MetricsListParams creates a new V1MetricsListParams object
// with the default values initialized.
func NewV1MetricsListParams() *V1MetricsListParams {
	var (
		discreteDefault                    = bool(false)
		includeControlPlaneMachinesDefault = bool(false)
		includeMasterMachinesDefault       = bool(false)
		metricKindDefault                  = string("all")
		periodDefault                      = int32(1)
	)
	return &V1MetricsListParams{
		Discrete:                    &discreteDefault,
		IncludeControlPlaneMachines: &includeControlPlaneMachinesDefault,
		IncludeMasterMachines:       &includeMasterMachinesDefault,
		MetricKind:                  &metricKindDefault,
		Period:                      &periodDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MetricsListParamsWithTimeout creates a new V1MetricsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MetricsListParamsWithTimeout(timeout time.Duration) *V1MetricsListParams {
	var (
		discreteDefault                    = bool(false)
		includeControlPlaneMachinesDefault = bool(false)
		includeMasterMachinesDefault       = bool(false)
		metricKindDefault                  = string("all")
		periodDefault                      = int32(1)
	)
	return &V1MetricsListParams{
		Discrete:                    &discreteDefault,
		IncludeControlPlaneMachines: &includeControlPlaneMachinesDefault,
		IncludeMasterMachines:       &includeMasterMachinesDefault,
		MetricKind:                  &metricKindDefault,
		Period:                      &periodDefault,

		timeout: timeout,
	}
}

// NewV1MetricsListParamsWithContext creates a new V1MetricsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MetricsListParamsWithContext(ctx context.Context) *V1MetricsListParams {
	var (
		discreteDefault                    = bool(false)
		includeControlPlaneMachinesDefault = bool(false)
		includeMasterMachinesDefault       = bool(false)
		metricKindDefault                  = string("all")
		periodDefault                      = int32(1)
	)
	return &V1MetricsListParams{
		Discrete:                    &discreteDefault,
		IncludeControlPlaneMachines: &includeControlPlaneMachinesDefault,
		IncludeMasterMachines:       &includeMasterMachinesDefault,
		MetricKind:                  &metricKindDefault,
		Period:                      &periodDefault,

		Context: ctx,
	}
}

// NewV1MetricsListParamsWithHTTPClient creates a new V1MetricsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MetricsListParamsWithHTTPClient(client *http.Client) *V1MetricsListParams {
	var (
		discreteDefault                    = bool(false)
		includeControlPlaneMachinesDefault = bool(false)
		includeMasterMachinesDefault       = bool(false)
		metricKindDefault                  = string("all")
		periodDefault                      = int32(1)
	)
	return &V1MetricsListParams{
		Discrete:                    &discreteDefault,
		IncludeControlPlaneMachines: &includeControlPlaneMachinesDefault,
		IncludeMasterMachines:       &includeMasterMachinesDefault,
		MetricKind:                  &metricKindDefault,
		Period:                      &periodDefault,
		HTTPClient:                  client,
	}
}

/*V1MetricsListParams contains all the parameters to send to the API endpoint
for the v1 metrics list operation typically these are written to a http.Request
*/
type V1MetricsListParams struct {

	/*Discrete
	  if true then api returns only aggregation values, else api returns all data points by default

	*/
	Discrete *bool
	/*EndTime
	  Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.

	*/
	EndTime *strfmt.DateTime
	/*IncludeControlPlaneMachines
	  includeControlPlaneMachines in boolean, group the data point by including control plane nodes if set to true

	*/
	IncludeControlPlaneMachines *bool
	/*IncludeMasterMachines
	  Deprecated. includeMasterMachines in boolean, group the data point by including control plane nodes if set to true

	*/
	IncludeMasterMachines *bool
	/*MetricKind*/
	MetricKind *string
	/*Period*/
	Period *int32
	/*ResourceKind*/
	ResourceKind string
	/*SpectroClusterUID*/
	SpectroClusterUID *string
	/*StartTime
	  Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.

	*/
	StartTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 metrics list params
func (o *V1MetricsListParams) WithTimeout(timeout time.Duration) *V1MetricsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 metrics list params
func (o *V1MetricsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 metrics list params
func (o *V1MetricsListParams) WithContext(ctx context.Context) *V1MetricsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 metrics list params
func (o *V1MetricsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 metrics list params
func (o *V1MetricsListParams) WithHTTPClient(client *http.Client) *V1MetricsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 metrics list params
func (o *V1MetricsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiscrete adds the discrete to the v1 metrics list params
func (o *V1MetricsListParams) WithDiscrete(discrete *bool) *V1MetricsListParams {
	o.SetDiscrete(discrete)
	return o
}

// SetDiscrete adds the discrete to the v1 metrics list params
func (o *V1MetricsListParams) SetDiscrete(discrete *bool) {
	o.Discrete = discrete
}

// WithEndTime adds the endTime to the v1 metrics list params
func (o *V1MetricsListParams) WithEndTime(endTime *strfmt.DateTime) *V1MetricsListParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the v1 metrics list params
func (o *V1MetricsListParams) SetEndTime(endTime *strfmt.DateTime) {
	o.EndTime = endTime
}

// WithIncludeControlPlaneMachines adds the includeControlPlaneMachines to the v1 metrics list params
func (o *V1MetricsListParams) WithIncludeControlPlaneMachines(includeControlPlaneMachines *bool) *V1MetricsListParams {
	o.SetIncludeControlPlaneMachines(includeControlPlaneMachines)
	return o
}

// SetIncludeControlPlaneMachines adds the includeControlPlaneMachines to the v1 metrics list params
func (o *V1MetricsListParams) SetIncludeControlPlaneMachines(includeControlPlaneMachines *bool) {
	o.IncludeControlPlaneMachines = includeControlPlaneMachines
}

// WithIncludeMasterMachines adds the includeMasterMachines to the v1 metrics list params
func (o *V1MetricsListParams) WithIncludeMasterMachines(includeMasterMachines *bool) *V1MetricsListParams {
	o.SetIncludeMasterMachines(includeMasterMachines)
	return o
}

// SetIncludeMasterMachines adds the includeMasterMachines to the v1 metrics list params
func (o *V1MetricsListParams) SetIncludeMasterMachines(includeMasterMachines *bool) {
	o.IncludeMasterMachines = includeMasterMachines
}

// WithMetricKind adds the metricKind to the v1 metrics list params
func (o *V1MetricsListParams) WithMetricKind(metricKind *string) *V1MetricsListParams {
	o.SetMetricKind(metricKind)
	return o
}

// SetMetricKind adds the metricKind to the v1 metrics list params
func (o *V1MetricsListParams) SetMetricKind(metricKind *string) {
	o.MetricKind = metricKind
}

// WithPeriod adds the period to the v1 metrics list params
func (o *V1MetricsListParams) WithPeriod(period *int32) *V1MetricsListParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the v1 metrics list params
func (o *V1MetricsListParams) SetPeriod(period *int32) {
	o.Period = period
}

// WithResourceKind adds the resourceKind to the v1 metrics list params
func (o *V1MetricsListParams) WithResourceKind(resourceKind string) *V1MetricsListParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the v1 metrics list params
func (o *V1MetricsListParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WithSpectroClusterUID adds the spectroClusterUID to the v1 metrics list params
func (o *V1MetricsListParams) WithSpectroClusterUID(spectroClusterUID *string) *V1MetricsListParams {
	o.SetSpectroClusterUID(spectroClusterUID)
	return o
}

// SetSpectroClusterUID adds the spectroClusterUid to the v1 metrics list params
func (o *V1MetricsListParams) SetSpectroClusterUID(spectroClusterUID *string) {
	o.SpectroClusterUID = spectroClusterUID
}

// WithStartTime adds the startTime to the v1 metrics list params
func (o *V1MetricsListParams) WithStartTime(startTime *strfmt.DateTime) *V1MetricsListParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the v1 metrics list params
func (o *V1MetricsListParams) SetStartTime(startTime *strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *V1MetricsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Discrete != nil {

		// query param discrete
		var qrDiscrete bool
		if o.Discrete != nil {
			qrDiscrete = *o.Discrete
		}
		qDiscrete := swag.FormatBool(qrDiscrete)
		if qDiscrete != "" {
			if err := r.SetQueryParam("discrete", qDiscrete); err != nil {
				return err
			}
		}

	}

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

	if o.IncludeControlPlaneMachines != nil {

		// query param includeControlPlaneMachines
		var qrIncludeControlPlaneMachines bool
		if o.IncludeControlPlaneMachines != nil {
			qrIncludeControlPlaneMachines = *o.IncludeControlPlaneMachines
		}
		qIncludeControlPlaneMachines := swag.FormatBool(qrIncludeControlPlaneMachines)
		if qIncludeControlPlaneMachines != "" {
			if err := r.SetQueryParam("includeControlPlaneMachines", qIncludeControlPlaneMachines); err != nil {
				return err
			}
		}

	}

	if o.IncludeMasterMachines != nil {

		// query param includeMasterMachines
		var qrIncludeMasterMachines bool
		if o.IncludeMasterMachines != nil {
			qrIncludeMasterMachines = *o.IncludeMasterMachines
		}
		qIncludeMasterMachines := swag.FormatBool(qrIncludeMasterMachines)
		if qIncludeMasterMachines != "" {
			if err := r.SetQueryParam("includeMasterMachines", qIncludeMasterMachines); err != nil {
				return err
			}
		}

	}

	if o.MetricKind != nil {

		// query param metricKind
		var qrMetricKind string
		if o.MetricKind != nil {
			qrMetricKind = *o.MetricKind
		}
		qMetricKind := qrMetricKind
		if qMetricKind != "" {
			if err := r.SetQueryParam("metricKind", qMetricKind); err != nil {
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

	// path param resourceKind
	if err := r.SetPathParam("resourceKind", o.ResourceKind); err != nil {
		return err
	}

	if o.SpectroClusterUID != nil {

		// query param spectroClusterUid
		var qrSpectroClusterUID string
		if o.SpectroClusterUID != nil {
			qrSpectroClusterUID = *o.SpectroClusterUID
		}
		qSpectroClusterUID := qrSpectroClusterUID
		if qSpectroClusterUID != "" {
			if err := r.SetQueryParam("spectroClusterUid", qSpectroClusterUID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
