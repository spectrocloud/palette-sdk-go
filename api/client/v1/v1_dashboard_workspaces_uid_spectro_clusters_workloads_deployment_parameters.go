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

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams object
// with the default values initialized.
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParamsWithTimeout creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParamsWithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams{

		timeout: timeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParamsWithContext creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParamsWithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams{

		Context: ctx,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParamsWithHTTPClient creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParamsWithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams{
		HTTPClient: client,
	}
}

/*V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams contains all the parameters to send to the API endpoint
for the v1 dashboard workspaces Uid spectro clusters workloads deployment operation typically these are written to a http.Request
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams struct {

	/*Body*/
	Body *models.V1WorkspaceWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) WithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) WithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) WithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) WithBody(body *models.V1WorkspaceWorkloadsSpec) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) SetBody(body *models.V1WorkspaceWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) WithUID(uid string) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads deployment params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
