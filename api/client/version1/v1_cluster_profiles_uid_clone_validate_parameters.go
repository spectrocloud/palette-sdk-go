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

// NewV1ClusterProfilesUIDCloneValidateParams creates a new V1ClusterProfilesUIDCloneValidateParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDCloneValidateParams() *V1ClusterProfilesUIDCloneValidateParams {
	var ()
	return &V1ClusterProfilesUIDCloneValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDCloneValidateParamsWithTimeout creates a new V1ClusterProfilesUIDCloneValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDCloneValidateParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDCloneValidateParams {
	var ()
	return &V1ClusterProfilesUIDCloneValidateParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDCloneValidateParamsWithContext creates a new V1ClusterProfilesUIDCloneValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDCloneValidateParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDCloneValidateParams {
	var ()
	return &V1ClusterProfilesUIDCloneValidateParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDCloneValidateParamsWithHTTPClient creates a new V1ClusterProfilesUIDCloneValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDCloneValidateParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDCloneValidateParams {
	var ()
	return &V1ClusterProfilesUIDCloneValidateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDCloneValidateParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid clone validate operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDCloneValidateParams struct {

	/*Body*/
	Body *models.V1ClusterProfileCloneMetaInputEntity
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDCloneValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDCloneValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDCloneValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) WithBody(body *models.V1ClusterProfileCloneMetaInputEntity) *V1ClusterProfilesUIDCloneValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) SetBody(body *models.V1ClusterProfileCloneMetaInputEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) WithUID(uid string) *V1ClusterProfilesUIDCloneValidateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid clone validate params
func (o *V1ClusterProfilesUIDCloneValidateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDCloneValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
