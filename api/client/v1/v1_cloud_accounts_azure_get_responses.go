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

// V1CloudAccountsAzureGetReader is a Reader for the V1CloudAccountsAzureGet structure.
type V1CloudAccountsAzureGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsAzureGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsAzureGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsAzureGetOK creates a V1CloudAccountsAzureGetOK with default headers values
func NewV1CloudAccountsAzureGetOK() *V1CloudAccountsAzureGetOK {
	return &V1CloudAccountsAzureGetOK{}
}

/*
V1CloudAccountsAzureGetOK handles this case with default header values.

OK
*/
type V1CloudAccountsAzureGetOK struct {
	Payload *models.V1AzureAccount
}

func (o *V1CloudAccountsAzureGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/azure/{uid}][%d] v1CloudAccountsAzureGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsAzureGetOK) GetPayload() *models.V1AzureAccount {
	return o.Payload
}

func (o *V1CloudAccountsAzureGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AzureAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
