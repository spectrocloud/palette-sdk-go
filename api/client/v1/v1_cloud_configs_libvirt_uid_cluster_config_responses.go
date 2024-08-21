// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsLibvirtUIDClusterConfigReader is a Reader for the V1CloudConfigsLibvirtUIDClusterConfig structure.
type V1CloudConfigsLibvirtUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsLibvirtUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsLibvirtUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsLibvirtUIDClusterConfigNoContent creates a V1CloudConfigsLibvirtUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsLibvirtUIDClusterConfigNoContent() *V1CloudConfigsLibvirtUIDClusterConfigNoContent {
	return &V1CloudConfigsLibvirtUIDClusterConfigNoContent{}
}

/*
V1CloudConfigsLibvirtUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsLibvirtUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsLibvirtUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/libvirt/{configUid}/clusterConfig][%d] v1CloudConfigsLibvirtUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsLibvirtUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}