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

// NewV1MaasTagsGetParams creates a new V1MaasTagsGetParams object
// with the default values initialized.
func NewV1MaasTagsGetParams() *V1MaasTagsGetParams {
	var ()
	return &V1MaasTagsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MaasTagsGetParamsWithTimeout creates a new V1MaasTagsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MaasTagsGetParamsWithTimeout(timeout time.Duration) *V1MaasTagsGetParams {
	var ()
	return &V1MaasTagsGetParams{

		timeout: timeout,
	}
}

// NewV1MaasTagsGetParamsWithContext creates a new V1MaasTagsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MaasTagsGetParamsWithContext(ctx context.Context) *V1MaasTagsGetParams {
	var ()
	return &V1MaasTagsGetParams{

		Context: ctx,
	}
}

// NewV1MaasTagsGetParamsWithHTTPClient creates a new V1MaasTagsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MaasTagsGetParamsWithHTTPClient(client *http.Client) *V1MaasTagsGetParams {
	var ()
	return &V1MaasTagsGetParams{
		HTTPClient: client,
	}
}

/*V1MaasTagsGetParams contains all the parameters to send to the API endpoint
for the v1 maas tags get operation typically these are written to a http.Request
*/
type V1MaasTagsGetParams struct {

	/*CloudAccountUID
	  Uid for the specific Maas cloud account

	*/
	CloudAccountUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 maas tags get params
func (o *V1MaasTagsGetParams) WithTimeout(timeout time.Duration) *V1MaasTagsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 maas tags get params
func (o *V1MaasTagsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 maas tags get params
func (o *V1MaasTagsGetParams) WithContext(ctx context.Context) *V1MaasTagsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 maas tags get params
func (o *V1MaasTagsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 maas tags get params
func (o *V1MaasTagsGetParams) WithHTTPClient(client *http.Client) *V1MaasTagsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 maas tags get params
func (o *V1MaasTagsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 maas tags get params
func (o *V1MaasTagsGetParams) WithCloudAccountUID(cloudAccountUID *string) *V1MaasTagsGetParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 maas tags get params
func (o *V1MaasTagsGetParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1MaasTagsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudAccountUID != nil {

		// query param cloudAccountUid
		var qrCloudAccountUID string
		if o.CloudAccountUID != nil {
			qrCloudAccountUID = *o.CloudAccountUID
		}
		qCloudAccountUID := qrCloudAccountUID
		if qCloudAccountUID != "" {
			if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
