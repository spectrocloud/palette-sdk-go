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

// V1CloudConfigsAwsGetReader is a Reader for the V1CloudConfigsAwsGet structure.
type V1CloudConfigsAwsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAwsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsAwsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAwsGetOK creates a V1CloudConfigsAwsGetOK with default headers values
func NewV1CloudConfigsAwsGetOK() *V1CloudConfigsAwsGetOK {
	return &V1CloudConfigsAwsGetOK{}
}

/*
V1CloudConfigsAwsGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsAwsGetOK struct {
	Payload *models.V1AwsCloudConfig
}

func (o *V1CloudConfigsAwsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/aws/{configUid}][%d] v1CloudConfigsAwsGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsAwsGetOK) GetPayload() *models.V1AwsCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsAwsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
