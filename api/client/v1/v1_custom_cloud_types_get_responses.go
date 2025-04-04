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

// V1CustomCloudTypesGetReader is a Reader for the V1CustomCloudTypesGet structure.
type V1CustomCloudTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CustomCloudTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CustomCloudTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CustomCloudTypesGetOK creates a V1CustomCloudTypesGetOK with default headers values
func NewV1CustomCloudTypesGetOK() *V1CustomCloudTypesGetOK {
	return &V1CustomCloudTypesGetOK{}
}

/*
V1CustomCloudTypesGetOK handles this case with default header values.

(empty)
*/
type V1CustomCloudTypesGetOK struct {
	Payload *models.V1CustomCloudTypes
}

func (o *V1CustomCloudTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/cloudTypes][%d] v1CustomCloudTypesGetOK  %+v", 200, o.Payload)
}

func (o *V1CustomCloudTypesGetOK) GetPayload() *models.V1CustomCloudTypes {
	return o.Payload
}

func (o *V1CustomCloudTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CustomCloudTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
