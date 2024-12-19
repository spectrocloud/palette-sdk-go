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

// NewV1SpectroClustersUIDValidateRepaveParams creates a new V1SpectroClustersUIDValidateRepaveParams object
// with the default values initialized.
func NewV1SpectroClustersUIDValidateRepaveParams() *V1SpectroClustersUIDValidateRepaveParams {
	var ()
	return &V1SpectroClustersUIDValidateRepaveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDValidateRepaveParamsWithTimeout creates a new V1SpectroClustersUIDValidateRepaveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDValidateRepaveParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDValidateRepaveParams {
	var ()
	return &V1SpectroClustersUIDValidateRepaveParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDValidateRepaveParamsWithContext creates a new V1SpectroClustersUIDValidateRepaveParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDValidateRepaveParamsWithContext(ctx context.Context) *V1SpectroClustersUIDValidateRepaveParams {
	var ()
	return &V1SpectroClustersUIDValidateRepaveParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDValidateRepaveParamsWithHTTPClient creates a new V1SpectroClustersUIDValidateRepaveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDValidateRepaveParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDValidateRepaveParams {
	var ()
	return &V1SpectroClustersUIDValidateRepaveParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersUIDValidateRepaveParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid validate repave operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDValidateRepaveParams struct {

	/*Body*/
	Body *models.V1SpectroClusterPacksEntity
	/*UID
	  cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDValidateRepaveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) WithContext(ctx context.Context) *V1SpectroClustersUIDValidateRepaveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDValidateRepaveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) WithBody(body *models.V1SpectroClusterPacksEntity) *V1SpectroClustersUIDValidateRepaveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) SetBody(body *models.V1SpectroClusterPacksEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) WithUID(uid string) *V1SpectroClustersUIDValidateRepaveParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid validate repave params
func (o *V1SpectroClustersUIDValidateRepaveParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDValidateRepaveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
