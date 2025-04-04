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

// V1OciRegistriesSummaryReader is a Reader for the V1OciRegistriesSummary structure.
type V1OciRegistriesSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OciRegistriesSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OciRegistriesSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OciRegistriesSummaryOK creates a V1OciRegistriesSummaryOK with default headers values
func NewV1OciRegistriesSummaryOK() *V1OciRegistriesSummaryOK {
	return &V1OciRegistriesSummaryOK{}
}

/*
V1OciRegistriesSummaryOK handles this case with default header values.

An array of oci registry items
*/
type V1OciRegistriesSummaryOK struct {
	Payload *models.V1OciRegistries
}

func (o *V1OciRegistriesSummaryOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/oci/summary][%d] v1OciRegistriesSummaryOK  %+v", 200, o.Payload)
}

func (o *V1OciRegistriesSummaryOK) GetPayload() *models.V1OciRegistries {
	return o.Payload
}

func (o *V1OciRegistriesSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OciRegistries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
