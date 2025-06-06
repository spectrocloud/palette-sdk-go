// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsEksUIDClusterConfigReader is a Reader for the V1CloudConfigsEksUIDClusterConfig structure.
type V1CloudConfigsEksUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsEksUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsEksUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsEksUIDClusterConfigNoContent creates a V1CloudConfigsEksUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsEksUIDClusterConfigNoContent() *V1CloudConfigsEksUIDClusterConfigNoContent {
	return &V1CloudConfigsEksUIDClusterConfigNoContent{}
}

/*
V1CloudConfigsEksUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsEksUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsEksUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/eks/{configUid}/clusterConfig][%d] v1CloudConfigsEksUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsEksUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
