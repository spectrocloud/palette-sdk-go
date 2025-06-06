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

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams object
// with the default values initialized.
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParamsWithTimeout creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParamsWithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams{

		timeout: timeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParamsWithContext creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParamsWithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams{

		Context: ctx,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParamsWithHTTPClient creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParamsWithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams{
		HTTPClient: client,
	}
}

/*
V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams contains all the parameters to send to the API endpoint
for the v1 dashboard workspaces Uid spectro clusters workloads role binding operation typically these are written to a http.Request
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams struct {

	/*Body*/
	Body *models.V1WorkspaceWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) WithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) WithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) WithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) WithBody(body *models.V1WorkspaceWorkloadsSpec) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) SetBody(body *models.V1WorkspaceWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) WithUID(uid string) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads role binding params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
