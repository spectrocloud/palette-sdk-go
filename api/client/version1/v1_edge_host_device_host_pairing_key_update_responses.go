// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1EdgeHostDeviceHostPairingKeyUpdateReader is a Reader for the V1EdgeHostDeviceHostPairingKeyUpdate structure.
type V1EdgeHostDeviceHostPairingKeyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeHostDeviceHostPairingKeyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1EdgeHostDeviceHostPairingKeyUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeHostDeviceHostPairingKeyUpdateNoContent creates a V1EdgeHostDeviceHostPairingKeyUpdateNoContent with default headers values
func NewV1EdgeHostDeviceHostPairingKeyUpdateNoContent() *V1EdgeHostDeviceHostPairingKeyUpdateNoContent {
	return &V1EdgeHostDeviceHostPairingKeyUpdateNoContent{}
}

/*
V1EdgeHostDeviceHostPairingKeyUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1EdgeHostDeviceHostPairingKeyUpdateNoContent struct {
}

func (o *V1EdgeHostDeviceHostPairingKeyUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/edgehosts/{uid}/hostPairingKey][%d] v1EdgeHostDeviceHostPairingKeyUpdateNoContent ", 204)
}

func (o *V1EdgeHostDeviceHostPairingKeyUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
