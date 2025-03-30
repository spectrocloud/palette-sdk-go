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

// NewV1WorkspaceOpsBackupGetParams creates a new V1WorkspaceOpsBackupGetParams object
// with the default values initialized.
func NewV1WorkspaceOpsBackupGetParams() *V1WorkspaceOpsBackupGetParams {
	var ()
	return &V1WorkspaceOpsBackupGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1WorkspaceOpsBackupGetParamsWithTimeout creates a new V1WorkspaceOpsBackupGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1WorkspaceOpsBackupGetParamsWithTimeout(timeout time.Duration) *V1WorkspaceOpsBackupGetParams {
	var ()
	return &V1WorkspaceOpsBackupGetParams{

		timeout: timeout,
	}
}

// NewV1WorkspaceOpsBackupGetParamsWithContext creates a new V1WorkspaceOpsBackupGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1WorkspaceOpsBackupGetParamsWithContext(ctx context.Context) *V1WorkspaceOpsBackupGetParams {
	var ()
	return &V1WorkspaceOpsBackupGetParams{

		Context: ctx,
	}
}

// NewV1WorkspaceOpsBackupGetParamsWithHTTPClient creates a new V1WorkspaceOpsBackupGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1WorkspaceOpsBackupGetParamsWithHTTPClient(client *http.Client) *V1WorkspaceOpsBackupGetParams {
	var ()
	return &V1WorkspaceOpsBackupGetParams{
		HTTPClient: client,
	}
}

/*
V1WorkspaceOpsBackupGetParams contains all the parameters to send to the API endpoint
for the v1 workspace ops backup get operation typically these are written to a http.Request
*/
type V1WorkspaceOpsBackupGetParams struct {

	/*BackupRequestUID*/
	BackupRequestUID *string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) WithTimeout(timeout time.Duration) *V1WorkspaceOpsBackupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) WithContext(ctx context.Context) *V1WorkspaceOpsBackupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) WithHTTPClient(client *http.Client) *V1WorkspaceOpsBackupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupRequestUID adds the backupRequestUID to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) WithBackupRequestUID(backupRequestUID *string) *V1WorkspaceOpsBackupGetParams {
	o.SetBackupRequestUID(backupRequestUID)
	return o
}

// SetBackupRequestUID adds the backupRequestUid to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) SetBackupRequestUID(backupRequestUID *string) {
	o.BackupRequestUID = backupRequestUID
}

// WithUID adds the uid to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) WithUID(uid string) *V1WorkspaceOpsBackupGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 workspace ops backup get params
func (o *V1WorkspaceOpsBackupGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1WorkspaceOpsBackupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupRequestUID != nil {

		// query param backupRequestUid
		var qrBackupRequestUID string
		if o.BackupRequestUID != nil {
			qrBackupRequestUID = *o.BackupRequestUID
		}
		qBackupRequestUID := qrBackupRequestUID
		if qBackupRequestUID != "" {
			if err := r.SetQueryParam("backupRequestUid", qBackupRequestUID); err != nil {
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
