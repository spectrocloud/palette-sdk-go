// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsGkeMachinePoolDeleteReader is a Reader for the V1CloudConfigsGkeMachinePoolDelete structure.
type V1CloudConfigsGkeMachinePoolDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsGkeMachinePoolDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsGkeMachinePoolDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsGkeMachinePoolDeleteNoContent creates a V1CloudConfigsGkeMachinePoolDeleteNoContent with default headers values
func NewV1CloudConfigsGkeMachinePoolDeleteNoContent() *V1CloudConfigsGkeMachinePoolDeleteNoContent {
	return &V1CloudConfigsGkeMachinePoolDeleteNoContent{}
}

/*
V1CloudConfigsGkeMachinePoolDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsGkeMachinePoolDeleteNoContent struct {
}

func (o *V1CloudConfigsGkeMachinePoolDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/gke/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsGkeMachinePoolDeleteNoContent ", 204)
}

func (o *V1CloudConfigsGkeMachinePoolDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
