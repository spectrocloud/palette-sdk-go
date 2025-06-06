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

// V1GcpStorageTypesReader is a Reader for the V1GcpStorageTypes structure.
type V1GcpStorageTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpStorageTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpStorageTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpStorageTypesOK creates a V1GcpStorageTypesOK with default headers values
func NewV1GcpStorageTypesOK() *V1GcpStorageTypesOK {
	return &V1GcpStorageTypesOK{}
}

/*
V1GcpStorageTypesOK handles this case with default header values.

(empty)
*/
type V1GcpStorageTypesOK struct {
	Payload *models.V1GcpStorageTypes
}

func (o *V1GcpStorageTypesOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/regions/{region}/storagetypes][%d] v1GcpStorageTypesOK  %+v", 200, o.Payload)
}

func (o *V1GcpStorageTypesOK) GetPayload() *models.V1GcpStorageTypes {
	return o.Payload
}

func (o *V1GcpStorageTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpStorageTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
