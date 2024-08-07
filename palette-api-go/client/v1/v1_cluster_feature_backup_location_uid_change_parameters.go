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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1ClusterFeatureBackupLocationUIDChangeParams creates a new V1ClusterFeatureBackupLocationUIDChangeParams object
// with the default values initialized.
func NewV1ClusterFeatureBackupLocationUIDChangeParams() *V1ClusterFeatureBackupLocationUIDChangeParams {
	var ()
	return &V1ClusterFeatureBackupLocationUIDChangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureBackupLocationUIDChangeParamsWithTimeout creates a new V1ClusterFeatureBackupLocationUIDChangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureBackupLocationUIDChangeParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureBackupLocationUIDChangeParams {
	var ()
	return &V1ClusterFeatureBackupLocationUIDChangeParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureBackupLocationUIDChangeParamsWithContext creates a new V1ClusterFeatureBackupLocationUIDChangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureBackupLocationUIDChangeParamsWithContext(ctx context.Context) *V1ClusterFeatureBackupLocationUIDChangeParams {
	var ()
	return &V1ClusterFeatureBackupLocationUIDChangeParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureBackupLocationUIDChangeParamsWithHTTPClient creates a new V1ClusterFeatureBackupLocationUIDChangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureBackupLocationUIDChangeParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureBackupLocationUIDChangeParams {
	var ()
	return &V1ClusterFeatureBackupLocationUIDChangeParams{
		HTTPClient: client,
	}
}

/*V1ClusterFeatureBackupLocationUIDChangeParams contains all the parameters to send to the API endpoint
for the v1 cluster feature backup location Uid change operation typically these are written to a http.Request
*/
type V1ClusterFeatureBackupLocationUIDChangeParams struct {

	/*Body*/
	Body *models.V1ClusterBackupLocationType
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureBackupLocationUIDChangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) WithContext(ctx context.Context) *V1ClusterFeatureBackupLocationUIDChangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureBackupLocationUIDChangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) WithBody(body *models.V1ClusterBackupLocationType) *V1ClusterFeatureBackupLocationUIDChangeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) SetBody(body *models.V1ClusterBackupLocationType) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) WithUID(uid string) *V1ClusterFeatureBackupLocationUIDChangeParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature backup location Uid change params
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureBackupLocationUIDChangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
