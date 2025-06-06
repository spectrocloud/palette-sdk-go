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

// NewV1SpectroClustersUIDOsPatchUpdateParams creates a new V1SpectroClustersUIDOsPatchUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDOsPatchUpdateParams() *V1SpectroClustersUIDOsPatchUpdateParams {
	var ()
	return &V1SpectroClustersUIDOsPatchUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDOsPatchUpdateParamsWithTimeout creates a new V1SpectroClustersUIDOsPatchUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDOsPatchUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDOsPatchUpdateParams {
	var ()
	return &V1SpectroClustersUIDOsPatchUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext creates a new V1SpectroClustersUIDOsPatchUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDOsPatchUpdateParams {
	var ()
	return &V1SpectroClustersUIDOsPatchUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDOsPatchUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDOsPatchUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDOsPatchUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDOsPatchUpdateParams {
	var ()
	return &V1SpectroClustersUIDOsPatchUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDOsPatchUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid os patch update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDOsPatchUpdateParams struct {

	/*Body*/
	Body *models.V1OsPatchEntity
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDOsPatchUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDOsPatchUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDOsPatchUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) WithBody(body *models.V1OsPatchEntity) *V1SpectroClustersUIDOsPatchUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) SetBody(body *models.V1OsPatchEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) WithUID(uid string) *V1SpectroClustersUIDOsPatchUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid os patch update params
func (o *V1SpectroClustersUIDOsPatchUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDOsPatchUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
