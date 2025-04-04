// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsVirtualMachinePoolDeleteReader is a Reader for the V1CloudConfigsVirtualMachinePoolDelete structure.
type V1CloudConfigsVirtualMachinePoolDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsVirtualMachinePoolDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsVirtualMachinePoolDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsVirtualMachinePoolDeleteNoContent creates a V1CloudConfigsVirtualMachinePoolDeleteNoContent with default headers values
func NewV1CloudConfigsVirtualMachinePoolDeleteNoContent() *V1CloudConfigsVirtualMachinePoolDeleteNoContent {
	return &V1CloudConfigsVirtualMachinePoolDeleteNoContent{}
}

/*
V1CloudConfigsVirtualMachinePoolDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsVirtualMachinePoolDeleteNoContent struct {
}

func (o *V1CloudConfigsVirtualMachinePoolDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/virtual/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsVirtualMachinePoolDeleteNoContent ", 204)
}

func (o *V1CloudConfigsVirtualMachinePoolDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
