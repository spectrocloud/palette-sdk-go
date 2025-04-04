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

// NewV1WorkspaceOpsBackupUpdateParams creates a new V1WorkspaceOpsBackupUpdateParams object
// with the default values initialized.
func NewV1WorkspaceOpsBackupUpdateParams() *V1WorkspaceOpsBackupUpdateParams {
	var ()
	return &V1WorkspaceOpsBackupUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1WorkspaceOpsBackupUpdateParamsWithTimeout creates a new V1WorkspaceOpsBackupUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1WorkspaceOpsBackupUpdateParamsWithTimeout(timeout time.Duration) *V1WorkspaceOpsBackupUpdateParams {
	var ()
	return &V1WorkspaceOpsBackupUpdateParams{

		timeout: timeout,
	}
}

// NewV1WorkspaceOpsBackupUpdateParamsWithContext creates a new V1WorkspaceOpsBackupUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1WorkspaceOpsBackupUpdateParamsWithContext(ctx context.Context) *V1WorkspaceOpsBackupUpdateParams {
	var ()
	return &V1WorkspaceOpsBackupUpdateParams{

		Context: ctx,
	}
}

// NewV1WorkspaceOpsBackupUpdateParamsWithHTTPClient creates a new V1WorkspaceOpsBackupUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1WorkspaceOpsBackupUpdateParamsWithHTTPClient(client *http.Client) *V1WorkspaceOpsBackupUpdateParams {
	var ()
	return &V1WorkspaceOpsBackupUpdateParams{
		HTTPClient: client,
	}
}

/*
V1WorkspaceOpsBackupUpdateParams contains all the parameters to send to the API endpoint
for the v1 workspace ops backup update operation typically these are written to a http.Request
*/
type V1WorkspaceOpsBackupUpdateParams struct {

	/*Body*/
	Body *models.V1WorkspaceBackupConfigEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) WithTimeout(timeout time.Duration) *V1WorkspaceOpsBackupUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) WithContext(ctx context.Context) *V1WorkspaceOpsBackupUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) WithHTTPClient(client *http.Client) *V1WorkspaceOpsBackupUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) WithBody(body *models.V1WorkspaceBackupConfigEntity) *V1WorkspaceOpsBackupUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) SetBody(body *models.V1WorkspaceBackupConfigEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) WithUID(uid string) *V1WorkspaceOpsBackupUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 workspace ops backup update params
func (o *V1WorkspaceOpsBackupUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1WorkspaceOpsBackupUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
