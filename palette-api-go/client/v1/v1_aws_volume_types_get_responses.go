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

// V1AwsVolumeTypesGetReader is a Reader for the V1AwsVolumeTypesGet structure.
type V1AwsVolumeTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsVolumeTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsVolumeTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsVolumeTypesGetOK creates a V1AwsVolumeTypesGetOK with default headers values
func NewV1AwsVolumeTypesGetOK() *V1AwsVolumeTypesGetOK {
	return &V1AwsVolumeTypesGetOK{}
}

/*V1AwsVolumeTypesGetOK handles this case with default header values.

(empty)
*/
type V1AwsVolumeTypesGetOK struct {
	Payload *models.V1AWSVolumeTypes
}

func (o *V1AwsVolumeTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/aws/volumeTypes][%d] v1AwsVolumeTypesGetOK  %+v", 200, o.Payload)
}

func (o *V1AwsVolumeTypesGetOK) GetPayload() *models.V1AWSVolumeTypes {
	return o.Payload
}

func (o *V1AwsVolumeTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AWSVolumeTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
