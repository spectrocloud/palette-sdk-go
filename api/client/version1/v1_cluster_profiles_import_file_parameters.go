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
	"github.com/go-openapi/swag"
)

// NewV1ClusterProfilesImportFileParams creates a new V1ClusterProfilesImportFileParams object
// with the default values initialized.
func NewV1ClusterProfilesImportFileParams() *V1ClusterProfilesImportFileParams {
	var (
		formatDefault = string("json")
	)
	return &V1ClusterProfilesImportFileParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesImportFileParamsWithTimeout creates a new V1ClusterProfilesImportFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesImportFileParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesImportFileParams {
	var (
		formatDefault = string("json")
	)
	return &V1ClusterProfilesImportFileParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewV1ClusterProfilesImportFileParamsWithContext creates a new V1ClusterProfilesImportFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesImportFileParamsWithContext(ctx context.Context) *V1ClusterProfilesImportFileParams {
	var (
		formatDefault = string("json")
	)
	return &V1ClusterProfilesImportFileParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewV1ClusterProfilesImportFileParamsWithHTTPClient creates a new V1ClusterProfilesImportFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesImportFileParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesImportFileParams {
	var (
		formatDefault = string("json")
	)
	return &V1ClusterProfilesImportFileParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesImportFileParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles import file operation typically these are written to a http.Request
*/
type V1ClusterProfilesImportFileParams struct {

	/*Format
	  Cluster profile import file format ["yaml", "json"]

	*/
	Format *string
	/*ImportFile
	  Cluster profile import file

	*/
	ImportFile runtime.NamedReadCloser
	/*Publish
	  If true then cluster profile will be published post successful import

	*/
	Publish *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesImportFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) WithContext(ctx context.Context) *V1ClusterProfilesImportFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesImportFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) WithFormat(format *string) *V1ClusterProfilesImportFileParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) SetFormat(format *string) {
	o.Format = format
}

// WithImportFile adds the importFile to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) WithImportFile(importFile runtime.NamedReadCloser) *V1ClusterProfilesImportFileParams {
	o.SetImportFile(importFile)
	return o
}

// SetImportFile adds the importFile to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) SetImportFile(importFile runtime.NamedReadCloser) {
	o.ImportFile = importFile
}

// WithPublish adds the publish to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) WithPublish(publish *bool) *V1ClusterProfilesImportFileParams {
	o.SetPublish(publish)
	return o
}

// SetPublish adds the publish to the v1 cluster profiles import file params
func (o *V1ClusterProfilesImportFileParams) SetPublish(publish *bool) {
	o.Publish = publish
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesImportFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	if o.ImportFile != nil {

		if o.ImportFile != nil {

			// form file param importFile
			if err := r.SetFileParam("importFile", o.ImportFile); err != nil {
				return err
			}

		}

	}

	if o.Publish != nil {

		// query param publish
		var qrPublish bool
		if o.Publish != nil {
			qrPublish = *o.Publish
		}
		qPublish := swag.FormatBool(qrPublish)
		if qPublish != "" {
			if err := r.SetQueryParam("publish", qPublish); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
