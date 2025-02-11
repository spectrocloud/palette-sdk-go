// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsCustomPoolMachinesUIDUpdateReader is a Reader for the V1CloudConfigsCustomPoolMachinesUIDUpdate structure.
type V1CloudConfigsCustomPoolMachinesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsCustomPoolMachinesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsCustomPoolMachinesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsCustomPoolMachinesUIDUpdateNoContent creates a V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent with default headers values
func NewV1CloudConfigsCustomPoolMachinesUIDUpdateNoContent() *V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent {
	return &V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent{}
}

/*
V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent struct {
}

func (o *V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/cloudTypes/{cloudType}/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsCustomPoolMachinesUidUpdateNoContent ", 204)
}

func (o *V1CloudConfigsCustomPoolMachinesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
