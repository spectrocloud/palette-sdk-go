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

// NewV1OpenstackAccountsUIDNetworksParams creates a new V1OpenstackAccountsUIDNetworksParams object
// with the default values initialized.
func NewV1OpenstackAccountsUIDNetworksParams() *V1OpenstackAccountsUIDNetworksParams {
	var ()
	return &V1OpenstackAccountsUIDNetworksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OpenstackAccountsUIDNetworksParamsWithTimeout creates a new V1OpenstackAccountsUIDNetworksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OpenstackAccountsUIDNetworksParamsWithTimeout(timeout time.Duration) *V1OpenstackAccountsUIDNetworksParams {
	var ()
	return &V1OpenstackAccountsUIDNetworksParams{

		timeout: timeout,
	}
}

// NewV1OpenstackAccountsUIDNetworksParamsWithContext creates a new V1OpenstackAccountsUIDNetworksParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OpenstackAccountsUIDNetworksParamsWithContext(ctx context.Context) *V1OpenstackAccountsUIDNetworksParams {
	var ()
	return &V1OpenstackAccountsUIDNetworksParams{

		Context: ctx,
	}
}

// NewV1OpenstackAccountsUIDNetworksParamsWithHTTPClient creates a new V1OpenstackAccountsUIDNetworksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OpenstackAccountsUIDNetworksParamsWithHTTPClient(client *http.Client) *V1OpenstackAccountsUIDNetworksParams {
	var ()
	return &V1OpenstackAccountsUIDNetworksParams{
		HTTPClient: client,
	}
}

/*
V1OpenstackAccountsUIDNetworksParams contains all the parameters to send to the API endpoint
for the v1 openstack accounts Uid networks operation typically these are written to a http.Request
*/
type V1OpenstackAccountsUIDNetworksParams struct {

	/*Domain*/
	Domain *string
	/*Project*/
	Project *string
	/*Region*/
	Region *string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithTimeout(timeout time.Duration) *V1OpenstackAccountsUIDNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithContext(ctx context.Context) *V1OpenstackAccountsUIDNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithHTTPClient(client *http.Client) *V1OpenstackAccountsUIDNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithDomain(domain *string) *V1OpenstackAccountsUIDNetworksParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithProject adds the project to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithProject(project *string) *V1OpenstackAccountsUIDNetworksParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetProject(project *string) {
	o.Project = project
}

// WithRegion adds the region to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithRegion(region *string) *V1OpenstackAccountsUIDNetworksParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetRegion(region *string) {
	o.Region = region
}

// WithUID adds the uid to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) WithUID(uid string) *V1OpenstackAccountsUIDNetworksParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 openstack accounts Uid networks params
func (o *V1OpenstackAccountsUIDNetworksParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OpenstackAccountsUIDNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Domain != nil {

		// query param domain
		var qrDomain string
		if o.Domain != nil {
			qrDomain = *o.Domain
		}
		qDomain := qrDomain
		if qDomain != "" {
			if err := r.SetQueryParam("domain", qDomain); err != nil {
				return err
			}
		}

	}

	if o.Project != nil {

		// query param project
		var qrProject string
		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {
			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
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
