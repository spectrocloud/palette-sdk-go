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

// V1GcpZonesReader is a Reader for the V1GcpZones structure.
type V1GcpZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpZonesOK creates a V1GcpZonesOK with default headers values
func NewV1GcpZonesOK() *V1GcpZonesOK {
	return &V1GcpZonesOK{}
}

/*
V1GcpZonesOK handles this case with default header values.

(empty)
*/
type V1GcpZonesOK struct {
	Payload *models.V1GcpZones
}

func (o *V1GcpZonesOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/projects/{project}/regions/{region}/zones][%d] v1GcpZonesOK  %+v", 200, o.Payload)
}

func (o *V1GcpZonesOK) GetPayload() *models.V1GcpZones {
	return o.Payload
}

func (o *V1GcpZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpZones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
