// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsVspherePoolMachinesUIDDeleteReader is a Reader for the V1CloudConfigsVspherePoolMachinesUIDDelete structure.
type V1CloudConfigsVspherePoolMachinesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsVspherePoolMachinesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDDeleteNoContent creates a V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent with default headers values
func NewV1CloudConfigsVspherePoolMachinesUIDDeleteNoContent() *V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent {
	return &V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent{}
}

/*
V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent struct {
}

func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/vsphere/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsVspherePoolMachinesUidDeleteNoContent ", 204)
}

func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
