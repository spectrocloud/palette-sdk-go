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

// V1CloudConfigsTkePoolMachinesListReader is a Reader for the V1CloudConfigsTkePoolMachinesList structure.
type V1CloudConfigsTkePoolMachinesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsTkePoolMachinesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsTkePoolMachinesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsTkePoolMachinesListOK creates a V1CloudConfigsTkePoolMachinesListOK with default headers values
func NewV1CloudConfigsTkePoolMachinesListOK() *V1CloudConfigsTkePoolMachinesListOK {
	return &V1CloudConfigsTkePoolMachinesListOK{}
}

/*
V1CloudConfigsTkePoolMachinesListOK handles this case with default header values.

An array of TKE machine items
*/
type V1CloudConfigsTkePoolMachinesListOK struct {
	Payload *models.V1TencentMachines
}

func (o *V1CloudConfigsTkePoolMachinesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/tke/{configUid}/machinePools/{machinePoolName}/machines][%d] v1CloudConfigsTkePoolMachinesListOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsTkePoolMachinesListOK) GetPayload() *models.V1TencentMachines {
	return o.Payload
}

func (o *V1CloudConfigsTkePoolMachinesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TencentMachines)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
