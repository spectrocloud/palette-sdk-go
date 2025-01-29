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

// V1EventsComponentsListReader is a Reader for the V1EventsComponentsList structure.
type V1EventsComponentsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EventsComponentsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EventsComponentsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EventsComponentsListOK creates a V1EventsComponentsListOK with default headers values
func NewV1EventsComponentsListOK() *V1EventsComponentsListOK {
	return &V1EventsComponentsListOK{}
}

/*
V1EventsComponentsListOK handles this case with default header values.

An array of component events items
*/
type V1EventsComponentsListOK struct {
	Payload *models.V1Events
}

func (o *V1EventsComponentsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/components][%d] v1EventsComponentsListOK  %+v", 200, o.Payload)
}

func (o *V1EventsComponentsListOK) GetPayload() *models.V1Events {
	return o.Payload
}

func (o *V1EventsComponentsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Events)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
