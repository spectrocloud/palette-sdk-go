// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsGkePoolMachinesUIDUpdateReader is a Reader for the V1CloudConfigsGkePoolMachinesUIDUpdate structure.
type V1CloudConfigsGkePoolMachinesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGkePoolMachinesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsGkePoolMachinesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGkePoolMachinesUIDUpdateNoContent creates a V1CloudConfigsGkePoolMachinesUIDUpdateNoContent with default headers values
func NewV1CloudConfigsGkePoolMachinesUIDUpdateNoContent() *V1CloudConfigsGkePoolMachinesUIDUpdateNoContent {
	return &V1CloudConfigsGkePoolMachinesUIDUpdateNoContent{}
}

/*
V1CloudConfigsGkePoolMachinesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsGkePoolMachinesUIDUpdateNoContent struct {
}

func (o *V1CloudConfigsGkePoolMachinesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/gke/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsGkePoolMachinesUidUpdateNoContent ", 204)
}

func (o *V1CloudConfigsGkePoolMachinesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
