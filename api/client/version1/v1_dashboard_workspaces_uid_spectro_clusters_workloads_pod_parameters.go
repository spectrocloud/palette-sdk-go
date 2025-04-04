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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams object
// with the default values initialized.
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParamsWithTimeout creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParamsWithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams{

		timeout: timeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParamsWithContext creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParamsWithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams{

		Context: ctx,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParamsWithHTTPClient creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParamsWithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams{
		HTTPClient: client,
	}
}

/*
V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams contains all the parameters to send to the API endpoint
for the v1 dashboard workspaces Uid spectro clusters workloads pod operation typically these are written to a http.Request
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams struct {

	/*Body*/
	Body *models.V1WorkspaceWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) WithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) WithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) WithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) WithBody(body *models.V1WorkspaceWorkloadsSpec) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) SetBody(body *models.V1WorkspaceWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) WithUID(uid string) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads pod params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
