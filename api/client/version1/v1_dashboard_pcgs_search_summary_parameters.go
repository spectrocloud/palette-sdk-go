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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1DashboardPcgsSearchSummaryParams creates a new V1DashboardPcgsSearchSummaryParams object
// with the default values initialized.
func NewV1DashboardPcgsSearchSummaryParams() *V1DashboardPcgsSearchSummaryParams {
	var ()
	return &V1DashboardPcgsSearchSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardPcgsSearchSummaryParamsWithTimeout creates a new V1DashboardPcgsSearchSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardPcgsSearchSummaryParamsWithTimeout(timeout time.Duration) *V1DashboardPcgsSearchSummaryParams {
	var ()
	return &V1DashboardPcgsSearchSummaryParams{

		timeout: timeout,
	}
}

// NewV1DashboardPcgsSearchSummaryParamsWithContext creates a new V1DashboardPcgsSearchSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardPcgsSearchSummaryParamsWithContext(ctx context.Context) *V1DashboardPcgsSearchSummaryParams {
	var ()
	return &V1DashboardPcgsSearchSummaryParams{

		Context: ctx,
	}
}

// NewV1DashboardPcgsSearchSummaryParamsWithHTTPClient creates a new V1DashboardPcgsSearchSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardPcgsSearchSummaryParamsWithHTTPClient(client *http.Client) *V1DashboardPcgsSearchSummaryParams {
	var ()
	return &V1DashboardPcgsSearchSummaryParams{
		HTTPClient: client,
	}
}

/*
V1DashboardPcgsSearchSummaryParams contains all the parameters to send to the API endpoint
for the v1 dashboard pcgs search summary operation typically these are written to a http.Request
*/
type V1DashboardPcgsSearchSummaryParams struct {

	/*Body*/
	Body *models.V1SearchFilterSummarySpec
	/*Continue
	  continue token to paginate the subsequent data items

	*/
	Continue *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. Maximum value of the limit is 50.
	If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results.

	*/
	Limit *int64
	/*Offset
	  offset is the next index number from which the response will start. The response offset value can be used along with continue token for the pagination.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithTimeout(timeout time.Duration) *V1DashboardPcgsSearchSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithContext(ctx context.Context) *V1DashboardPcgsSearchSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithHTTPClient(client *http.Client) *V1DashboardPcgsSearchSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithBody(body *models.V1SearchFilterSummarySpec) *V1DashboardPcgsSearchSummaryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetBody(body *models.V1SearchFilterSummarySpec) {
	o.Body = body
}

// WithContinue adds the continueVar to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithContinue(continueVar *string) *V1DashboardPcgsSearchSummaryParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithLimit adds the limit to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithLimit(limit *int64) *V1DashboardPcgsSearchSummaryParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) WithOffset(offset *int64) *V1DashboardPcgsSearchSummaryParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the v1 dashboard pcgs search summary params
func (o *V1DashboardPcgsSearchSummaryParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardPcgsSearchSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Continue != nil {

		// query param continue
		var qrContinue string
		if o.Continue != nil {
			qrContinue = *o.Continue
		}
		qContinue := qrContinue
		if qContinue != "" {
			if err := r.SetQueryParam("continue", qContinue); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
