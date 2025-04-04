// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsGcpMachinePoolDeleteReader is a Reader for the V1CloudConfigsGcpMachinePoolDelete structure.
type V1CloudConfigsGcpMachinePoolDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGcpMachinePoolDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsGcpMachinePoolDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGcpMachinePoolDeleteNoContent creates a V1CloudConfigsGcpMachinePoolDeleteNoContent with default headers values
func NewV1CloudConfigsGcpMachinePoolDeleteNoContent() *V1CloudConfigsGcpMachinePoolDeleteNoContent {
	return &V1CloudConfigsGcpMachinePoolDeleteNoContent{}
}

/*
V1CloudConfigsGcpMachinePoolDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsGcpMachinePoolDeleteNoContent struct {
}

func (o *V1CloudConfigsGcpMachinePoolDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/gcp/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsGcpMachinePoolDeleteNoContent ", 204)
}

func (o *V1CloudConfigsGcpMachinePoolDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
