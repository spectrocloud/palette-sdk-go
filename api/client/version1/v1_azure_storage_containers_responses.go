// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1AzureStorageContainersReader is a Reader for the V1AzureStorageContainers structure.
type V1AzureStorageContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AzureStorageContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AzureStorageContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AzureStorageContainersOK creates a V1AzureStorageContainersOK with default headers values
func NewV1AzureStorageContainersOK() *V1AzureStorageContainersOK {
	return &V1AzureStorageContainersOK{}
}

/*
V1AzureStorageContainersOK handles this case with default header values.

(empty)
*/
type V1AzureStorageContainersOK struct {
	Payload *models.V1AzureStorageContainers
}

func (o *V1AzureStorageContainersOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/azure/resourceGroups/{resourceGroup}/storageAccounts/{storageAccountName}/containers][%d] v1AzureStorageContainersOK  %+v", 200, o.Payload)
}

func (o *V1AzureStorageContainersOK) GetPayload() *models.V1AzureStorageContainers {
	return o.Payload
}

func (o *V1AzureStorageContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AzureStorageContainers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
