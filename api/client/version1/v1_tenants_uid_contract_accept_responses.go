// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantsUIDContractAcceptReader is a Reader for the V1TenantsUIDContractAccept structure.
type V1TenantsUIDContractAcceptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantsUIDContractAcceptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantsUIDContractAcceptNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantsUIDContractAcceptNoContent creates a V1TenantsUIDContractAcceptNoContent with default headers values
func NewV1TenantsUIDContractAcceptNoContent() *V1TenantsUIDContractAcceptNoContent {
	return &V1TenantsUIDContractAcceptNoContent{}
}

/*
V1TenantsUIDContractAcceptNoContent handles this case with default header values.

Ok response without content
*/
type V1TenantsUIDContractAcceptNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1TenantsUIDContractAcceptNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/tenants/{tenantUid}/contract/accept][%d] v1TenantsUidContractAcceptNoContent ", 204)
}

func (o *V1TenantsUIDContractAcceptNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
