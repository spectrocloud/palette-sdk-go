// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAksMachinePoolUpdateReader is a Reader for the V1CloudConfigsAksMachinePoolUpdate structure.
type V1CloudConfigsAksMachinePoolUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAksMachinePoolUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAksMachinePoolUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAksMachinePoolUpdateNoContent creates a V1CloudConfigsAksMachinePoolUpdateNoContent with default headers values
func NewV1CloudConfigsAksMachinePoolUpdateNoContent() *V1CloudConfigsAksMachinePoolUpdateNoContent {
	return &V1CloudConfigsAksMachinePoolUpdateNoContent{}
}

/*
V1CloudConfigsAksMachinePoolUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsAksMachinePoolUpdateNoContent struct {
}

func (o *V1CloudConfigsAksMachinePoolUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/aks/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsAksMachinePoolUpdateNoContent ", 204)
}

func (o *V1CloudConfigsAksMachinePoolUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
