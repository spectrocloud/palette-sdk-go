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

// NewV1ClusterProfilesUpdateParams creates a new V1ClusterProfilesUpdateParams object
// with the default values initialized.
func NewV1ClusterProfilesUpdateParams() *V1ClusterProfilesUpdateParams {
	var ()
	return &V1ClusterProfilesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUpdateParamsWithTimeout creates a new V1ClusterProfilesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUpdateParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUpdateParams {
	var ()
	return &V1ClusterProfilesUpdateParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUpdateParamsWithContext creates a new V1ClusterProfilesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUpdateParamsWithContext(ctx context.Context) *V1ClusterProfilesUpdateParams {
	var ()
	return &V1ClusterProfilesUpdateParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUpdateParamsWithHTTPClient creates a new V1ClusterProfilesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUpdateParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUpdateParams {
	var ()
	return &V1ClusterProfilesUpdateParams{
		HTTPClient: client,
	}
}

/*V1ClusterProfilesUpdateParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles update operation typically these are written to a http.Request
*/
type V1ClusterProfilesUpdateParams struct {

	/*Body*/
	Body *models.V1ClusterProfileUpdateEntity
	/*IncludePackMeta
	  Comma seperated pack meta such as schema, presets

	*/
	IncludePackMeta *string
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) WithContext(ctx context.Context) *V1ClusterProfilesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) WithBody(body *models.V1ClusterProfileUpdateEntity) *V1ClusterProfilesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) SetBody(body *models.V1ClusterProfileUpdateEntity) {
	o.Body = body
}

// WithIncludePackMeta adds the includePackMeta to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) WithIncludePackMeta(includePackMeta *string) *V1ClusterProfilesUpdateParams {
	o.SetIncludePackMeta(includePackMeta)
	return o
}

// SetIncludePackMeta adds the includePackMeta to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) SetIncludePackMeta(includePackMeta *string) {
	o.IncludePackMeta = includePackMeta
}

// WithUID adds the uid to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) WithUID(uid string) *V1ClusterProfilesUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles update params
func (o *V1ClusterProfilesUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.IncludePackMeta != nil {

		// query param includePackMeta
		var qrIncludePackMeta string
		if o.IncludePackMeta != nil {
			qrIncludePackMeta = *o.IncludePackMeta
		}
		qIncludePackMeta := qrIncludePackMeta
		if qIncludePackMeta != "" {
			if err := r.SetQueryParam("includePackMeta", qIncludePackMeta); err != nil {
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
