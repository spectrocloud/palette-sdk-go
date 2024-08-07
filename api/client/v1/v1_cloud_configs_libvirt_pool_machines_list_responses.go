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

// V1CloudConfigsLibvirtPoolMachinesListReader is a Reader for the V1CloudConfigsLibvirtPoolMachinesList structure.
type V1CloudConfigsLibvirtPoolMachinesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsLibvirtPoolMachinesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsLibvirtPoolMachinesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsLibvirtPoolMachinesListOK creates a V1CloudConfigsLibvirtPoolMachinesListOK with default headers values
func NewV1CloudConfigsLibvirtPoolMachinesListOK() *V1CloudConfigsLibvirtPoolMachinesListOK {
	return &V1CloudConfigsLibvirtPoolMachinesListOK{}
}

/*V1CloudConfigsLibvirtPoolMachinesListOK handles this case with default header values.

An array of Libvirt machine items
*/
type V1CloudConfigsLibvirtPoolMachinesListOK struct {
	Payload *models.V1LibvirtMachines
}

func (o *V1CloudConfigsLibvirtPoolMachinesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/libvirt/{configUid}/machinePools/{machinePoolName}/machines][%d] v1CloudConfigsLibvirtPoolMachinesListOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsLibvirtPoolMachinesListOK) GetPayload() *models.V1LibvirtMachines {
	return o.Payload
}

func (o *V1CloudConfigsLibvirtPoolMachinesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1LibvirtMachines)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
