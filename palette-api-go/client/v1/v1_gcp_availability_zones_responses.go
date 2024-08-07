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

// V1GcpAvailabilityZonesReader is a Reader for the V1GcpAvailabilityZones structure.
type V1GcpAvailabilityZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpAvailabilityZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpAvailabilityZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpAvailabilityZonesOK creates a V1GcpAvailabilityZonesOK with default headers values
func NewV1GcpAvailabilityZonesOK() *V1GcpAvailabilityZonesOK {
	return &V1GcpAvailabilityZonesOK{}
}

/*V1GcpAvailabilityZonesOK handles this case with default header values.

(empty)
*/
type V1GcpAvailabilityZonesOK struct {
	Payload *models.V1GcpZones
}

func (o *V1GcpAvailabilityZonesOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/projects/{project}/zones][%d] v1GcpAvailabilityZonesOK  %+v", 200, o.Payload)
}

func (o *V1GcpAvailabilityZonesOK) GetPayload() *models.V1GcpZones {
	return o.Payload
}

func (o *V1GcpAvailabilityZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpZones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
