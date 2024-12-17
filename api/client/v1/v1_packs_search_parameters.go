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
	"github.com/go-openapi/swag"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1PacksSearchParams creates a new V1PacksSearchParams object
// with the default values initialized.
func NewV1PacksSearchParams() *V1PacksSearchParams {
	var (
		limitDefault = int64(50)
	)
	return &V1PacksSearchParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1PacksSearchParamsWithTimeout creates a new V1PacksSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1PacksSearchParamsWithTimeout(timeout time.Duration) *V1PacksSearchParams {
	var (
		limitDefault = int64(50)
	)
	return &V1PacksSearchParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1PacksSearchParamsWithContext creates a new V1PacksSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1PacksSearchParamsWithContext(ctx context.Context) *V1PacksSearchParams {
	var (
		limitDefault = int64(50)
	)
	return &V1PacksSearchParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1PacksSearchParamsWithHTTPClient creates a new V1PacksSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1PacksSearchParamsWithHTTPClient(client *http.Client) *V1PacksSearchParams {
	var (
		limitDefault = int64(50)
	)
	return &V1PacksSearchParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*V1PacksSearchParams contains all the parameters to send to the API endpoint
for the v1 packs search operation typically these are written to a http.Request
*/
type V1PacksSearchParams struct {

	/*Body*/
	Body *models.V1PacksFilterSpec
	/*Continue
	  continue token to paginate the subsequent data items

	*/
	Continue *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. Default and maximum value of the limit is 50.
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

// WithTimeout adds the timeout to the v1 packs search params
func (o *V1PacksSearchParams) WithTimeout(timeout time.Duration) *V1PacksSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 packs search params
func (o *V1PacksSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 packs search params
func (o *V1PacksSearchParams) WithContext(ctx context.Context) *V1PacksSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 packs search params
func (o *V1PacksSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 packs search params
func (o *V1PacksSearchParams) WithHTTPClient(client *http.Client) *V1PacksSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 packs search params
func (o *V1PacksSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 packs search params
func (o *V1PacksSearchParams) WithBody(body *models.V1PacksFilterSpec) *V1PacksSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 packs search params
func (o *V1PacksSearchParams) SetBody(body *models.V1PacksFilterSpec) {
	o.Body = body
}

// WithContinue adds the continueVar to the v1 packs search params
func (o *V1PacksSearchParams) WithContinue(continueVar *string) *V1PacksSearchParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 packs search params
func (o *V1PacksSearchParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithLimit adds the limit to the v1 packs search params
func (o *V1PacksSearchParams) WithLimit(limit *int64) *V1PacksSearchParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 packs search params
func (o *V1PacksSearchParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the v1 packs search params
func (o *V1PacksSearchParams) WithOffset(offset *int64) *V1PacksSearchParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the v1 packs search params
func (o *V1PacksSearchParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *V1PacksSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
