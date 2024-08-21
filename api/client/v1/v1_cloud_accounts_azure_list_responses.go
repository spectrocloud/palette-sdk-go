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

// V1CloudAccountsAzureListReader is a Reader for the V1CloudAccountsAzureList structure.
type V1CloudAccountsAzureListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsAzureListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsAzureListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsAzureListOK creates a V1CloudAccountsAzureListOK with default headers values
func NewV1CloudAccountsAzureListOK() *V1CloudAccountsAzureListOK {
	return &V1CloudAccountsAzureListOK{}
}

/*
V1CloudAccountsAzureListOK handles this case with default header values.

An array of azure cloud account items
*/
type V1CloudAccountsAzureListOK struct {
	Payload *models.V1AzureAccounts
}

func (o *V1CloudAccountsAzureListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/azure][%d] v1CloudAccountsAzureListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsAzureListOK) GetPayload() *models.V1AzureAccounts {
	return o.Payload
}

func (o *V1CloudAccountsAzureListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AzureAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}