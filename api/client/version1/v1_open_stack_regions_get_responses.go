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

// V1OpenStackRegionsGetReader is a Reader for the V1OpenStackRegionsGet structure.
type V1OpenStackRegionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenStackRegionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenStackRegionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenStackRegionsGetOK creates a V1OpenStackRegionsGetOK with default headers values
func NewV1OpenStackRegionsGetOK() *V1OpenStackRegionsGetOK {
	return &V1OpenStackRegionsGetOK{}
}

/*
V1OpenStackRegionsGetOK handles this case with default header values.

(empty)
*/
type V1OpenStackRegionsGetOK struct {
	Payload *models.V1OpenStackRegions
}

func (o *V1OpenStackRegionsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/openstack/regions][%d] v1OpenStackRegionsGetOK  %+v", 200, o.Payload)
}

func (o *V1OpenStackRegionsGetOK) GetPayload() *models.V1OpenStackRegions {
	return o.Payload
}

func (o *V1OpenStackRegionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
