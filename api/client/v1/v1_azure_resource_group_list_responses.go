// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1AzureResourceGroupListReader is a Reader for the V1AzureResourceGroupList structure.
type V1AzureResourceGroupListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AzureResourceGroupListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AzureResourceGroupListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AzureResourceGroupListOK creates a V1AzureResourceGroupListOK with default headers values
func NewV1AzureResourceGroupListOK() *V1AzureResourceGroupListOK {
	return &V1AzureResourceGroupListOK{}
}

/*
V1AzureResourceGroupListOK handles this case with default header values.

(empty)
*/
type V1AzureResourceGroupListOK struct {
	Payload *models.V1AzureResourceGroupList
}

func (o *V1AzureResourceGroupListOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/azure/regions/{region}/subscriptions/{subscriptionId}/resourceGroups][%d] v1AzureResourceGroupListOK  %+v", 200, o.Payload)
}

func (o *V1AzureResourceGroupListOK) GetPayload() *models.V1AzureResourceGroupList {
	return o.Payload
}

func (o *V1AzureResourceGroupListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AzureResourceGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}