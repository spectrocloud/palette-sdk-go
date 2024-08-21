// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAksPoolMachinesUIDDeleteReader is a Reader for the V1CloudConfigsAksPoolMachinesUIDDelete structure.
type V1CloudConfigsAksPoolMachinesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAksPoolMachinesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAksPoolMachinesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAksPoolMachinesUIDDeleteNoContent creates a V1CloudConfigsAksPoolMachinesUIDDeleteNoContent with default headers values
func NewV1CloudConfigsAksPoolMachinesUIDDeleteNoContent() *V1CloudConfigsAksPoolMachinesUIDDeleteNoContent {
	return &V1CloudConfigsAksPoolMachinesUIDDeleteNoContent{}
}

/*
V1CloudConfigsAksPoolMachinesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsAksPoolMachinesUIDDeleteNoContent struct {
}

func (o *V1CloudConfigsAksPoolMachinesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/aks/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsAksPoolMachinesUidDeleteNoContent ", 204)
}

func (o *V1CloudConfigsAksPoolMachinesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}