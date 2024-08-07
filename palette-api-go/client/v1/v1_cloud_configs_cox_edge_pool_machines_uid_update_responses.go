// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsCoxEdgePoolMachinesUIDUpdateReader is a Reader for the V1CloudConfigsCoxEdgePoolMachinesUIDUpdate structure.
type V1CloudConfigsCoxEdgePoolMachinesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent creates a V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent with default headers values
func NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent() *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent {
	return &V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent{}
}

/*V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent struct {
}

func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/coxedge/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsCoxEdgePoolMachinesUidUpdateNoContent ", 204)
}

func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
