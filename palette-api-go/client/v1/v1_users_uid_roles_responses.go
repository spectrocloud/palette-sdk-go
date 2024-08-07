// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1UsersUIDRolesReader is a Reader for the V1UsersUIDRoles structure.
type V1UsersUIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersUIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersUIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersUIDRolesOK creates a V1UsersUIDRolesOK with default headers values
func NewV1UsersUIDRolesOK() *V1UsersUIDRolesOK {
	return &V1UsersUIDRolesOK{}
}

/*V1UsersUIDRolesOK handles this case with default header values.

OK
*/
type V1UsersUIDRolesOK struct {
	Payload *models.V1UserRolesEntity
}

func (o *V1UsersUIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{uid}/roles][%d] v1UsersUidRolesOK  %+v", 200, o.Payload)
}

func (o *V1UsersUIDRolesOK) GetPayload() *models.V1UserRolesEntity {
	return o.Payload
}

func (o *V1UsersUIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserRolesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
