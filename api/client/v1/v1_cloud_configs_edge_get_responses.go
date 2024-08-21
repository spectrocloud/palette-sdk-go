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

// V1CloudConfigsEdgeGetReader is a Reader for the V1CloudConfigsEdgeGet structure.
type V1CloudConfigsEdgeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEdgeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsEdgeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEdgeGetOK creates a V1CloudConfigsEdgeGetOK with default headers values
func NewV1CloudConfigsEdgeGetOK() *V1CloudConfigsEdgeGetOK {
	return &V1CloudConfigsEdgeGetOK{}
}

/*
V1CloudConfigsEdgeGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsEdgeGetOK struct {
	Payload *models.V1EdgeCloudConfig
}

func (o *V1CloudConfigsEdgeGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/edge/{configUid}][%d] v1CloudConfigsEdgeGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsEdgeGetOK) GetPayload() *models.V1EdgeCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsEdgeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}