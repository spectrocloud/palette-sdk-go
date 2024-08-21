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

// V1AuditsListReader is a Reader for the V1AuditsList structure.
type V1AuditsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AuditsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AuditsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AuditsListOK creates a V1AuditsListOK with default headers values
func NewV1AuditsListOK() *V1AuditsListOK {
	return &V1AuditsListOK{}
}

/*
V1AuditsListOK handles this case with default header values.

(empty)
*/
type V1AuditsListOK struct {
	Payload *models.V1Audits
}

func (o *V1AuditsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/audits][%d] v1AuditsListOK  %+v", 200, o.Payload)
}

func (o *V1AuditsListOK) GetPayload() *models.V1Audits {
	return o.Payload
}

func (o *V1AuditsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Audits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}