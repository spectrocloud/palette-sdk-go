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

// V1CloudConfigsGenericGetReader is a Reader for the V1CloudConfigsGenericGet structure.
type V1CloudConfigsGenericGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGenericGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsGenericGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGenericGetOK creates a V1CloudConfigsGenericGetOK with default headers values
func NewV1CloudConfigsGenericGetOK() *V1CloudConfigsGenericGetOK {
	return &V1CloudConfigsGenericGetOK{}
}

/*
V1CloudConfigsGenericGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsGenericGetOK struct {
	Payload *models.V1GenericCloudConfig
}

func (o *V1CloudConfigsGenericGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/generic/{configUid}][%d] v1CloudConfigsGenericGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsGenericGetOK) GetPayload() *models.V1GenericCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsGenericGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GenericCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
