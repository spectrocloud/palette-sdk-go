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

// V1CloudConfigsCoxEdgePoolMachinesUIDGetReader is a Reader for the V1CloudConfigsCoxEdgePoolMachinesUIDGet structure.
type V1CloudConfigsCoxEdgePoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsCoxEdgePoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsCoxEdgePoolMachinesUIDGetOK creates a V1CloudConfigsCoxEdgePoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsCoxEdgePoolMachinesUIDGetOK() *V1CloudConfigsCoxEdgePoolMachinesUIDGetOK {
	return &V1CloudConfigsCoxEdgePoolMachinesUIDGetOK{}
}

/*
V1CloudConfigsCoxEdgePoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsCoxEdgePoolMachinesUIDGetOK struct {
	Payload *models.V1CoxEdgeMachine
}

func (o *V1CloudConfigsCoxEdgePoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/coxedge/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsCoxEdgePoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsCoxEdgePoolMachinesUIDGetOK) GetPayload() *models.V1CoxEdgeMachine {
	return o.Payload
}

func (o *V1CloudConfigsCoxEdgePoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CoxEdgeMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
