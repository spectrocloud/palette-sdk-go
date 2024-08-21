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

// V1CloudConfigsGcpPoolMachinesUIDGetReader is a Reader for the V1CloudConfigsGcpPoolMachinesUIDGet structure.
type V1CloudConfigsGcpPoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGcpPoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsGcpPoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGcpPoolMachinesUIDGetOK creates a V1CloudConfigsGcpPoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsGcpPoolMachinesUIDGetOK() *V1CloudConfigsGcpPoolMachinesUIDGetOK {
	return &V1CloudConfigsGcpPoolMachinesUIDGetOK{}
}

/*
V1CloudConfigsGcpPoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsGcpPoolMachinesUIDGetOK struct {
	Payload *models.V1GcpMachine
}

func (o *V1CloudConfigsGcpPoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/gcp/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsGcpPoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsGcpPoolMachinesUIDGetOK) GetPayload() *models.V1GcpMachine {
	return o.Payload
}

func (o *V1CloudConfigsGcpPoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}