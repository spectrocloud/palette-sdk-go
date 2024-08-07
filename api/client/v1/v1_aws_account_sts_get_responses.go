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

// V1AwsAccountStsGetReader is a Reader for the V1AwsAccountStsGet structure.
type V1AwsAccountStsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsAccountStsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsAccountStsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsAccountStsGetOK creates a V1AwsAccountStsGetOK with default headers values
func NewV1AwsAccountStsGetOK() *V1AwsAccountStsGetOK {
	return &V1AwsAccountStsGetOK{}
}

/*
V1AwsAccountStsGetOK handles this case with default header values.

(empty)
*/
type V1AwsAccountStsGetOK struct {
	Payload *models.V1AwsAccountSts
}

func (o *V1AwsAccountStsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/aws/account/sts][%d] v1AwsAccountStsGetOK  %+v", 200, o.Payload)
}

func (o *V1AwsAccountStsGetOK) GetPayload() *models.V1AwsAccountSts {
	return o.Payload
}

func (o *V1AwsAccountStsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsAccountSts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
