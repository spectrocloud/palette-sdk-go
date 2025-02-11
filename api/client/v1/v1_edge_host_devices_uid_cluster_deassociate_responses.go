// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1EdgeHostDevicesUIDClusterDeassociateReader is a Reader for the V1EdgeHostDevicesUIDClusterDeassociate structure.
type V1EdgeHostDevicesUIDClusterDeassociateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeHostDevicesUIDClusterDeassociateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1EdgeHostDevicesUIDClusterDeassociateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeHostDevicesUIDClusterDeassociateNoContent creates a V1EdgeHostDevicesUIDClusterDeassociateNoContent with default headers values
func NewV1EdgeHostDevicesUIDClusterDeassociateNoContent() *V1EdgeHostDevicesUIDClusterDeassociateNoContent {
	return &V1EdgeHostDevicesUIDClusterDeassociateNoContent{}
}

/*
V1EdgeHostDevicesUIDClusterDeassociateNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1EdgeHostDevicesUIDClusterDeassociateNoContent struct {
}

func (o *V1EdgeHostDevicesUIDClusterDeassociateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/edgehosts/{uid}/cluster/associate][%d] v1EdgeHostDevicesUidClusterDeassociateNoContent ", 204)
}

func (o *V1EdgeHostDevicesUIDClusterDeassociateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
