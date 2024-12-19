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

// V1ProjectsUIDAlertsUIDGetReader is a Reader for the V1ProjectsUIDAlertsUIDGet structure.
type V1ProjectsUIDAlertsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ProjectsUIDAlertsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ProjectsUIDAlertsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ProjectsUIDAlertsUIDGetOK creates a V1ProjectsUIDAlertsUIDGetOK with default headers values
func NewV1ProjectsUIDAlertsUIDGetOK() *V1ProjectsUIDAlertsUIDGetOK {
	return &V1ProjectsUIDAlertsUIDGetOK{}
}

/*
V1ProjectsUIDAlertsUIDGetOK handles this case with default header values.

OK
*/
type V1ProjectsUIDAlertsUIDGetOK struct {
	Payload *models.V1Channel
}

func (o *V1ProjectsUIDAlertsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{uid}/alerts/{component}/{alertUid}][%d] v1ProjectsUidAlertsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1ProjectsUIDAlertsUIDGetOK) GetPayload() *models.V1Channel {
	return o.Payload
}

func (o *V1ProjectsUIDAlertsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Channel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
