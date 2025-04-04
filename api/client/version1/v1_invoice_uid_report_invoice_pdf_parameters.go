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

// NewV1InvoiceUIDReportInvoicePdfParams creates a new V1InvoiceUIDReportInvoicePdfParams object
// with the default values initialized.
func NewV1InvoiceUIDReportInvoicePdfParams() *V1InvoiceUIDReportInvoicePdfParams {
	var ()
	return &V1InvoiceUIDReportInvoicePdfParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1InvoiceUIDReportInvoicePdfParamsWithTimeout creates a new V1InvoiceUIDReportInvoicePdfParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1InvoiceUIDReportInvoicePdfParamsWithTimeout(timeout time.Duration) *V1InvoiceUIDReportInvoicePdfParams {
	var ()
	return &V1InvoiceUIDReportInvoicePdfParams{

		timeout: timeout,
	}
}

// NewV1InvoiceUIDReportInvoicePdfParamsWithContext creates a new V1InvoiceUIDReportInvoicePdfParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1InvoiceUIDReportInvoicePdfParamsWithContext(ctx context.Context) *V1InvoiceUIDReportInvoicePdfParams {
	var ()
	return &V1InvoiceUIDReportInvoicePdfParams{

		Context: ctx,
	}
}

// NewV1InvoiceUIDReportInvoicePdfParamsWithHTTPClient creates a new V1InvoiceUIDReportInvoicePdfParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1InvoiceUIDReportInvoicePdfParamsWithHTTPClient(client *http.Client) *V1InvoiceUIDReportInvoicePdfParams {
	var ()
	return &V1InvoiceUIDReportInvoicePdfParams{
		HTTPClient: client,
	}
}

/*
V1InvoiceUIDReportInvoicePdfParams contains all the parameters to send to the API endpoint
for the v1 invoice Uid report invoice pdf operation typically these are written to a http.Request
*/
type V1InvoiceUIDReportInvoicePdfParams struct {

	/*InvoiceUID
	  Specify the invoice uid

	*/
	InvoiceUID string
	/*TenantUID
	  Specify the tenant uid

	*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) WithTimeout(timeout time.Duration) *V1InvoiceUIDReportInvoicePdfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) WithContext(ctx context.Context) *V1InvoiceUIDReportInvoicePdfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) WithHTTPClient(client *http.Client) *V1InvoiceUIDReportInvoicePdfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvoiceUID adds the invoiceUID to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) WithInvoiceUID(invoiceUID string) *V1InvoiceUIDReportInvoicePdfParams {
	o.SetInvoiceUID(invoiceUID)
	return o
}

// SetInvoiceUID adds the invoiceUid to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) SetInvoiceUID(invoiceUID string) {
	o.InvoiceUID = invoiceUID
}

// WithTenantUID adds the tenantUID to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) WithTenantUID(tenantUID string) *V1InvoiceUIDReportInvoicePdfParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 invoice Uid report invoice pdf params
func (o *V1InvoiceUIDReportInvoicePdfParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1InvoiceUIDReportInvoicePdfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invoiceUid
	if err := r.SetPathParam("invoiceUid", o.InvoiceUID); err != nil {
		return err
	}

	// path param tenantUid
	if err := r.SetPathParam("tenantUid", o.TenantUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
