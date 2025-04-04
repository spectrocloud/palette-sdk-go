// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantUIDOidcConfigUpdateReader is a Reader for the V1TenantUIDOidcConfigUpdate structure.
type V1TenantUIDOidcConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantUIDOidcConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantUIDOidcConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantUIDOidcConfigUpdateNoContent creates a V1TenantUIDOidcConfigUpdateNoContent with default headers values
func NewV1TenantUIDOidcConfigUpdateNoContent() *V1TenantUIDOidcConfigUpdateNoContent {
	return &V1TenantUIDOidcConfigUpdateNoContent{}
}

/*
V1TenantUIDOidcConfigUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantUIDOidcConfigUpdateNoContent struct {
}

func (o *V1TenantUIDOidcConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/tenants/{tenantUid}/oidc/config][%d] v1TenantUidOidcConfigUpdateNoContent ", 204)
}

func (o *V1TenantUIDOidcConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
