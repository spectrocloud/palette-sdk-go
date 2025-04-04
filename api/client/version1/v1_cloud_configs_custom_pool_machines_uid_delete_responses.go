// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsCustomPoolMachinesUIDDeleteReader is a Reader for the V1CloudConfigsCustomPoolMachinesUIDDelete structure.
type V1CloudConfigsCustomPoolMachinesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsCustomPoolMachinesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsCustomPoolMachinesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsCustomPoolMachinesUIDDeleteNoContent creates a V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent with default headers values
func NewV1CloudConfigsCustomPoolMachinesUIDDeleteNoContent() *V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent {
	return &V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent{}
}

/*
V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent struct {
}

func (o *V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/cloudTypes/{cloudType}/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsCustomPoolMachinesUidDeleteNoContent ", 204)
}

func (o *V1CloudConfigsCustomPoolMachinesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
