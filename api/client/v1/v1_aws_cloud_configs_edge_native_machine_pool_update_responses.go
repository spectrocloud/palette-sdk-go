// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AwsCloudConfigsEdgeNativeMachinePoolUpdateReader is a Reader for the V1AwsCloudConfigsEdgeNativeMachinePoolUpdate structure.
type V1AwsCloudConfigsEdgeNativeMachinePoolUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent creates a V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent with default headers values
func NewV1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent() *V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent {
	return &V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent{}
}

/*V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent struct {
}

func (o *V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/aws/{configUid}/edge-native/machinePools/{machinePoolName}][%d] v1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent ", 204)
}

func (o *V1AwsCloudConfigsEdgeNativeMachinePoolUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
