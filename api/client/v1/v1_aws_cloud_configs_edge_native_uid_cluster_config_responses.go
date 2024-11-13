// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AwsCloudConfigsEdgeNativeUIDClusterConfigReader is a Reader for the V1AwsCloudConfigsEdgeNativeUIDClusterConfig structure.
type V1AwsCloudConfigsEdgeNativeUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsCloudConfigsEdgeNativeUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent creates a V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent with default headers values
func NewV1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent() *V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent {
	return &V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent{}
}

/*
V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent struct {
}

func (o *V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/aws/{configUid}/edge-native/machinePools/{machinePoolName}/clusterConfig][%d] v1AwsCloudConfigsEdgeNativeUidClusterConfigNoContent ", 204)
}

func (o *V1AwsCloudConfigsEdgeNativeUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}