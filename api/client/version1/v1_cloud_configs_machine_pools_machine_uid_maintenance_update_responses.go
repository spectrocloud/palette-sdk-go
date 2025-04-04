// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateReader is a Reader for the V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdate structure.
type V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent creates a V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent with default headers values
func NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent() *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent {
	return &V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent{}
}

/*
V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent struct {
}

func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/{cloudType}/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}/maintenance][%d] v1CloudConfigsMachinePoolsMachineUidMaintenanceUpdateNoContent ", 204)
}

func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
