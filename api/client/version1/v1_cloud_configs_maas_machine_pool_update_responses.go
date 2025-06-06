// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsMaasMachinePoolUpdateReader is a Reader for the V1CloudConfigsMaasMachinePoolUpdate structure.
type V1CloudConfigsMaasMachinePoolUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsMaasMachinePoolUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsMaasMachinePoolUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsMaasMachinePoolUpdateNoContent creates a V1CloudConfigsMaasMachinePoolUpdateNoContent with default headers values
func NewV1CloudConfigsMaasMachinePoolUpdateNoContent() *V1CloudConfigsMaasMachinePoolUpdateNoContent {
	return &V1CloudConfigsMaasMachinePoolUpdateNoContent{}
}

/*
V1CloudConfigsMaasMachinePoolUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsMaasMachinePoolUpdateNoContent struct {
}

func (o *V1CloudConfigsMaasMachinePoolUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/maas/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsMaasMachinePoolUpdateNoContent ", 204)
}

func (o *V1CloudConfigsMaasMachinePoolUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
