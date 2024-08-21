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

// V1CloudConfigsGenericPoolMachinesUIDGetReader is a Reader for the V1CloudConfigsGenericPoolMachinesUIDGet structure.
type V1CloudConfigsGenericPoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGenericPoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsGenericPoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGenericPoolMachinesUIDGetOK creates a V1CloudConfigsGenericPoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsGenericPoolMachinesUIDGetOK() *V1CloudConfigsGenericPoolMachinesUIDGetOK {
	return &V1CloudConfigsGenericPoolMachinesUIDGetOK{}
}

/*
V1CloudConfigsGenericPoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsGenericPoolMachinesUIDGetOK struct {
	Payload *models.V1GenericMachine
}

func (o *V1CloudConfigsGenericPoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/generic/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsGenericPoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsGenericPoolMachinesUIDGetOK) GetPayload() *models.V1GenericMachine {
	return o.Payload
}

func (o *V1CloudConfigsGenericPoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GenericMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}