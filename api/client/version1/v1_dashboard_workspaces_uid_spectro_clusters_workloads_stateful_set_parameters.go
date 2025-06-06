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

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams object
// with the default values initialized.
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParamsWithTimeout creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParamsWithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams{

		timeout: timeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParamsWithContext creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParamsWithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams{

		Context: ctx,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParamsWithHTTPClient creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParamsWithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams{
		HTTPClient: client,
	}
}

/*
V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams contains all the parameters to send to the API endpoint
for the v1 dashboard workspaces Uid spectro clusters workloads stateful set operation typically these are written to a http.Request
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams struct {

	/*Body*/
	Body *models.V1WorkspaceWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) WithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) WithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) WithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) WithBody(body *models.V1WorkspaceWorkloadsSpec) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) SetBody(body *models.V1WorkspaceWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) WithUID(uid string) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads stateful set params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsStatefulSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
