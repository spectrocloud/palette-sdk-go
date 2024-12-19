// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsTkeMachinePoolUpdateReader is a Reader for the V1CloudConfigsTkeMachinePoolUpdate structure.
type V1CloudConfigsTkeMachinePoolUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsTkeMachinePoolUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsTkeMachinePoolUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsTkeMachinePoolUpdateNoContent creates a V1CloudConfigsTkeMachinePoolUpdateNoContent with default headers values
func NewV1CloudConfigsTkeMachinePoolUpdateNoContent() *V1CloudConfigsTkeMachinePoolUpdateNoContent {
	return &V1CloudConfigsTkeMachinePoolUpdateNoContent{}
}

/*V1CloudConfigsTkeMachinePoolUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsTkeMachinePoolUpdateNoContent struct {
}

func (o *V1CloudConfigsTkeMachinePoolUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/tke/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsTkeMachinePoolUpdateNoContent ", 204)
}

func (o *V1CloudConfigsTkeMachinePoolUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
