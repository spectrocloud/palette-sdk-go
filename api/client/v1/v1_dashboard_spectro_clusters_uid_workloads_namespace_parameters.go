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

// NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParams creates a new V1DashboardSpectroClustersUIDWorkloadsNamespaceParams object
// with the default values initialized.
func NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParams() *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParamsWithTimeout creates a new V1DashboardSpectroClustersUIDWorkloadsNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParamsWithTimeout(timeout time.Duration) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsNamespaceParams{

		timeout: timeout,
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParamsWithContext creates a new V1DashboardSpectroClustersUIDWorkloadsNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParamsWithContext(ctx context.Context) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsNamespaceParams{

		Context: ctx,
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParamsWithHTTPClient creates a new V1DashboardSpectroClustersUIDWorkloadsNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardSpectroClustersUIDWorkloadsNamespaceParamsWithHTTPClient(client *http.Client) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	var ()
	return &V1DashboardSpectroClustersUIDWorkloadsNamespaceParams{
		HTTPClient: client,
	}
}

/*
V1DashboardSpectroClustersUIDWorkloadsNamespaceParams contains all the parameters to send to the API endpoint
for the v1 dashboard spectro clusters Uid workloads namespace operation typically these are written to a http.Request
*/
type V1DashboardSpectroClustersUIDWorkloadsNamespaceParams struct {

	/*Body*/
	Body *models.V1ClusterWorkloadsSpec
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) WithTimeout(timeout time.Duration) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) WithContext(ctx context.Context) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) WithHTTPClient(client *http.Client) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) WithBody(body *models.V1ClusterWorkloadsSpec) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) SetBody(body *models.V1ClusterWorkloadsSpec) {
	o.Body = body
}

// WithUID adds the uid to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) WithUID(uid string) *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 dashboard spectro clusters Uid workloads namespace params
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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