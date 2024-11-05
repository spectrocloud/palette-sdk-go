// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsTkeUIDClusterConfigReader is a Reader for the V1CloudConfigsTkeUIDClusterConfig structure.
type V1CloudConfigsTkeUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsTkeUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsTkeUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsTkeUIDClusterConfigNoContent creates a V1CloudConfigsTkeUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsTkeUIDClusterConfigNoContent() *V1CloudConfigsTkeUIDClusterConfigNoContent {
	return &V1CloudConfigsTkeUIDClusterConfigNoContent{}
}

/*
V1CloudConfigsTkeUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsTkeUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsTkeUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/tke/{configUid}/clusterConfig][%d] v1CloudConfigsTkeUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsTkeUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
