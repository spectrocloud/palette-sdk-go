// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1BasicOciRegistriesValidateReader is a Reader for the V1BasicOciRegistriesValidate structure.
type V1BasicOciRegistriesValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1BasicOciRegistriesValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1BasicOciRegistriesValidateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1BasicOciRegistriesValidateNoContent creates a V1BasicOciRegistriesValidateNoContent with default headers values
func NewV1BasicOciRegistriesValidateNoContent() *V1BasicOciRegistriesValidateNoContent {
	return &V1BasicOciRegistriesValidateNoContent{}
}

/*V1BasicOciRegistriesValidateNoContent handles this case with default header values.

Ok response without content
*/
type V1BasicOciRegistriesValidateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1BasicOciRegistriesValidateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/registries/oci/basic/validate][%d] v1BasicOciRegistriesValidateNoContent ", 204)
}

func (o *V1BasicOciRegistriesValidateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
