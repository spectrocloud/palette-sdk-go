// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TeamsUIDResourceRolesCreateReader is a Reader for the V1TeamsUIDResourceRolesCreate structure.
type V1TeamsUIDResourceRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsUIDResourceRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TeamsUIDResourceRolesCreateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsUIDResourceRolesCreateNoContent creates a V1TeamsUIDResourceRolesCreateNoContent with default headers values
func NewV1TeamsUIDResourceRolesCreateNoContent() *V1TeamsUIDResourceRolesCreateNoContent {
	return &V1TeamsUIDResourceRolesCreateNoContent{}
}

/*
V1TeamsUIDResourceRolesCreateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TeamsUIDResourceRolesCreateNoContent struct {
}

func (o *V1TeamsUIDResourceRolesCreateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{uid}/resourceRoles][%d] v1TeamsUidResourceRolesCreateNoContent ", 204)
}

func (o *V1TeamsUIDResourceRolesCreateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
