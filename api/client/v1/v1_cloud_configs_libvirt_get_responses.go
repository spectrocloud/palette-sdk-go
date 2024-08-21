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

// V1CloudConfigsLibvirtGetReader is a Reader for the V1CloudConfigsLibvirtGet structure.
type V1CloudConfigsLibvirtGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsLibvirtGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsLibvirtGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsLibvirtGetOK creates a V1CloudConfigsLibvirtGetOK with default headers values
func NewV1CloudConfigsLibvirtGetOK() *V1CloudConfigsLibvirtGetOK {
	return &V1CloudConfigsLibvirtGetOK{}
}

/*
V1CloudConfigsLibvirtGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsLibvirtGetOK struct {
	Payload *models.V1LibvirtCloudConfig
}

func (o *V1CloudConfigsLibvirtGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/libvirt/{configUid}][%d] v1CloudConfigsLibvirtGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsLibvirtGetOK) GetPayload() *models.V1LibvirtCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsLibvirtGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1LibvirtCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}