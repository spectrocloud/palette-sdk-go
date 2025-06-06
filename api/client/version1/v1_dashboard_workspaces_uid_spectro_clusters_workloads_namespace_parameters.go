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

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams object
// with the default values initialized.
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParamsWithTimeout creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParamsWithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams{

		timeout: timeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParamsWithContext creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParamsWithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams{

		Context: ctx,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParamsWithHTTPClient creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParamsWithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams{
		HTTPClient: client,
	}
}

/*
V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams contains all the parameters to send to the API endpoint
for the v1 dashboard workspaces Uid spectro clusters workloads namespace operation typically these are written to a http.Request
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams struct {

	/*Body*/
	Body *models.V1WorkspaceWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) WithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) WithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) WithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) WithBody(body *models.V1WorkspaceWorkloadsSpec) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) SetBody(body *models.V1WorkspaceWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) WithUID(uid string) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads namespace params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
