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

// V1CloudConfigsVirtualPoolMachinesUIDGetReader is a Reader for the V1CloudConfigsVirtualPoolMachinesUIDGet structure.
type V1CloudConfigsVirtualPoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsVirtualPoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsVirtualPoolMachinesUIDGetOK creates a V1CloudConfigsVirtualPoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsVirtualPoolMachinesUIDGetOK() *V1CloudConfigsVirtualPoolMachinesUIDGetOK {
	return &V1CloudConfigsVirtualPoolMachinesUIDGetOK{}
}

/*
V1CloudConfigsVirtualPoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsVirtualPoolMachinesUIDGetOK struct {
	Payload *models.V1VirtualMachine
}

func (o *V1CloudConfigsVirtualPoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/virtual/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsVirtualPoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsVirtualPoolMachinesUIDGetOK) GetPayload() *models.V1VirtualMachine {
	return o.Payload
}

func (o *V1CloudConfigsVirtualPoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VirtualMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}