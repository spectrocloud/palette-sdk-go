// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AwsCloudConfigsUIDHybridConfigReader is a Reader for the V1AwsCloudConfigsUIDHybridConfig structure.
type V1AwsCloudConfigsUIDHybridConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsCloudConfigsUIDHybridConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AwsCloudConfigsUIDHybridConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsCloudConfigsUIDHybridConfigNoContent creates a V1AwsCloudConfigsUIDHybridConfigNoContent with default headers values
func NewV1AwsCloudConfigsUIDHybridConfigNoContent() *V1AwsCloudConfigsUIDHybridConfigNoContent {
	return &V1AwsCloudConfigsUIDHybridConfigNoContent{}
}

/*
V1AwsCloudConfigsUIDHybridConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1AwsCloudConfigsUIDHybridConfigNoContent struct {
}

func (o *V1AwsCloudConfigsUIDHybridConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/aws/{configUid}/clusterConfig/hybridConfig][%d] v1AwsCloudConfigsUidHybridConfigNoContent ", 204)
}

func (o *V1AwsCloudConfigsUIDHybridConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
