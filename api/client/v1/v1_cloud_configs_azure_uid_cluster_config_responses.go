// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudConfigsAzureUIDClusterConfigReader is a Reader for the V1CloudConfigsAzureUIDClusterConfig structure.
type V1CloudConfigsAzureUIDClusterConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAzureUIDClusterConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudConfigsAzureUIDClusterConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAzureUIDClusterConfigNoContent creates a V1CloudConfigsAzureUIDClusterConfigNoContent with default headers values
func NewV1CloudConfigsAzureUIDClusterConfigNoContent() *V1CloudConfigsAzureUIDClusterConfigNoContent {
	return &V1CloudConfigsAzureUIDClusterConfigNoContent{}
}

/*V1CloudConfigsAzureUIDClusterConfigNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudConfigsAzureUIDClusterConfigNoContent struct {
}

func (o *V1CloudConfigsAzureUIDClusterConfigNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudconfigs/azure/{configUid}/clusterConfig][%d] v1CloudConfigsAzureUidClusterConfigNoContent ", 204)
}

func (o *V1CloudConfigsAzureUIDClusterConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
