// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1OpenStackAccountValidateReader is a Reader for the V1OpenStackAccountValidate structure.
type V1OpenStackAccountValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenStackAccountValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1OpenStackAccountValidateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenStackAccountValidateNoContent creates a V1OpenStackAccountValidateNoContent with default headers values
func NewV1OpenStackAccountValidateNoContent() *V1OpenStackAccountValidateNoContent {
	return &V1OpenStackAccountValidateNoContent{}
}

/*
V1OpenStackAccountValidateNoContent handles this case with default header values.

Ok response without content
*/
type V1OpenStackAccountValidateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1OpenStackAccountValidateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/clouds/openstack/account/validate][%d] v1OpenStackAccountValidateNoContent ", 204)
}

func (o *V1OpenStackAccountValidateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}