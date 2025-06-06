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

// V1GcpRegionsReader is a Reader for the V1GcpRegions structure.
type V1GcpRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpRegionsOK creates a V1GcpRegionsOK with default headers values
func NewV1GcpRegionsOK() *V1GcpRegionsOK {
	return &V1GcpRegionsOK{}
}

/*
V1GcpRegionsOK handles this case with default header values.

(empty)
*/
type V1GcpRegionsOK struct {
	Payload *models.V1GcpRegions
}

func (o *V1GcpRegionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/projects/{project}/regions][%d] v1GcpRegionsOK  %+v", 200, o.Payload)
}

func (o *V1GcpRegionsOK) GetPayload() *models.V1GcpRegions {
	return o.Payload
}

func (o *V1GcpRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
