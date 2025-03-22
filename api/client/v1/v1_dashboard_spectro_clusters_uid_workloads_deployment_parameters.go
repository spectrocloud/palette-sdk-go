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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParams creates a new V1DashboardSpectroClustersUIDWorkloadsDeploymentParams object
// with the default values initialized.
func NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParams() *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParamsWithTimeout creates a new V1DashboardSpectroClustersUIDWorkloadsDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParamsWithTimeout(timeout time.Duration) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsDeploymentParams{

		timeout: timeout,
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParamsWithContext creates a new V1DashboardSpectroClustersUIDWorkloadsDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParamsWithContext(ctx context.Context) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsDeploymentParams{

		Context: ctx,
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParamsWithHTTPClient creates a new V1DashboardSpectroClustersUIDWorkloadsDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardSpectroClustersUIDWorkloadsDeploymentParamsWithHTTPClient(client *http.Client) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsDeploymentParams{
		HTTPClient: client,
	}
}

/*V1DashboardSpectroClustersUIDWorkloadsDeploymentParams contains all the parameters to send to the API endpoint
for the v1 dashboard spectro clusters Uid workloads deployment operation typically these are written to a http.Request
*/
type V1DashboardSpectroClustersUIDWorkloadsDeploymentParams struct {

	/*Body*/
	Body *models.V1ClusterWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) WithTimeout(timeout time.Duration) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) WithContext(ctx context.Context) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) WithHTTPClient(client *http.Client) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) WithBody(body *models.V1ClusterWorkloadsSpec) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) SetBody(body *models.V1ClusterWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) WithUID(uid string) *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard spectro clusters Uid workloads deployment params
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardSpectroClustersUIDWorkloadsDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
