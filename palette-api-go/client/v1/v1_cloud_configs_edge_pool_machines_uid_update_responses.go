// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEdgePoolMachinesUIDUpdateReader is a Reader for the V1CloudConfigsEdgePoolMachinesUIDUpdate structure.
type V1CloudConfigsEdgePoolMachinesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEdgePoolMachinesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEdgePoolMachinesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEdgePoolMachinesUIDUpdateNoContent creates a V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent with default headers values
func NewV1CloudConfigsEdgePoolMachinesUIDUpdateNoContent() *V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent {
	return &V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent{}
}

/*V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent struct {
}

func (o *V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/edge/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsEdgePoolMachinesUidUpdateNoContent ", 204)
}

func (o *V1CloudConfigsEdgePoolMachinesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
