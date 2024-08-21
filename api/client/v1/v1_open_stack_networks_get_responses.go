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

// V1OpenStackNetworksGetReader is a Reader for the V1OpenStackNetworksGet structure.
type V1OpenStackNetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenStackNetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenStackNetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenStackNetworksGetOK creates a V1OpenStackNetworksGetOK with default headers values
func NewV1OpenStackNetworksGetOK() *V1OpenStackNetworksGetOK {
	return &V1OpenStackNetworksGetOK{}
}

/*
V1OpenStackNetworksGetOK handles this case with default header values.

(empty)
*/
type V1OpenStackNetworksGetOK struct {
	Payload *models.V1OpenStackNetworks
}

func (o *V1OpenStackNetworksGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/openstack/networks][%d] v1OpenStackNetworksGetOK  %+v", 200, o.Payload)
}

func (o *V1OpenStackNetworksGetOK) GetPayload() *models.V1OpenStackNetworks {
	return o.Payload
}

func (o *V1OpenStackNetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}