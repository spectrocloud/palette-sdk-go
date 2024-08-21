// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsMaasUIDClusterConfigReader is a Reader for the V1CloudConfigsMaasUIDClusterConfig structure.
type V1CloudConfigsMaasUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsMaasUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsMaasUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsMaasUIDClusterConfigNoContent creates a V1CloudConfigsMaasUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsMaasUIDClusterConfigNoContent() *V1CloudConfigsMaasUIDClusterConfigNoContent {
	return &V1CloudConfigsMaasUIDClusterConfigNoContent{}
}

/*
V1CloudConfigsMaasUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsMaasUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsMaasUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/maas/{configUid}/clusterConfig][%d] v1CloudConfigsMaasUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsMaasUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}