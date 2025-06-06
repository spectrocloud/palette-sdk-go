// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEdgeNativeUIDClusterConfigReader is a Reader for the V1CloudConfigsEdgeNativeUIDClusterConfig structure.
type V1CloudConfigsEdgeNativeUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEdgeNativeUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEdgeNativeUIDClusterConfigNoContent creates a V1CloudConfigsEdgeNativeUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsEdgeNativeUIDClusterConfigNoContent() *V1CloudConfigsEdgeNativeUIDClusterConfigNoContent {
	return &V1CloudConfigsEdgeNativeUIDClusterConfigNoContent{}
}

/*
V1CloudConfigsEdgeNativeUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsEdgeNativeUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsEdgeNativeUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/edge-native/{configUid}/clusterConfig][%d] v1CloudConfigsEdgeNativeUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsEdgeNativeUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
