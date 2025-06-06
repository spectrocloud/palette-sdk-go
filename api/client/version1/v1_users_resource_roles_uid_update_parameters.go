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

// NewV1UsersResourceRolesUIDUpdateParams creates a new V1UsersResourceRolesUIDUpdateParams object
// with the default values initialized.
func NewV1UsersResourceRolesUIDUpdateParams() *V1UsersResourceRolesUIDUpdateParams {
	var ()
	return &V1UsersResourceRolesUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersResourceRolesUIDUpdateParamsWithTimeout creates a new V1UsersResourceRolesUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersResourceRolesUIDUpdateParamsWithTimeout(timeout time.Duration) *V1UsersResourceRolesUIDUpdateParams {
	var ()
	return &V1UsersResourceRolesUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1UsersResourceRolesUIDUpdateParamsWithContext creates a new V1UsersResourceRolesUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersResourceRolesUIDUpdateParamsWithContext(ctx context.Context) *V1UsersResourceRolesUIDUpdateParams {
	var ()
	return &V1UsersResourceRolesUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1UsersResourceRolesUIDUpdateParamsWithHTTPClient creates a new V1UsersResourceRolesUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersResourceRolesUIDUpdateParamsWithHTTPClient(client *http.Client) *V1UsersResourceRolesUIDUpdateParams {
	var ()
	return &V1UsersResourceRolesUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1UsersResourceRolesUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 users resource roles Uid update operation typically these are written to a http.Request
*/
type V1UsersResourceRolesUIDUpdateParams struct {

	/*Body*/
	Body *models.V1ResourceRolesUpdateEntity
	/*ResourceRoleUID*/
	ResourceRoleUID string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) WithTimeout(timeout time.Duration) *V1UsersResourceRolesUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) WithContext(ctx context.Context) *V1UsersResourceRolesUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) WithHTTPClient(client *http.Client) *V1UsersResourceRolesUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) WithBody(body *models.V1ResourceRolesUpdateEntity) *V1UsersResourceRolesUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) SetBody(body *models.V1ResourceRolesUpdateEntity) {
	o.Body = body
}

// WithResourceRoleUID adds the resourceRoleUID to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) WithResourceRoleUID(resourceRoleUID string) *V1UsersResourceRolesUIDUpdateParams {
	o.SetResourceRoleUID(resourceRoleUID)
	return o
}

// SetResourceRoleUID adds the resourceRoleUid to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) SetResourceRoleUID(resourceRoleUID string) {
	o.ResourceRoleUID = resourceRoleUID
}

// WithUID adds the uid to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) WithUID(uid string) *V1UsersResourceRolesUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users resource roles Uid update params
func (o *V1UsersResourceRolesUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersResourceRolesUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resourceRoleUid
	if err := r.SetPathParam("resourceRoleUid", o.ResourceRoleUID); err != nil {
		return err
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
