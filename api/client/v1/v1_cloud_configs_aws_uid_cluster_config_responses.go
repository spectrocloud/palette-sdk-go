// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAwsUIDClusterConfigReader is a Reader for the V1CloudConfigsAwsUIDClusterConfig structure.
type V1CloudConfigsAwsUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAwsUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAwsUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAwsUIDClusterConfigNoContent creates a V1CloudConfigsAwsUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsAwsUIDClusterConfigNoContent() *V1CloudConfigsAwsUIDClusterConfigNoContent {
	return &V1CloudConfigsAwsUIDClusterConfigNoContent{}
}

/*V1CloudConfigsAwsUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsAwsUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsAwsUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/aws/{configUid}/clusterConfig][%d] v1CloudConfigsAwsUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsAwsUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
