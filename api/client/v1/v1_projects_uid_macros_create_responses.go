// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ProjectsUIDMacrosCreateReader is a Reader for the V1ProjectsUIDMacrosCreate structure.
type V1ProjectsUIDMacrosCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ProjectsUIDMacrosCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ProjectsUIDMacrosCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ProjectsUIDMacrosCreateNoContent creates a V1ProjectsUIDMacrosCreateNoContent with default headers values
func NewV1ProjectsUIDMacrosCreateNoContent() *V1ProjectsUIDMacrosCreateNoContent {
	return &V1ProjectsUIDMacrosCreateNoContent{}
}

/*V1ProjectsUIDMacrosCreateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1ProjectsUIDMacrosCreateNoContent struct {
}

func (o *V1ProjectsUIDMacrosCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{uid}/macros][%d] v1ProjectsUidMacrosCreateNoContent ", 204)
}

func (o *V1ProjectsUIDMacrosCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
