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

// V1AwsAccountSecretCredentialsReader is a Reader for the V1AwsAccountSecretCredentials structure.
type V1AwsAccountSecretCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsAccountSecretCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsAccountSecretCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsAccountSecretCredentialsOK creates a V1AwsAccountSecretCredentialsOK with default headers values
func NewV1AwsAccountSecretCredentialsOK() *V1AwsAccountSecretCredentialsOK {
	return &V1AwsAccountSecretCredentialsOK{}
}

/*
V1AwsAccountSecretCredentialsOK handles this case with default header values.

(empty)
*/
type V1AwsAccountSecretCredentialsOK struct {
	Payload *models.V1AwsAccountCredentials
}

func (o *V1AwsAccountSecretCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /v1/clouds/aws/account/secret/credentials][%d] v1AwsAccountSecretCredentialsOK  %+v", 200, o.Payload)
}

func (o *V1AwsAccountSecretCredentialsOK) GetPayload() *models.V1AwsAccountCredentials {
	return o.Payload
}

func (o *V1AwsAccountSecretCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsAccountCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
