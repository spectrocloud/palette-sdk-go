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

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams object
// with the default values initialized.
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParamsWithTimeout creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParamsWithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams{

		timeout: timeout,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParamsWithContext creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParamsWithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams{

		Context: ctx,
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParamsWithHTTPClient creates a new V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParamsWithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	var ()
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams{
		HTTPClient: client,
	}
}

/*
V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams contains all the parameters to send to the API endpoint
for the v1 dashboard workspaces Uid spectro clusters workloads job operation typically these are written to a http.Request
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams struct {

	/*Body*/
	Body *models.V1WorkspaceWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) WithTimeout(timeout time.Duration) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) WithContext(ctx context.Context) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) WithHTTPClient(client *http.Client) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) WithBody(body *models.V1WorkspaceWorkloadsSpec) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) SetBody(body *models.V1WorkspaceWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) WithUID(uid string) *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard workspaces Uid spectro clusters workloads job params
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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