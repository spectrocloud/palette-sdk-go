// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEdgeMachinePoolDeleteReader is a Reader for the V1CloudConfigsEdgeMachinePoolDelete structure.
type V1CloudConfigsEdgeMachinePoolDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEdgeMachinePoolDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEdgeMachinePoolDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEdgeMachinePoolDeleteNoContent creates a V1CloudConfigsEdgeMachinePoolDeleteNoContent with default headers values
func NewV1CloudConfigsEdgeMachinePoolDeleteNoContent() *V1CloudConfigsEdgeMachinePoolDeleteNoContent {
	return &V1CloudConfigsEdgeMachinePoolDeleteNoContent{}
}

/*
V1CloudConfigsEdgeMachinePoolDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsEdgeMachinePoolDeleteNoContent struct {
}

func (o *V1CloudConfigsEdgeMachinePoolDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/edge/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsEdgeMachinePoolDeleteNoContent ", 204)
}

func (o *V1CloudConfigsEdgeMachinePoolDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
