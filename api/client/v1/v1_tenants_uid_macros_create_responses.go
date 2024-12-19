// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantsUIDMacrosCreateReader is a Reader for the V1TenantsUIDMacrosCreate structure.
type V1TenantsUIDMacrosCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantsUIDMacrosCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantsUIDMacrosCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantsUIDMacrosCreateNoContent creates a V1TenantsUIDMacrosCreateNoContent with default headers values
func NewV1TenantsUIDMacrosCreateNoContent() *V1TenantsUIDMacrosCreateNoContent {
	return &V1TenantsUIDMacrosCreateNoContent{}
}

/*V1TenantsUIDMacrosCreateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantsUIDMacrosCreateNoContent struct {
}

func (o *V1TenantsUIDMacrosCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/tenants/{tenantUid}/macros][%d] v1TenantsUidMacrosCreateNoContent ", 204)
}

func (o *V1TenantsUIDMacrosCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
