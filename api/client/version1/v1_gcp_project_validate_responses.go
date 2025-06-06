// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1GcpProjectValidateReader is a Reader for the V1GcpProjectValidate structure.
type V1GcpProjectValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpProjectValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1GcpProjectValidateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpProjectValidateNoContent creates a V1GcpProjectValidateNoContent with default headers values
func NewV1GcpProjectValidateNoContent() *V1GcpProjectValidateNoContent {
	return &V1GcpProjectValidateNoContent{}
}

/*
V1GcpProjectValidateNoContent handles this case with default header values.

Ok response without content
*/
type V1GcpProjectValidateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1GcpProjectValidateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/clouds/gcp/projects/{project}/validate][%d] v1GcpProjectValidateNoContent ", 204)
}

func (o *V1GcpProjectValidateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
