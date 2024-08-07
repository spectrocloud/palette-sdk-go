// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1CloudConfigsMaasPoolMachinesUIDGetReader is a Reader for the V1CloudConfigsMaasPoolMachinesUIDGet structure.
type V1CloudConfigsMaasPoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsMaasPoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsMaasPoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsMaasPoolMachinesUIDGetOK creates a V1CloudConfigsMaasPoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsMaasPoolMachinesUIDGetOK() *V1CloudConfigsMaasPoolMachinesUIDGetOK {
	return &V1CloudConfigsMaasPoolMachinesUIDGetOK{}
}

/*V1CloudConfigsMaasPoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsMaasPoolMachinesUIDGetOK struct {
	Payload *models.V1MaasMachine
}

func (o *V1CloudConfigsMaasPoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/maas/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsMaasPoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsMaasPoolMachinesUIDGetOK) GetPayload() *models.V1MaasMachine {
	return o.Payload
}

func (o *V1CloudConfigsMaasPoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MaasMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
