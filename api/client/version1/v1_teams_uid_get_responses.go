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

// V1TeamsUIDGetReader is a Reader for the V1TeamsUIDGet structure.
type V1TeamsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1TeamsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsUIDGetOK creates a V1TeamsUIDGetOK with default headers values
func NewV1TeamsUIDGetOK() *V1TeamsUIDGetOK {
	return &V1TeamsUIDGetOK{}
}

/*
V1TeamsUIDGetOK handles this case with default header values.

OK
*/
type V1TeamsUIDGetOK struct {
	Payload *models.V1Team
}

func (o *V1TeamsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{uid}][%d] v1TeamsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1TeamsUIDGetOK) GetPayload() *models.V1Team {
	return o.Payload
}

func (o *V1TeamsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
