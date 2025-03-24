// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SamlLogoutReader is a Reader for the V1SamlLogout structure.
type V1SamlLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SamlLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SamlLogoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SamlLogoutNoContent creates a V1SamlLogoutNoContent with default headers values
func NewV1SamlLogoutNoContent() *V1SamlLogoutNoContent {
	return &V1SamlLogoutNoContent{}
}

/*V1SamlLogoutNoContent handles this case with default header values.

Ok response without content
*/
type V1SamlLogoutNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1SamlLogoutNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/auth/org/{org}/saml/logout][%d] v1SamlLogoutNoContent ", 204)
}

func (o *V1SamlLogoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
