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

// V1EdgeHostDevicesUIDConfigGetReader is a Reader for the V1EdgeHostDevicesUIDConfigGet structure.
type V1EdgeHostDevicesUIDConfigGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeHostDevicesUIDConfigGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EdgeHostDevicesUIDConfigGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeHostDevicesUIDConfigGetOK creates a V1EdgeHostDevicesUIDConfigGetOK with default headers values
func NewV1EdgeHostDevicesUIDConfigGetOK() *V1EdgeHostDevicesUIDConfigGetOK {
	return &V1EdgeHostDevicesUIDConfigGetOK{}
}

/*
V1EdgeHostDevicesUIDConfigGetOK handles this case with default header values.

OK
*/
type V1EdgeHostDevicesUIDConfigGetOK struct {
	Payload *models.V1EdgeHostConfig
}

func (o *V1EdgeHostDevicesUIDConfigGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/edgehosts/{uid}/config][%d] v1EdgeHostDevicesUidConfigGetOK  %+v", 200, o.Payload)
}

func (o *V1EdgeHostDevicesUIDConfigGetOK) GetPayload() *models.V1EdgeHostConfig {
	return o.Payload
}

func (o *V1EdgeHostDevicesUIDConfigGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeHostConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
