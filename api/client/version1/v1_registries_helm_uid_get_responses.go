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

// V1RegistriesHelmUIDGetReader is a Reader for the V1RegistriesHelmUIDGet structure.
type V1RegistriesHelmUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RegistriesHelmUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1RegistriesHelmUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RegistriesHelmUIDGetOK creates a V1RegistriesHelmUIDGetOK with default headers values
func NewV1RegistriesHelmUIDGetOK() *V1RegistriesHelmUIDGetOK {
	return &V1RegistriesHelmUIDGetOK{}
}

/*
V1RegistriesHelmUIDGetOK handles this case with default header values.

OK
*/
type V1RegistriesHelmUIDGetOK struct {
	Payload *models.V1HelmRegistry
}

func (o *V1RegistriesHelmUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/helm/{uid}][%d] v1RegistriesHelmUidGetOK  %+v", 200, o.Payload)
}

func (o *V1RegistriesHelmUIDGetOK) GetPayload() *models.V1HelmRegistry {
	return o.Payload
}

func (o *V1RegistriesHelmUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1HelmRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
