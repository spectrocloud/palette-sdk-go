// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEksMachinePoolUpdateReader is a Reader for the V1CloudConfigsEksMachinePoolUpdate structure.
type V1CloudConfigsEksMachinePoolUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEksMachinePoolUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEksMachinePoolUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEksMachinePoolUpdateNoContent creates a V1CloudConfigsEksMachinePoolUpdateNoContent with default headers values
func NewV1CloudConfigsEksMachinePoolUpdateNoContent() *V1CloudConfigsEksMachinePoolUpdateNoContent {
	return &V1CloudConfigsEksMachinePoolUpdateNoContent{}
}

/*
V1CloudConfigsEksMachinePoolUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsEksMachinePoolUpdateNoContent struct {
}

func (o *V1CloudConfigsEksMachinePoolUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/eks/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsEksMachinePoolUpdateNoContent ", 204)
}

func (o *V1CloudConfigsEksMachinePoolUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
