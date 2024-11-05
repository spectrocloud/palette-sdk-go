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

// NewV1VirtualClustersPacksValuesParams creates a new V1VirtualClustersPacksValuesParams object
// with the default values initialized.
func NewV1VirtualClustersPacksValuesParams() *V1VirtualClustersPacksValuesParams {
	var (
		kubernetesDistroTypeDefault = string("k3s")
	)
	return &V1VirtualClustersPacksValuesParams{
		KubernetesDistroType: &kubernetesDistroTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VirtualClustersPacksValuesParamsWithTimeout creates a new V1VirtualClustersPacksValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VirtualClustersPacksValuesParamsWithTimeout(timeout time.Duration) *V1VirtualClustersPacksValuesParams {
	var (
		kubernetesDistroTypeDefault = string("k3s")
	)
	return &V1VirtualClustersPacksValuesParams{
		KubernetesDistroType: &kubernetesDistroTypeDefault,

		timeout: timeout,
	}
}

// NewV1VirtualClustersPacksValuesParamsWithContext creates a new V1VirtualClustersPacksValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VirtualClustersPacksValuesParamsWithContext(ctx context.Context) *V1VirtualClustersPacksValuesParams {
	var (
		kubernetesDistroTypeDefault = string("k3s")
	)
	return &V1VirtualClustersPacksValuesParams{
		KubernetesDistroType: &kubernetesDistroTypeDefault,

		Context: ctx,
	}
}

// NewV1VirtualClustersPacksValuesParamsWithHTTPClient creates a new V1VirtualClustersPacksValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VirtualClustersPacksValuesParamsWithHTTPClient(client *http.Client) *V1VirtualClustersPacksValuesParams {
	var (
		kubernetesDistroTypeDefault = string("k3s")
	)
	return &V1VirtualClustersPacksValuesParams{
		KubernetesDistroType: &kubernetesDistroTypeDefault,
		HTTPClient:           client,
	}
}

/*
V1VirtualClustersPacksValuesParams contains all the parameters to send to the API endpoint
for the v1 virtual clusters packs values operation typically these are written to a http.Request
*/
type V1VirtualClustersPacksValuesParams struct {

	/*KubernetesDistroType
	  Kubernetes distribution type

	*/
	KubernetesDistroType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) WithTimeout(timeout time.Duration) *V1VirtualClustersPacksValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) WithContext(ctx context.Context) *V1VirtualClustersPacksValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) WithHTTPClient(client *http.Client) *V1VirtualClustersPacksValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKubernetesDistroType adds the kubernetesDistroType to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) WithKubernetesDistroType(kubernetesDistroType *string) *V1VirtualClustersPacksValuesParams {
	o.SetKubernetesDistroType(kubernetesDistroType)
	return o
}

// SetKubernetesDistroType adds the kubernetesDistroType to the v1 virtual clusters packs values params
func (o *V1VirtualClustersPacksValuesParams) SetKubernetesDistroType(kubernetesDistroType *string) {
	o.KubernetesDistroType = kubernetesDistroType
}

// WriteToRequest writes these params to a swagger request
func (o *V1VirtualClustersPacksValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KubernetesDistroType != nil {

		// query param kubernetesDistroType
		var qrKubernetesDistroType string
		if o.KubernetesDistroType != nil {
			qrKubernetesDistroType = *o.KubernetesDistroType
		}
		qKubernetesDistroType := qrKubernetesDistroType
		if qKubernetesDistroType != "" {
			if err := r.SetQueryParam("kubernetesDistroType", qKubernetesDistroType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
