// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAzurePoolMachinesUIDUpdateReader is a Reader for the V1CloudConfigsAzurePoolMachinesUIDUpdate structure.
type V1CloudConfigsAzurePoolMachinesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAzurePoolMachinesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAzurePoolMachinesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAzurePoolMachinesUIDUpdateNoContent creates a V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent with default headers values
func NewV1CloudConfigsAzurePoolMachinesUIDUpdateNoContent() *V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent {
	return &V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent{}
}

/*
V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent struct {
}

func (o *V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/azure/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsAzurePoolMachinesUidUpdateNoContent ", 204)
}

func (o *V1CloudConfigsAzurePoolMachinesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
