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

// V1TeamsProjectRolesReader is a Reader for the V1TeamsProjectRoles structure.
type V1TeamsProjectRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsProjectRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1TeamsProjectRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsProjectRolesOK creates a V1TeamsProjectRolesOK with default headers values
func NewV1TeamsProjectRolesOK() *V1TeamsProjectRolesOK {
	return &V1TeamsProjectRolesOK{}
}

/*
V1TeamsProjectRolesOK handles this case with default header values.

OK
*/
type V1TeamsProjectRolesOK struct {
	Payload *models.V1ProjectRolesEntity
}

func (o *V1TeamsProjectRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{uid}/projects][%d] v1TeamsProjectRolesOK  %+v", 200, o.Payload)
}

func (o *V1TeamsProjectRolesOK) GetPayload() *models.V1ProjectRolesEntity {
	return o.Payload
}

func (o *V1TeamsProjectRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProjectRolesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}