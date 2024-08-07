// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1GcpNetworksReader is a Reader for the V1GcpNetworks structure.
type V1GcpNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpNetworksOK creates a V1GcpNetworksOK with default headers values
func NewV1GcpNetworksOK() *V1GcpNetworksOK {
	return &V1GcpNetworksOK{}
}

/*V1GcpNetworksOK handles this case with default header values.

(empty)
*/
type V1GcpNetworksOK struct {
	Payload *models.V1GcpNetworks
}

func (o *V1GcpNetworksOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/projects/{project}/regions/{region}/networks][%d] v1GcpNetworksOK  %+v", 200, o.Payload)
}

func (o *V1GcpNetworksOK) GetPayload() *models.V1GcpNetworks {
	return o.Payload
}

func (o *V1GcpNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
