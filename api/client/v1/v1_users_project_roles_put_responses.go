// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersProjectRolesPutReader is a Reader for the V1UsersProjectRolesPut structure.
type V1UsersProjectRolesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersProjectRolesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersProjectRolesPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersProjectRolesPutNoContent creates a V1UsersProjectRolesPutNoContent with default headers values
func NewV1UsersProjectRolesPutNoContent() *V1UsersProjectRolesPutNoContent {
	return &V1UsersProjectRolesPutNoContent{}
}

/*
V1UsersProjectRolesPutNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1UsersProjectRolesPutNoContent struct {
}

func (o *V1UsersProjectRolesPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{uid}/projects][%d] v1UsersProjectRolesPutNoContent ", 204)
}

func (o *V1UsersProjectRolesPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
