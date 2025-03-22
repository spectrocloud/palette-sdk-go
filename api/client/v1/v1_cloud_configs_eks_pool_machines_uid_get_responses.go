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

// V1CloudConfigsEksPoolMachinesUIDGetReader is a Reader for the V1CloudConfigsEksPoolMachinesUIDGet structure.
type V1CloudConfigsEksPoolMachinesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEksPoolMachinesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsEksPoolMachinesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDGetOK creates a V1CloudConfigsEksPoolMachinesUIDGetOK with default headers values
func NewV1CloudConfigsEksPoolMachinesUIDGetOK() *V1CloudConfigsEksPoolMachinesUIDGetOK {
	return &V1CloudConfigsEksPoolMachinesUIDGetOK{}
}

/*V1CloudConfigsEksPoolMachinesUIDGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsEksPoolMachinesUIDGetOK struct {
	Payload *models.V1AwsMachine
}

func (o *V1CloudConfigsEksPoolMachinesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/eks/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsEksPoolMachinesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsEksPoolMachinesUIDGetOK) GetPayload() *models.V1AwsMachine {
	return o.Payload
}

func (o *V1CloudConfigsEksPoolMachinesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
