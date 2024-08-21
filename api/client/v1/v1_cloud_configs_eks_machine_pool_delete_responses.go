// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEksMachinePoolDeleteReader is a Reader for the V1CloudConfigsEksMachinePoolDelete structure.
type V1CloudConfigsEksMachinePoolDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEksMachinePoolDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEksMachinePoolDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEksMachinePoolDeleteNoContent creates a V1CloudConfigsEksMachinePoolDeleteNoContent with default headers values
func NewV1CloudConfigsEksMachinePoolDeleteNoContent() *V1CloudConfigsEksMachinePoolDeleteNoContent {
	return &V1CloudConfigsEksMachinePoolDeleteNoContent{}
}

/*
V1CloudConfigsEksMachinePoolDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CloudConfigsEksMachinePoolDeleteNoContent struct {
}

func (o *V1CloudConfigsEksMachinePoolDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/cloudconfigs/eks/{configUid}/machinePools/{machinePoolName}][%d] v1CloudConfigsEksMachinePoolDeleteNoContent ", 204)
}

func (o *V1CloudConfigsEksMachinePoolDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}