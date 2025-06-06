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

// V1VsphereEnvReader is a Reader for the V1VsphereEnv structure.
type V1VsphereEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VsphereEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VsphereEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VsphereEnvOK creates a V1VsphereEnvOK with default headers values
func NewV1VsphereEnvOK() *V1VsphereEnvOK {
	return &V1VsphereEnvOK{}
}

/*
V1VsphereEnvOK handles this case with default header values.

(empty)
*/
type V1VsphereEnvOK struct {
	Payload *models.V1VsphereEnv
}

func (o *V1VsphereEnvOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/vsphere/env][%d] v1VsphereEnvOK  %+v", 200, o.Payload)
}

func (o *V1VsphereEnvOK) GetPayload() *models.V1VsphereEnv {
	return o.Payload
}

func (o *V1VsphereEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VsphereEnv)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
