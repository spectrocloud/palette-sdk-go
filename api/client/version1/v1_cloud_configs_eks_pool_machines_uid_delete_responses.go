// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEksPoolMachinesUIDDeleteReader is a Reader for the V1CloudConfigsEksPoolMachinesUIDDelete structure.
type V1CloudConfigsEksPoolMachinesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEksPoolMachinesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDDeleteNoContent creates a V1CloudConfigsEksPoolMachinesUIDDeleteNoContent with default headers values
func NewV1CloudConfigsEksPoolMachinesUIDDeleteNoContent() *V1CloudConfigsEksPoolMachinesUIDDeleteNoContent {
	return &V1CloudConfigsEksPoolMachinesUIDDeleteNoContent{}
}

/*
V1CloudConfigsEksPoolMachinesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsEksPoolMachinesUIDDeleteNoContent struct {
}

func (o *V1CloudConfigsEksPoolMachinesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/eks/{configUid}/machinePools/{machinePoolName}/machines/{machineUid}][%d] v1CloudConfigsEksPoolMachinesUidDeleteNoContent ", 204)
}

func (o *V1CloudConfigsEksPoolMachinesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
