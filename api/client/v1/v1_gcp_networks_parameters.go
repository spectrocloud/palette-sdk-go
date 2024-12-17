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

// NewV1GcpNetworksParams creates a new V1GcpNetworksParams object
// with the default values initialized.
func NewV1GcpNetworksParams() *V1GcpNetworksParams {
	var ()
	return &V1GcpNetworksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GcpNetworksParamsWithTimeout creates a new V1GcpNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GcpNetworksParamsWithTimeout(timeout time.Duration) *V1GcpNetworksParams {
	var ()
	return &V1GcpNetworksParams{

		timeout: timeout,
	}
}

// NewV1GcpNetworksParamsWithContext creates a new V1GcpNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GcpNetworksParamsWithContext(ctx context.Context) *V1GcpNetworksParams {
	var ()
	return &V1GcpNetworksParams{

		Context: ctx,
	}
}

// NewV1GcpNetworksParamsWithHTTPClient creates a new V1GcpNetworksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GcpNetworksParamsWithHTTPClient(client *http.Client) *V1GcpNetworksParams {
	var ()
	return &V1GcpNetworksParams{
		HTTPClient: client,
	}
}

/*V1GcpNetworksParams contains all the parameters to send to the API endpoint
for the v1 gcp networks operation typically these are written to a http.Request
*/
type V1GcpNetworksParams struct {

	/*CloudAccountUID
	  Uid for the specific GCP cloud account

	*/
	CloudAccountUID string
	/*Project
	  Project Name for which GCP networks are requested

	*/
	Project string
	/*Region
	  Region for which GCP networks are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 gcp networks params
func (o *V1GcpNetworksParams) WithTimeout(timeout time.Duration) *V1GcpNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 gcp networks params
func (o *V1GcpNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 gcp networks params
func (o *V1GcpNetworksParams) WithContext(ctx context.Context) *V1GcpNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 gcp networks params
func (o *V1GcpNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 gcp networks params
func (o *V1GcpNetworksParams) WithHTTPClient(client *http.Client) *V1GcpNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 gcp networks params
func (o *V1GcpNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 gcp networks params
func (o *V1GcpNetworksParams) WithCloudAccountUID(cloudAccountUID string) *V1GcpNetworksParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 gcp networks params
func (o *V1GcpNetworksParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithProject adds the project to the v1 gcp networks params
func (o *V1GcpNetworksParams) WithProject(project string) *V1GcpNetworksParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the v1 gcp networks params
func (o *V1GcpNetworksParams) SetProject(project string) {
	o.Project = project
}

// WithRegion adds the region to the v1 gcp networks params
func (o *V1GcpNetworksParams) WithRegion(region string) *V1GcpNetworksParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 gcp networks params
func (o *V1GcpNetworksParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1GcpNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloudAccountUid
	qrCloudAccountUID := o.CloudAccountUID
	qCloudAccountUID := qrCloudAccountUID
	if qCloudAccountUID != "" {
		if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
			return err
		}
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
