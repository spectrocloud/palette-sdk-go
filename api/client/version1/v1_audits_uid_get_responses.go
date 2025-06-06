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

// V1AuditsUIDGetReader is a Reader for the V1AuditsUIDGet structure.
type V1AuditsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AuditsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AuditsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AuditsUIDGetOK creates a V1AuditsUIDGetOK with default headers values
func NewV1AuditsUIDGetOK() *V1AuditsUIDGetOK {
	return &V1AuditsUIDGetOK{}
}

/*
V1AuditsUIDGetOK handles this case with default header values.

(empty)
*/
type V1AuditsUIDGetOK struct {
	Payload *models.V1Audit
}

func (o *V1AuditsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/audits/{uid}][%d] v1AuditsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1AuditsUIDGetOK) GetPayload() *models.V1Audit {
	return o.Payload
}

func (o *V1AuditsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Audit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
