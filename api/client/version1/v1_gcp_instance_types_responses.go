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

// V1GcpInstanceTypesReader is a Reader for the V1GcpInstanceTypes structure.
type V1GcpInstanceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpInstanceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpInstanceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpInstanceTypesOK creates a V1GcpInstanceTypesOK with default headers values
func NewV1GcpInstanceTypesOK() *V1GcpInstanceTypesOK {
	return &V1GcpInstanceTypesOK{}
}

/*
V1GcpInstanceTypesOK handles this case with default header values.

(empty)
*/
type V1GcpInstanceTypesOK struct {
	Payload *models.V1GcpInstanceTypes
}

func (o *V1GcpInstanceTypesOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/regions/{region}/instancetypes][%d] v1GcpInstanceTypesOK  %+v", 200, o.Payload)
}

func (o *V1GcpInstanceTypesOK) GetPayload() *models.V1GcpInstanceTypes {
	return o.Payload
}

func (o *V1GcpInstanceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpInstanceTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
