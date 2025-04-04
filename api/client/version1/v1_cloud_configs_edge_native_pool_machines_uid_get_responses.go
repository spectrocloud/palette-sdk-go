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

// V1CloudConfigsEdgeNativePoolMachinesUIDGetReader is a Reader for the V1CloudConfigsEdgeNativePoolMachinesUIDGet structure.
type V1CloudConfigsEdgeNativePoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsEdgeNativePoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesUIDGetOK creates a V1CloudConfigsEdgeNativePoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsEdgeNativePoolMachinesUIDGetOK() *V1CloudConfigsEdgeNativePoolMachinesUIDGetOK {
	return &V1CloudConfigsEdgeNativePoolMachinesUIDGetOK{}
}

/*
V1CloudConfigsEdgeNativePoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsEdgeNativePoolMachinesUIDGetOK struct {
	Payload *models.V1EdgeNativeMachine
}

func (o *V1CloudConfigsEdgeNativePoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/edge-native/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsEdgeNativePoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsEdgeNativePoolMachinesUIDGetOK) GetPayload() *models.V1EdgeNativeMachine {
	return o.Payload
}

func (o *V1CloudConfigsEdgeNativePoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeNativeMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
