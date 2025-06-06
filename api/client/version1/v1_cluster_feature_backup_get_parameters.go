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
)

// NewV1ClusterFeatureBackupGetParams creates a new V1ClusterFeatureBackupGetParams object
// with the default values initialized.
func NewV1ClusterFeatureBackupGetParams() *V1ClusterFeatureBackupGetParams {
	var ()
	return &V1ClusterFeatureBackupGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureBackupGetParamsWithTimeout creates a new V1ClusterFeatureBackupGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureBackupGetParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureBackupGetParams {
	var ()
	return &V1ClusterFeatureBackupGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureBackupGetParamsWithContext creates a new V1ClusterFeatureBackupGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureBackupGetParamsWithContext(ctx context.Context) *V1ClusterFeatureBackupGetParams {
	var ()
	return &V1ClusterFeatureBackupGetParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureBackupGetParamsWithHTTPClient creates a new V1ClusterFeatureBackupGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureBackupGetParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureBackupGetParams {
	var ()
	return &V1ClusterFeatureBackupGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureBackupGetParams contains all the parameters to send to the API endpoint
for the v1 cluster feature backup get operation typically these are written to a http.Request
*/
type V1ClusterFeatureBackupGetParams struct {

	/*BackupRequestUID*/
	BackupRequestUID *string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureBackupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) WithContext(ctx context.Context) *V1ClusterFeatureBackupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureBackupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupRequestUID adds the backupRequestUID to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) WithBackupRequestUID(backupRequestUID *string) *V1ClusterFeatureBackupGetParams {
	o.SetBackupRequestUID(backupRequestUID)
	return o
}

// SetBackupRequestUID adds the backupRequestUid to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) SetBackupRequestUID(backupRequestUID *string) {
	o.BackupRequestUID = backupRequestUID
}

// WithUID adds the uid to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) WithUID(uid string) *V1ClusterFeatureBackupGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature backup get params
func (o *V1ClusterFeatureBackupGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureBackupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
