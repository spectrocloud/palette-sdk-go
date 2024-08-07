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

// V1UsersMeGetReader is a Reader for the V1UsersMeGet structure.
type V1UsersMeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersMeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersMeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersMeGetOK creates a V1UsersMeGetOK with default headers values
func NewV1UsersMeGetOK() *V1UsersMeGetOK {
	return &V1UsersMeGetOK{}
}

/*
V1UsersMeGetOK handles this case with default header values.

OK
*/
type V1UsersMeGetOK struct {
	Payload *models.V1UserMe
}

func (o *V1UsersMeGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/me][%d] v1UsersMeGetOK  %+v", 200, o.Payload)
}

func (o *V1UsersMeGetOK) GetPayload() *models.V1UserMe {
	return o.Payload
}

func (o *V1UsersMeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserMe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
