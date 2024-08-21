// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersSystemMacrosCreateReader is a Reader for the V1UsersSystemMacrosCreate structure.
type V1UsersSystemMacrosCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersSystemMacrosCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersSystemMacrosCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersSystemMacrosCreateNoContent creates a V1UsersSystemMacrosCreateNoContent with default headers values
func NewV1UsersSystemMacrosCreateNoContent() *V1UsersSystemMacrosCreateNoContent {
	return &V1UsersSystemMacrosCreateNoContent{}
}

/*
V1UsersSystemMacrosCreateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1UsersSystemMacrosCreateNoContent struct {
}

func (o *V1UsersSystemMacrosCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/users/system/macros][%d] v1UsersSystemMacrosCreateNoContent ", 204)
}

func (o *V1UsersSystemMacrosCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}