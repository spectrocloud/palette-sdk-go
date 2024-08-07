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

// NewV1CloudInstanceSpotPriceGetParams creates a new V1CloudInstanceSpotPriceGetParams object
// with the default values initialized.
func NewV1CloudInstanceSpotPriceGetParams() *V1CloudInstanceSpotPriceGetParams {
	var ()
	return &V1CloudInstanceSpotPriceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudInstanceSpotPriceGetParamsWithTimeout creates a new V1CloudInstanceSpotPriceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudInstanceSpotPriceGetParamsWithTimeout(timeout time.Duration) *V1CloudInstanceSpotPriceGetParams {
	var ()
	return &V1CloudInstanceSpotPriceGetParams{

		timeout: timeout,
	}
}

// NewV1CloudInstanceSpotPriceGetParamsWithContext creates a new V1CloudInstanceSpotPriceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudInstanceSpotPriceGetParamsWithContext(ctx context.Context) *V1CloudInstanceSpotPriceGetParams {
	var ()
	return &V1CloudInstanceSpotPriceGetParams{

		Context: ctx,
	}
}

// NewV1CloudInstanceSpotPriceGetParamsWithHTTPClient creates a new V1CloudInstanceSpotPriceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudInstanceSpotPriceGetParamsWithHTTPClient(client *http.Client) *V1CloudInstanceSpotPriceGetParams {
	var ()
	return &V1CloudInstanceSpotPriceGetParams{
		HTTPClient: client,
	}
}

/*V1CloudInstanceSpotPriceGetParams contains all the parameters to send to the API endpoint
for the v1 cloud instance spot price get operation typically these are written to a http.Request
*/
type V1CloudInstanceSpotPriceGetParams struct {

	/*CloudType
	  Cloud type [aws/azure/gcp/tencent]

	*/
	CloudType string
	/*InstanceType
	  Instance type for a specific cloud type

	*/
	InstanceType string
	/*Timestamp
	  Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.

	*/
	Timestamp *strfmt.DateTime
	/*Zone
	  Availability zone for a specific cloud type

	*/
	Zone string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithTimeout(timeout time.Duration) *V1CloudInstanceSpotPriceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithContext(ctx context.Context) *V1CloudInstanceSpotPriceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithHTTPClient(client *http.Client) *V1CloudInstanceSpotPriceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithCloudType(cloudType string) *V1CloudInstanceSpotPriceGetParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithInstanceType adds the instanceType to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithInstanceType(instanceType string) *V1CloudInstanceSpotPriceGetParams {
	o.SetInstanceType(instanceType)
	return o
}

// SetInstanceType adds the instanceType to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetInstanceType(instanceType string) {
	o.InstanceType = instanceType
}

// WithTimestamp adds the timestamp to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithTimestamp(timestamp *strfmt.DateTime) *V1CloudInstanceSpotPriceGetParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetTimestamp(timestamp *strfmt.DateTime) {
	o.Timestamp = timestamp
}

// WithZone adds the zone to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) WithZone(zone string) *V1CloudInstanceSpotPriceGetParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the v1 cloud instance spot price get params
func (o *V1CloudInstanceSpotPriceGetParams) SetZone(zone string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudInstanceSpotPriceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
	}

	// query param instanceType
	qrInstanceType := o.InstanceType
	qInstanceType := qrInstanceType
	if qInstanceType != "" {
		if err := r.SetQueryParam("instanceType", qInstanceType); err != nil {
			return err
		}
	}

	if o.Timestamp != nil {

		// query param timestamp
		var qrTimestamp strfmt.DateTime
		if o.Timestamp != nil {
			qrTimestamp = *o.Timestamp
		}
		qTimestamp := qrTimestamp.String()
		if qTimestamp != "" {
			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}

	}

	// query param zone
	qrZone := o.Zone
	qZone := qrZone
	if qZone != "" {
		if err := r.SetQueryParam("zone", qZone); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
