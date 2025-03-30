// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAzureMachinePoolUpdateReader is a Reader for the V1CloudConfigsAzureMachinePoolUpdate structure.
type V1CloudConfigsAzureMachinePoolUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAzureMachinePoolUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAzureMachinePoolUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAzureMachinePoolUpdateNoContent creates a V1CloudConfigsAzureMachinePoolUpdateNoContent with default headers values
func NewV1CloudConfigsAzureMachinePoolUpdateNoContent() *V1CloudConfigsAzureMachinePoolUpdateNoContent {
	return &V1CloudConfigsAzureMachinePoolUpdateNoContent{}
}

/*
V1CloudConfigsAzureMachinePoolUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsAzureMachinePoolUpdateNoContent struct {
}

func (o *V1CloudConfigsAzureMachinePoolUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/azure/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsAzureMachinePoolUpdateNoContent ", 204)
}

func (o *V1CloudConfigsAzureMachinePoolUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
