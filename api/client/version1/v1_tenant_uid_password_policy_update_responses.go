// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantUIDPasswordPolicyUpdateReader is a Reader for the V1TenantUIDPasswordPolicyUpdate structure.
type V1TenantUIDPasswordPolicyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantUIDPasswordPolicyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantUIDPasswordPolicyUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantUIDPasswordPolicyUpdateNoContent creates a V1TenantUIDPasswordPolicyUpdateNoContent with default headers values
func NewV1TenantUIDPasswordPolicyUpdateNoContent() *V1TenantUIDPasswordPolicyUpdateNoContent {
	return &V1TenantUIDPasswordPolicyUpdateNoContent{}
}

/*
V1TenantUIDPasswordPolicyUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantUIDPasswordPolicyUpdateNoContent struct {
}

func (o *V1TenantUIDPasswordPolicyUpdateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/tenants/{tenantUid}/password/policy][%d] v1TenantUidPasswordPolicyUpdateNoContent ", 204)
}

func (o *V1TenantUIDPasswordPolicyUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
