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

// V1TeamsListReader is a Reader for the V1TeamsList structure.
type V1TeamsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1TeamsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsListOK creates a V1TeamsListOK with default headers values
func NewV1TeamsListOK() *V1TeamsListOK {
	return &V1TeamsListOK{}
}

/*
V1TeamsListOK handles this case with default header values.

An array of teams
*/
type V1TeamsListOK struct {
	Payload *models.V1Teams
}

func (o *V1TeamsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams][%d] v1TeamsListOK  %+v", 200, o.Payload)
}

func (o *V1TeamsListOK) GetPayload() *models.V1Teams {
	return o.Payload
}

func (o *V1TeamsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Teams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
