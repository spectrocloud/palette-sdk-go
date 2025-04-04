// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersWorkspaceRolesPutReader is a Reader for the V1UsersWorkspaceRolesPut structure.
type V1UsersWorkspaceRolesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersWorkspaceRolesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersWorkspaceRolesPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersWorkspaceRolesPutNoContent creates a V1UsersWorkspaceRolesPutNoContent with default headers values
func NewV1UsersWorkspaceRolesPutNoContent() *V1UsersWorkspaceRolesPutNoContent {
	return &V1UsersWorkspaceRolesPutNoContent{}
}

/*
V1UsersWorkspaceRolesPutNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1UsersWorkspaceRolesPutNoContent struct {
}

func (o *V1UsersWorkspaceRolesPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/workspaces/users/{userUid}/roles][%d] v1UsersWorkspaceRolesPutNoContent ", 204)
}

func (o *V1UsersWorkspaceRolesPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
