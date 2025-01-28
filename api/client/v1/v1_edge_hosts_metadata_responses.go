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

// V1EdgeHostsMetadataReader is a Reader for the V1EdgeHostsMetadata structure.
type V1EdgeHostsMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeHostsMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EdgeHostsMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeHostsMetadataOK creates a V1EdgeHostsMetadataOK with default headers values
func NewV1EdgeHostsMetadataOK() *V1EdgeHostsMetadataOK {
	return &V1EdgeHostsMetadataOK{}
}

/*V1EdgeHostsMetadataOK handles this case with default header values.

An array of edgenative pair summary items
*/
type V1EdgeHostsMetadataOK struct {
	Payload *models.V1EdgeHostsMetadataSummary
}

func (o *V1EdgeHostsMetadataOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/appliances/metadata][%d] v1EdgeHostsMetadataOK  %+v", 200, o.Payload)
}

func (o *V1EdgeHostsMetadataOK) GetPayload() *models.V1EdgeHostsMetadataSummary {
	return o.Payload
}

func (o *V1EdgeHostsMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeHostsMetadataSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
