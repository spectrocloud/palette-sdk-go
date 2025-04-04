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

// V1VsphereMappingGetReader is a Reader for the V1VsphereMappingGet structure.
type V1VsphereMappingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VsphereMappingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VsphereMappingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VsphereMappingGetOK creates a V1VsphereMappingGetOK with default headers values
func NewV1VsphereMappingGetOK() *V1VsphereMappingGetOK {
	return &V1VsphereMappingGetOK{}
}

/*
V1VsphereMappingGetOK handles this case with default header values.

(empty)
*/
type V1VsphereMappingGetOK struct {
	Payload *models.V1VsphereDNSMapping
}

func (o *V1VsphereMappingGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/assets/vsphere/dnsMapping][%d] v1VsphereMappingGetOK  %+v", 200, o.Payload)
}

func (o *V1VsphereMappingGetOK) GetPayload() *models.V1VsphereDNSMapping {
	return o.Payload
}

func (o *V1VsphereMappingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VsphereDNSMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
