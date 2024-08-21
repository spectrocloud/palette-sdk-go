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

// NewV1SpectroClustersUIDConfigRbacsUIDUpdateParams creates a new V1SpectroClustersUIDConfigRbacsUIDUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDConfigRbacsUIDUpdateParams() *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDUpdateParamsWithTimeout creates a new V1SpectroClustersUIDConfigRbacsUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDConfigRbacsUIDUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDUpdateParamsWithContext creates a new V1SpectroClustersUIDConfigRbacsUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDConfigRbacsUIDUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDConfigRbacsUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDConfigRbacsUIDUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	var ()
	return &V1SpectroClustersUIDConfigRbacsUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDConfigRbacsUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid config rbacs Uid update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDConfigRbacsUIDUpdateParams struct {

	/*Body*/
	Body *models.V1ClusterRbacInputEntity
	/*RbacUID
	  RBAC resource uid

	*/
	RbacUID string
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WithBody(body *models.V1ClusterRbacInputEntity) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) SetBody(body *models.V1ClusterRbacInputEntity) {
	o.Body = body
}

// WithRbacUID adds the rbacUID to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WithRbacUID(rbacUID string) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	o.SetRbacUID(rbacUID)
	return o
}

// SetRbacUID adds the rbacUid to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) SetRbacUID(rbacUID string) {
	o.RbacUID = rbacUID
}

// WithUID adds the uid to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WithUID(uid string) *V1SpectroClustersUIDConfigRbacsUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid config rbacs Uid update params
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param rbacUid
	if err := r.SetPathParam("rbacUid", o.RbacUID); err != nil {
		return err
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