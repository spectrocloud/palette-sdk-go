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

// V1RegistriesMetadataReader is a Reader for the V1RegistriesMetadata structure.
type V1RegistriesMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RegistriesMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1RegistriesMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RegistriesMetadataOK creates a V1RegistriesMetadataOK with default headers values
func NewV1RegistriesMetadataOK() *V1RegistriesMetadataOK {
	return &V1RegistriesMetadataOK{}
}

/*
V1RegistriesMetadataOK handles this case with default header values.

An array of registry metadata items
*/
type V1RegistriesMetadataOK struct {
	Payload *models.V1RegistriesMetadata
}

func (o *V1RegistriesMetadataOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/metadata][%d] v1RegistriesMetadataOK  %+v", 200, o.Payload)
}

func (o *V1RegistriesMetadataOK) GetPayload() *models.V1RegistriesMetadata {
	return o.Payload
}

func (o *V1RegistriesMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RegistriesMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
