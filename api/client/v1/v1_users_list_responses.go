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

// V1UsersListReader is a Reader for the V1UsersList structure.
type V1UsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersListOK creates a V1UsersListOK with default headers values
func NewV1UsersListOK() *V1UsersListOK {
	return &V1UsersListOK{}
}

/*
V1UsersListOK handles this case with default header values.

OK
*/
type V1UsersListOK struct {
	Payload *models.V1Users
}

func (o *V1UsersListOK) Error() string {
	return fmt.Sprintf("[GET /v1/users][%d] v1UsersListOK  %+v", 200, o.Payload)
}

func (o *V1UsersListOK) GetPayload() *models.V1Users {
	return o.Payload
}

func (o *V1UsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Users)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}