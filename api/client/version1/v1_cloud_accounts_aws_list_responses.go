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

// V1CloudAccountsAwsListReader is a Reader for the V1CloudAccountsAwsList structure.
type V1CloudAccountsAwsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsAwsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsAwsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsAwsListOK creates a V1CloudAccountsAwsListOK with default headers values
func NewV1CloudAccountsAwsListOK() *V1CloudAccountsAwsListOK {
	return &V1CloudAccountsAwsListOK{}
}

/*
V1CloudAccountsAwsListOK handles this case with default header values.

An array of cloud account items
*/
type V1CloudAccountsAwsListOK struct {
	Payload *models.V1AwsAccounts
}

func (o *V1CloudAccountsAwsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/aws][%d] v1CloudAccountsAwsListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsAwsListOK) GetPayload() *models.V1AwsAccounts {
	return o.Payload
}

func (o *V1CloudAccountsAwsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
