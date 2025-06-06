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

// NewV1ClusterProfilesUIDValidatePacksParams creates a new V1ClusterProfilesUIDValidatePacksParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDValidatePacksParams() *V1ClusterProfilesUIDValidatePacksParams {
	var ()
	return &V1ClusterProfilesUIDValidatePacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDValidatePacksParamsWithTimeout creates a new V1ClusterProfilesUIDValidatePacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDValidatePacksParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDValidatePacksParams {
	var ()
	return &V1ClusterProfilesUIDValidatePacksParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDValidatePacksParamsWithContext creates a new V1ClusterProfilesUIDValidatePacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDValidatePacksParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDValidatePacksParams {
	var ()
	return &V1ClusterProfilesUIDValidatePacksParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDValidatePacksParamsWithHTTPClient creates a new V1ClusterProfilesUIDValidatePacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDValidatePacksParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDValidatePacksParams {
	var ()
	return &V1ClusterProfilesUIDValidatePacksParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDValidatePacksParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid validate packs operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDValidatePacksParams struct {

	/*Body*/
	Body *models.V1ClusterProfileTemplateDraft
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDValidatePacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDValidatePacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDValidatePacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) WithBody(body *models.V1ClusterProfileTemplateDraft) *V1ClusterProfilesUIDValidatePacksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) SetBody(body *models.V1ClusterProfileTemplateDraft) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) WithUID(uid string) *V1ClusterProfilesUIDValidatePacksParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid validate packs params
func (o *V1ClusterProfilesUIDValidatePacksParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDValidatePacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
