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

// NewV1OpenstackAccountsUIDFlavorsParams creates a new V1OpenstackAccountsUIDFlavorsParams object
// with the default values initialized.
func NewV1OpenstackAccountsUIDFlavorsParams() *V1OpenstackAccountsUIDFlavorsParams {
	var ()
	return &V1OpenstackAccountsUIDFlavorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OpenstackAccountsUIDFlavorsParamsWithTimeout creates a new V1OpenstackAccountsUIDFlavorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OpenstackAccountsUIDFlavorsParamsWithTimeout(timeout time.Duration) *V1OpenstackAccountsUIDFlavorsParams {
	var ()
	return &V1OpenstackAccountsUIDFlavorsParams{

		timeout: timeout,
	}
}

// NewV1OpenstackAccountsUIDFlavorsParamsWithContext creates a new V1OpenstackAccountsUIDFlavorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OpenstackAccountsUIDFlavorsParamsWithContext(ctx context.Context) *V1OpenstackAccountsUIDFlavorsParams {
	var ()
	return &V1OpenstackAccountsUIDFlavorsParams{

		Context: ctx,
	}
}

// NewV1OpenstackAccountsUIDFlavorsParamsWithHTTPClient creates a new V1OpenstackAccountsUIDFlavorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OpenstackAccountsUIDFlavorsParamsWithHTTPClient(client *http.Client) *V1OpenstackAccountsUIDFlavorsParams {
	var ()
	return &V1OpenstackAccountsUIDFlavorsParams{
		HTTPClient: client,
	}
}

/*
V1OpenstackAccountsUIDFlavorsParams contains all the parameters to send to the API endpoint
for the v1 openstack accounts Uid flavors operation typically these are written to a http.Request
*/
type V1OpenstackAccountsUIDFlavorsParams struct {

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

// WithTimeout adds the timeout to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithTimeout(timeout time.Duration) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithContext(ctx context.Context) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithHTTPClient(client *http.Client) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithDomain(domain *string) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithProject adds the project to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithProject(project *string) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetProject(project *string) {
	o.Project = project
}

// WithRegion adds the region to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithRegion(region *string) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetRegion(region *string) {
	o.Region = region
}

// WithUID adds the uid to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) WithUID(uid string) *V1OpenstackAccountsUIDFlavorsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 openstack accounts Uid flavors params
func (o *V1OpenstackAccountsUIDFlavorsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OpenstackAccountsUIDFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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