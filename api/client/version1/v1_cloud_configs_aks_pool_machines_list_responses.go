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

// V1CloudConfigsAksPoolMachinesListReader is a Reader for the V1CloudConfigsAksPoolMachinesList structure.
type V1CloudConfigsAksPoolMachinesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAksPoolMachinesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsAksPoolMachinesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAksPoolMachinesListOK creates a V1CloudConfigsAksPoolMachinesListOK with default headers values
func NewV1CloudConfigsAksPoolMachinesListOK() *V1CloudConfigsAksPoolMachinesListOK {
	return &V1CloudConfigsAksPoolMachinesListOK{}
}

/*
V1CloudConfigsAksPoolMachinesListOK handles this case with default header values.

An array of AKS machine items
*/
type V1CloudConfigsAksPoolMachinesListOK struct {
	Payload *models.V1AzureMachines
}

func (o *V1CloudConfigsAksPoolMachinesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/aks/{configUid}/machinePools/{machinePoolName}/machines][%d] v1CloudConfigsAksPoolMachinesListOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsAksPoolMachinesListOK) GetPayload() *models.V1AzureMachines {
	return o.Payload
}

func (o *V1CloudConfigsAksPoolMachinesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AzureMachines)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
