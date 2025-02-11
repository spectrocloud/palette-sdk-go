// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1OidcLogoutReader is a Reader for the V1OidcLogout structure.
type V1OidcLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OidcLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1OidcLogoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OidcLogoutNoContent creates a V1OidcLogoutNoContent with default headers values
func NewV1OidcLogoutNoContent() *V1OidcLogoutNoContent {
	return &V1OidcLogoutNoContent{}
}

/*
V1OidcLogoutNoContent handles this case with default header values.

Ok response without content
*/
type V1OidcLogoutNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1OidcLogoutNoContent) Error() string {
	return fmt.Sprintf("[GET /v1/auth/org/{org}/oidc/logout][%d] v1OidcLogoutNoContent ", 204)
}

func (o *V1OidcLogoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
