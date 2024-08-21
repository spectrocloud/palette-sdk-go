// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantsUIDMacrosUpdateReader is a Reader for the V1TenantsUIDMacrosUpdate structure.
type V1TenantsUIDMacrosUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantsUIDMacrosUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantsUIDMacrosUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantsUIDMacrosUpdateNoContent creates a V1TenantsUIDMacrosUpdateNoContent with default headers values
func NewV1TenantsUIDMacrosUpdateNoContent() *V1TenantsUIDMacrosUpdateNoContent {
	return &V1TenantsUIDMacrosUpdateNoContent{}
}

/*
V1TenantsUIDMacrosUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantsUIDMacrosUpdateNoContent struct {
}

func (o *V1TenantsUIDMacrosUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/tenants/{tenantUid}/macros][%d] v1TenantsUidMacrosUpdateNoContent ", 204)
}

func (o *V1TenantsUIDMacrosUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}