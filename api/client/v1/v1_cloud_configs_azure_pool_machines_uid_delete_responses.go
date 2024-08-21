// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAzurePoolMachinesUIDDeleteReader is a Reader for the V1CloudConfigsAzurePoolMachinesUIDDelete structure.
type V1CloudConfigsAzurePoolMachinesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAzurePoolMachinesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAzurePoolMachinesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAzurePoolMachinesUIDDeleteNoContent creates a V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent with default headers values
func NewV1CloudConfigsAzurePoolMachinesUIDDeleteNoContent() *V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent {
	return &V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent{}
}

/*
V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent struct {
}

func (o *V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/azure/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsAzurePoolMachinesUidDeleteNoContent ", 204)
}

func (o *V1CloudConfigsAzurePoolMachinesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}