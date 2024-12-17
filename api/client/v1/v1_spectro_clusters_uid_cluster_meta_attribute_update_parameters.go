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

// NewV1SpectroClustersUIDClusterMetaAttributeUpdateParams creates a new V1SpectroClustersUIDClusterMetaAttributeUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDClusterMetaAttributeUpdateParams() *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	var ()
	return &V1SpectroClustersUIDClusterMetaAttributeUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithTimeout creates a new V1SpectroClustersUIDClusterMetaAttributeUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	var ()
	return &V1SpectroClustersUIDClusterMetaAttributeUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithContext creates a new V1SpectroClustersUIDClusterMetaAttributeUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	var ()
	return &V1SpectroClustersUIDClusterMetaAttributeUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDClusterMetaAttributeUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	var ()
	return &V1SpectroClustersUIDClusterMetaAttributeUpdateParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersUIDClusterMetaAttributeUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid cluster meta attribute update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDClusterMetaAttributeUpdateParams struct {

	/*Body*/
	Body *models.V1ClusterMetaAttributeEntity
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) WithBody(body *models.V1ClusterMetaAttributeEntity) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) SetBody(body *models.V1ClusterMetaAttributeEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) WithUID(uid string) *V1SpectroClustersUIDClusterMetaAttributeUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid cluster meta attribute update params
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
