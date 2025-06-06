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

// V1EcrRegistriesUIDGetReader is a Reader for the V1EcrRegistriesUIDGet structure.
type V1EcrRegistriesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EcrRegistriesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EcrRegistriesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EcrRegistriesUIDGetOK creates a V1EcrRegistriesUIDGetOK with default headers values
func NewV1EcrRegistriesUIDGetOK() *V1EcrRegistriesUIDGetOK {
	return &V1EcrRegistriesUIDGetOK{}
}

/*
V1EcrRegistriesUIDGetOK handles this case with default header values.

OK
*/
type V1EcrRegistriesUIDGetOK struct {
	Payload *models.V1EcrRegistry
}

func (o *V1EcrRegistriesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/oci/{uid}/ecr][%d] v1EcrRegistriesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1EcrRegistriesUIDGetOK) GetPayload() *models.V1EcrRegistry {
	return o.Payload
}

func (o *V1EcrRegistriesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EcrRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
