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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1TenantClustersNodesAutoRemediationSettingUpdateParams creates a new V1TenantClustersNodesAutoRemediationSettingUpdateParams object
// with the default values initialized.
func NewV1TenantClustersNodesAutoRemediationSettingUpdateParams() *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	var ()
	return &V1TenantClustersNodesAutoRemediationSettingUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithTimeout creates a new V1TenantClustersNodesAutoRemediationSettingUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithTimeout(timeout time.Duration) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	var ()
	return &V1TenantClustersNodesAutoRemediationSettingUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithContext creates a new V1TenantClustersNodesAutoRemediationSettingUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithContext(ctx context.Context) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	var ()
	return &V1TenantClustersNodesAutoRemediationSettingUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithHTTPClient creates a new V1TenantClustersNodesAutoRemediationSettingUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithHTTPClient(client *http.Client) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	var ()
	return &V1TenantClustersNodesAutoRemediationSettingUpdateParams{
		HTTPClient: client,
	}
}

/*V1TenantClustersNodesAutoRemediationSettingUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant clusters nodes auto remediation setting update operation typically these are written to a http.Request
*/
type V1TenantClustersNodesAutoRemediationSettingUpdateParams struct {

	/*Body*/
	Body *models.V1NodesAutoRemediationSettings
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) WithTimeout(timeout time.Duration) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) WithContext(ctx context.Context) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) WithHTTPClient(client *http.Client) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) WithBody(body *models.V1NodesAutoRemediationSettings) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) SetBody(body *models.V1NodesAutoRemediationSettings) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) WithTenantUID(tenantUID string) *V1TenantClustersNodesAutoRemediationSettingUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant clusters nodes auto remediation setting update params
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantClustersNodesAutoRemediationSettingUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
