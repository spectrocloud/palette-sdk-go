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

// V1ProjectsUIDMacrosListReader is a Reader for the V1ProjectsUIDMacrosList structure.
type V1ProjectsUIDMacrosListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ProjectsUIDMacrosListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ProjectsUIDMacrosListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ProjectsUIDMacrosListOK creates a V1ProjectsUIDMacrosListOK with default headers values
func NewV1ProjectsUIDMacrosListOK() *V1ProjectsUIDMacrosListOK {
	return &V1ProjectsUIDMacrosListOK{}
}

/*
V1ProjectsUIDMacrosListOK handles this case with default header values.

OK
*/
type V1ProjectsUIDMacrosListOK struct {
	Payload *models.V1Macros
}

func (o *V1ProjectsUIDMacrosListOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{uid}/macros][%d] v1ProjectsUidMacrosListOK  %+v", 200, o.Payload)
}

func (o *V1ProjectsUIDMacrosListOK) GetPayload() *models.V1Macros {
	return o.Payload
}

func (o *V1ProjectsUIDMacrosListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Macros)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}