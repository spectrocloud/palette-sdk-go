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

// V1OpenStackFlavorsGetReader is a Reader for the V1OpenStackFlavorsGet structure.
type V1OpenStackFlavorsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenStackFlavorsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenStackFlavorsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenStackFlavorsGetOK creates a V1OpenStackFlavorsGetOK with default headers values
func NewV1OpenStackFlavorsGetOK() *V1OpenStackFlavorsGetOK {
	return &V1OpenStackFlavorsGetOK{}
}

/*
V1OpenStackFlavorsGetOK handles this case with default header values.

(empty)
*/
type V1OpenStackFlavorsGetOK struct {
	Payload *models.V1OpenStackFlavors
}

func (o *V1OpenStackFlavorsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/openstack/flavors][%d] v1OpenStackFlavorsGetOK  %+v", 200, o.Payload)
}

func (o *V1OpenStackFlavorsGetOK) GetPayload() *models.V1OpenStackFlavors {
	return o.Payload
}

func (o *V1OpenStackFlavorsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackFlavors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
