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

// NewV1AuditsUIDMsgUpdateParams creates a new V1AuditsUIDMsgUpdateParams object
// with the default values initialized.
func NewV1AuditsUIDMsgUpdateParams() *V1AuditsUIDMsgUpdateParams {
	var ()
	return &V1AuditsUIDMsgUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AuditsUIDMsgUpdateParamsWithTimeout creates a new V1AuditsUIDMsgUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AuditsUIDMsgUpdateParamsWithTimeout(timeout time.Duration) *V1AuditsUIDMsgUpdateParams {
	var ()
	return &V1AuditsUIDMsgUpdateParams{

		timeout: timeout,
	}
}

// NewV1AuditsUIDMsgUpdateParamsWithContext creates a new V1AuditsUIDMsgUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AuditsUIDMsgUpdateParamsWithContext(ctx context.Context) *V1AuditsUIDMsgUpdateParams {
	var ()
	return &V1AuditsUIDMsgUpdateParams{

		Context: ctx,
	}
}

// NewV1AuditsUIDMsgUpdateParamsWithHTTPClient creates a new V1AuditsUIDMsgUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AuditsUIDMsgUpdateParamsWithHTTPClient(client *http.Client) *V1AuditsUIDMsgUpdateParams {
	var ()
	return &V1AuditsUIDMsgUpdateParams{
		HTTPClient: client,
	}
}

/*
V1AuditsUIDMsgUpdateParams contains all the parameters to send to the API endpoint
for the v1 audits Uid msg update operation typically these are written to a http.Request
*/
type V1AuditsUIDMsgUpdateParams struct {

	/*Body*/
	Body *models.V1AuditMsgUpdate
	/*UID
	  Specify the audit uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) WithTimeout(timeout time.Duration) *V1AuditsUIDMsgUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) WithContext(ctx context.Context) *V1AuditsUIDMsgUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) WithHTTPClient(client *http.Client) *V1AuditsUIDMsgUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) WithBody(body *models.V1AuditMsgUpdate) *V1AuditsUIDMsgUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) SetBody(body *models.V1AuditMsgUpdate) {
	o.Body = body
}

// WithUID adds the uid to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) WithUID(uid string) *V1AuditsUIDMsgUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 audits Uid msg update params
func (o *V1AuditsUIDMsgUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AuditsUIDMsgUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
