// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1UsersWorkspaceGetRolesReader is a Reader for the V1UsersWorkspaceGetRoles structure.
type V1UsersWorkspaceGetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersWorkspaceGetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersWorkspaceGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersWorkspaceGetRolesOK creates a V1UsersWorkspaceGetRolesOK with default headers values
func NewV1UsersWorkspaceGetRolesOK() *V1UsersWorkspaceGetRolesOK {
	return &V1UsersWorkspaceGetRolesOK{}
}

/*
V1UsersWorkspaceGetRolesOK handles this case with default header values.

OK
*/
type V1UsersWorkspaceGetRolesOK struct {
	Payload *models.V1WorkspaceScopeRoles
}

func (o *V1UsersWorkspaceGetRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/workspaces/users/{userUid}/roles][%d] v1UsersWorkspaceGetRolesOK  %+v", 200, o.Payload)
}

func (o *V1UsersWorkspaceGetRolesOK) GetPayload() *models.V1WorkspaceScopeRoles {
	return o.Payload
}

func (o *V1UsersWorkspaceGetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1WorkspaceScopeRoles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}