// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1RolesListReader is a Reader for the V1RolesList structure.
type V1RolesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RolesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1RolesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RolesListOK creates a V1RolesListOK with default headers values
func NewV1RolesListOK() *V1RolesListOK {
	return &V1RolesListOK{}
}

/*
V1RolesListOK handles this case with default header values.

OK
*/
type V1RolesListOK struct {
	Payload *models.V1Roles
}

func (o *V1RolesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] v1RolesListOK  %+v", 200, o.Payload)
}

func (o *V1RolesListOK) GetPayload() *models.V1Roles {
	return o.Payload
}

func (o *V1RolesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Roles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
