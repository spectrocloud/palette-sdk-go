// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsVsphereUIDClusterConfigReader is a Reader for the V1CloudConfigsVsphereUIDClusterConfig structure.
type V1CloudConfigsVsphereUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsVsphereUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsVsphereUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsVsphereUIDClusterConfigNoContent creates a V1CloudConfigsVsphereUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsVsphereUIDClusterConfigNoContent() *V1CloudConfigsVsphereUIDClusterConfigNoContent {
	return &V1CloudConfigsVsphereUIDClusterConfigNoContent{}
}

/*
V1CloudConfigsVsphereUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsVsphereUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsVsphereUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/vsphere/{configUid}/clusterConfig][%d] v1CloudConfigsVsphereUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsVsphereUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
