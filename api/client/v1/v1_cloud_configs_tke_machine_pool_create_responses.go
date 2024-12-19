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

// V1CloudConfigsTkeMachinePoolCreateReader is a Reader for the V1CloudConfigsTkeMachinePoolCreate structure.
type V1CloudConfigsTkeMachinePoolCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsTkeMachinePoolCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1CloudConfigsTkeMachinePoolCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsTkeMachinePoolCreateCreated creates a V1CloudConfigsTkeMachinePoolCreateCreated with default headers values
func NewV1CloudConfigsTkeMachinePoolCreateCreated() *V1CloudConfigsTkeMachinePoolCreateCreated {
	return &V1CloudConfigsTkeMachinePoolCreateCreated{}
}

/*V1CloudConfigsTkeMachinePoolCreateCreated handles this case with default header values.

Created successfully
*/
type V1CloudConfigsTkeMachinePoolCreateCreated struct {
	/*Audit uid for the request
	 */
	AuditUID string

	Payload *models.V1UID
}

func (o *V1CloudConfigsTkeMachinePoolCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/cloudconfigs/tke/{configUid}/machinePools][%d] v1CloudConfigsTkeMachinePoolCreateCreated  %+v", 201, o.Payload)
}

func (o *V1CloudConfigsTkeMachinePoolCreateCreated) GetPayload() *models.V1UID {
	return o.Payload
}

func (o *V1CloudConfigsTkeMachinePoolCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	o.Payload = new(models.V1UID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
