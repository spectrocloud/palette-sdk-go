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

// V1GcpImageURLReader is a Reader for the V1GcpImageURL structure.
type V1GcpImageURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpImageURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpImageURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpImageURLOK creates a V1GcpImageURLOK with default headers values
func NewV1GcpImageURLOK() *V1GcpImageURLOK {
	return &V1GcpImageURLOK{}
}

/*
V1GcpImageURLOK handles this case with default header values.

(empty)
*/
type V1GcpImageURLOK struct {
	Payload *models.V1GcpImageURLEntity
}

func (o *V1GcpImageURLOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/images/{imageName}/url][%d] v1GcpImageUrlOK  %+v", 200, o.Payload)
}

func (o *V1GcpImageURLOK) GetPayload() *models.V1GcpImageURLEntity {
	return o.Payload
}

func (o *V1GcpImageURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpImageURLEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
