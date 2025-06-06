// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsMaasPoolMachinesUIDDeleteReader is a Reader for the V1CloudConfigsMaasPoolMachinesUIDDelete structure.
type V1CloudConfigsMaasPoolMachinesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsMaasPoolMachinesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsMaasPoolMachinesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsMaasPoolMachinesUIDDeleteNoContent creates a V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent with default headers values
func NewV1CloudConfigsMaasPoolMachinesUIDDeleteNoContent() *V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent {
	return &V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent{}
}

/*
V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent struct {
}

func (o *V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/maas/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsMaasPoolMachinesUidDeleteNoContent ", 204)
}

func (o *V1CloudConfigsMaasPoolMachinesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
