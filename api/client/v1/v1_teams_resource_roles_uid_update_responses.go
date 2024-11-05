// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TeamsResourceRolesUIDUpdateReader is a Reader for the V1TeamsResourceRolesUIDUpdate structure.
type V1TeamsResourceRolesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsResourceRolesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TeamsResourceRolesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsResourceRolesUIDUpdateNoContent creates a V1TeamsResourceRolesUIDUpdateNoContent with default headers values
func NewV1TeamsResourceRolesUIDUpdateNoContent() *V1TeamsResourceRolesUIDUpdateNoContent {
	return &V1TeamsResourceRolesUIDUpdateNoContent{}
}

/*V1TeamsResourceRolesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TeamsResourceRolesUIDUpdateNoContent struct {
}

func (o *V1TeamsResourceRolesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/teams/{uid}/resourceRoles/{resourceRoleUid}][%d] v1TeamsResourceRolesUidUpdateNoContent ", 204)
}

func (o *V1TeamsResourceRolesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
