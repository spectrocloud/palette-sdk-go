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

// NewV1ClusterVMSnapshotsListParams creates a new V1ClusterVMSnapshotsListParams object
// with the default values initialized.
func NewV1ClusterVMSnapshotsListParams() *V1ClusterVMSnapshotsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1ClusterVMSnapshotsListParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterVMSnapshotsListParamsWithTimeout creates a new V1ClusterVMSnapshotsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterVMSnapshotsListParamsWithTimeout(timeout time.Duration) *V1ClusterVMSnapshotsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1ClusterVMSnapshotsListParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewV1ClusterVMSnapshotsListParamsWithContext creates a new V1ClusterVMSnapshotsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterVMSnapshotsListParamsWithContext(ctx context.Context) *V1ClusterVMSnapshotsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1ClusterVMSnapshotsListParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewV1ClusterVMSnapshotsListParamsWithHTTPClient creates a new V1ClusterVMSnapshotsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterVMSnapshotsListParamsWithHTTPClient(client *http.Client) *V1ClusterVMSnapshotsListParams {
	var (
		limitDefault = int64(50)
	)
	return &V1ClusterVMSnapshotsListParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*
V1ClusterVMSnapshotsListParams contains all the parameters to send to the API endpoint
for the v1 cluster VM snapshots list operation typically these are written to a http.Request
*/
type V1ClusterVMSnapshotsListParams struct {

	/*Continue
	  continue token to paginate the subsequent data items

	*/
	Continue *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. Default and maximum value of the limit is 50.
	If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results.

	*/
	Limit *int64
	/*Namespace
	  Namespace names, comma separated value (ex: dev,test). If namespace is empty it returns the specific resource under all namespace

	*/
	Namespace []string
	/*UID
	  Cluster uid

	*/
	UID string
	/*VMName
	  vmName is comma separated value (ex: name1,name2).

	*/
	VMName []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithTimeout(timeout time.Duration) *V1ClusterVMSnapshotsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithContext(ctx context.Context) *V1ClusterVMSnapshotsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithHTTPClient(client *http.Client) *V1ClusterVMSnapshotsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinue adds the continueVar to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithContinue(continueVar *string) *V1ClusterVMSnapshotsListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithLimit adds the limit to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithLimit(limit *int64) *V1ClusterVMSnapshotsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithNamespace(namespace []string) *V1ClusterVMSnapshotsListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetNamespace(namespace []string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithUID(uid string) *V1ClusterVMSnapshotsListParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) WithVMName(vMName []string) *V1ClusterVMSnapshotsListParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 cluster VM snapshots list params
func (o *V1ClusterVMSnapshotsListParams) SetVMName(vMName []string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterVMSnapshotsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	valuesNamespace := o.Namespace

	joinedNamespace := swag.JoinByFormat(valuesNamespace, "csv")
	// query array param namespace
	if err := r.SetQueryParam("namespace", joinedNamespace...); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	valuesVMName := o.VMName

	joinedVMName := swag.JoinByFormat(valuesVMName, "csv")
	// query array param vmName
	if err := r.SetQueryParam("vmName", joinedVMName...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
