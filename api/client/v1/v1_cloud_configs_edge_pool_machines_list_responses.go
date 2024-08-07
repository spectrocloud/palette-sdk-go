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

// V1CloudConfigsEdgePoolMachinesListReader is a Reader for the V1CloudConfigsEdgePoolMachinesList structure.
type V1CloudConfigsEdgePoolMachinesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEdgePoolMachinesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsEdgePoolMachinesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEdgePoolMachinesListOK creates a V1CloudConfigsEdgePoolMachinesListOK with default headers values
func NewV1CloudConfigsEdgePoolMachinesListOK() *V1CloudConfigsEdgePoolMachinesListOK {
	return &V1CloudConfigsEdgePoolMachinesListOK{}
}

/*V1CloudConfigsEdgePoolMachinesListOK handles this case with default header values.

An array of Edge machine items
*/
type V1CloudConfigsEdgePoolMachinesListOK struct {
	Payload *models.V1EdgeMachines
}

func (o *V1CloudConfigsEdgePoolMachinesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/edge/{configUid}/machinePools/{machinePoolName}/machines][%d] v1CloudConfigsEdgePoolMachinesListOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsEdgePoolMachinesListOK) GetPayload() *models.V1EdgeMachines {
	return o.Payload
}

func (o *V1CloudConfigsEdgePoolMachinesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeMachines)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
