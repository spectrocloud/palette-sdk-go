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

// V1CloudConfigsGkeMachinePoolCreateReader is a Reader for the V1CloudConfigsGkeMachinePoolCreate structure.
type V1CloudConfigsGkeMachinePoolCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGkeMachinePoolCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1CloudConfigsGkeMachinePoolCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGkeMachinePoolCreateCreated creates a V1CloudConfigsGkeMachinePoolCreateCreated with default headers values
func NewV1CloudConfigsGkeMachinePoolCreateCreated() *V1CloudConfigsGkeMachinePoolCreateCreated {
	return &V1CloudConfigsGkeMachinePoolCreateCreated{}
}

/*
V1CloudConfigsGkeMachinePoolCreateCreated handles this case with default header values.

Created successfully
*/
type V1CloudConfigsGkeMachinePoolCreateCreated struct {
	/*Audit uid for the request
	 */
	AuditUID string

	Payload *models.V1UID
}

func (o *V1CloudConfigsGkeMachinePoolCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/cloudconfigs/gke/{configUid}/machinePools][%d] v1CloudConfigsGkeMachinePoolCreateCreated  %+v", 201, o.Payload)
}

func (o *V1CloudConfigsGkeMachinePoolCreateCreated) GetPayload() *models.V1UID {
	return o.Payload
}

func (o *V1CloudConfigsGkeMachinePoolCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	o.Payload = new(models.V1UID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
