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

// V1AwsZonesReader is a Reader for the V1AwsZones structure.
type V1AwsZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsZonesOK creates a V1AwsZonesOK with default headers values
func NewV1AwsZonesOK() *V1AwsZonesOK {
	return &V1AwsZonesOK{}
}

/*V1AwsZonesOK handles this case with default header values.

(empty)
*/
type V1AwsZonesOK struct {
	Payload *models.V1AwsAvailabilityZones
}

func (o *V1AwsZonesOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/aws/regions/{region}/availabilityzones][%d] v1AwsZonesOK  %+v", 200, o.Payload)
}

func (o *V1AwsZonesOK) GetPayload() *models.V1AwsAvailabilityZones {
	return o.Payload
}

func (o *V1AwsZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsAvailabilityZones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
