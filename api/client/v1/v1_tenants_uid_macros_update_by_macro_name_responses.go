// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantsUIDMacrosUpdateByMacroNameReader is a Reader for the V1TenantsUIDMacrosUpdateByMacroName structure.
type V1TenantsUIDMacrosUpdateByMacroNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantsUIDMacrosUpdateByMacroNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantsUIDMacrosUpdateByMacroNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantsUIDMacrosUpdateByMacroNameNoContent creates a V1TenantsUIDMacrosUpdateByMacroNameNoContent with default headers values
func NewV1TenantsUIDMacrosUpdateByMacroNameNoContent() *V1TenantsUIDMacrosUpdateByMacroNameNoContent {
	return &V1TenantsUIDMacrosUpdateByMacroNameNoContent{}
}

/*
V1TenantsUIDMacrosUpdateByMacroNameNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantsUIDMacrosUpdateByMacroNameNoContent struct {
}

func (o *V1TenantsUIDMacrosUpdateByMacroNameNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/tenants/{tenantUid}/macros][%d] v1TenantsUidMacrosUpdateByMacroNameNoContent ", 204)
}

func (o *V1TenantsUIDMacrosUpdateByMacroNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}