// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AuthUserOrgForgotReader is a Reader for the V1AuthUserOrgForgot structure.
type V1AuthUserOrgForgotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AuthUserOrgForgotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AuthUserOrgForgotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AuthUserOrgForgotNoContent creates a V1AuthUserOrgForgotNoContent with default headers values
func NewV1AuthUserOrgForgotNoContent() *V1AuthUserOrgForgotNoContent {
	return &V1AuthUserOrgForgotNoContent{}
}

/*
V1AuthUserOrgForgotNoContent handles this case with default header values.

Ok response without content
*/
type V1AuthUserOrgForgotNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1AuthUserOrgForgotNoContent) Error() string {
	return fmt.Sprintf("[GET /v1/auth/user/org/forgot][%d] v1AuthUserOrgForgotNoContent ", 204)
}

func (o *V1AuthUserOrgForgotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
