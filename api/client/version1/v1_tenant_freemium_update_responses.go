// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantFreemiumUpdateReader is a Reader for the V1TenantFreemiumUpdate structure.
type V1TenantFreemiumUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantFreemiumUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantFreemiumUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantFreemiumUpdateNoContent creates a V1TenantFreemiumUpdateNoContent with default headers values
func NewV1TenantFreemiumUpdateNoContent() *V1TenantFreemiumUpdateNoContent {
	return &V1TenantFreemiumUpdateNoContent{}
}

/*
V1TenantFreemiumUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantFreemiumUpdateNoContent struct {
}

func (o *V1TenantFreemiumUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/tenants/{tenantUid}/freemium][%d] v1TenantFreemiumUpdateNoContent ", 204)
}

func (o *V1TenantFreemiumUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
