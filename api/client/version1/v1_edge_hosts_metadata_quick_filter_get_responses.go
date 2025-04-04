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

// V1EdgeHostsMetadataQuickFilterGetReader is a Reader for the V1EdgeHostsMetadataQuickFilterGet structure.
type V1EdgeHostsMetadataQuickFilterGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeHostsMetadataQuickFilterGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EdgeHostsMetadataQuickFilterGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeHostsMetadataQuickFilterGetOK creates a V1EdgeHostsMetadataQuickFilterGetOK with default headers values
func NewV1EdgeHostsMetadataQuickFilterGetOK() *V1EdgeHostsMetadataQuickFilterGetOK {
	return &V1EdgeHostsMetadataQuickFilterGetOK{}
}

/*
V1EdgeHostsMetadataQuickFilterGetOK handles this case with default header values.

An array of edge host metadata
*/
type V1EdgeHostsMetadataQuickFilterGetOK struct {
	Payload *models.V1EdgeHostsMeta
}

func (o *V1EdgeHostsMetadataQuickFilterGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/edgehosts/metadata][%d] v1EdgeHostsMetadataQuickFilterGetOK  %+v", 200, o.Payload)
}

func (o *V1EdgeHostsMetadataQuickFilterGetOK) GetPayload() *models.V1EdgeHostsMeta {
	return o.Payload
}

func (o *V1EdgeHostsMetadataQuickFilterGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeHostsMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
