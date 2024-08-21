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

// V1CloudConfigsCustomGetReader is a Reader for the V1CloudConfigsCustomGet structure.
type V1CloudConfigsCustomGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsCustomGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsCustomGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsCustomGetOK creates a V1CloudConfigsCustomGetOK with default headers values
func NewV1CloudConfigsCustomGetOK() *V1CloudConfigsCustomGetOK {
	return &V1CloudConfigsCustomGetOK{}
}

/*
V1CloudConfigsCustomGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsCustomGetOK struct {
	Payload *models.V1CustomCloudConfig
}

func (o *V1CloudConfigsCustomGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/cloudTypes/{cloudType}/{configUid}][%d] v1CloudConfigsCustomGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsCustomGetOK) GetPayload() *models.V1CustomCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsCustomGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CustomCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}