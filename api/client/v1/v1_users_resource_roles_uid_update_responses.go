// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersResourceRolesUIDUpdateReader is a Reader for the V1UsersResourceRolesUIDUpdate structure.
type V1UsersResourceRolesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersResourceRolesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersResourceRolesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersResourceRolesUIDUpdateNoContent creates a V1UsersResourceRolesUIDUpdateNoContent with default headers values
func NewV1UsersResourceRolesUIDUpdateNoContent() *V1UsersResourceRolesUIDUpdateNoContent {
	return &V1UsersResourceRolesUIDUpdateNoContent{}
}

/*
V1UsersResourceRolesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1UsersResourceRolesUIDUpdateNoContent struct {
}

func (o *V1UsersResourceRolesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/{uid}/resourceRoles/{resourceRoleUid}][%d] v1UsersResourceRolesUidUpdateNoContent ", 204)
}

func (o *V1UsersResourceRolesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}