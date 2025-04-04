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

// V1SamlCallbackReader is a Reader for the V1SamlCallback structure.
type V1SamlCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SamlCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SamlCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SamlCallbackOK creates a V1SamlCallbackOK with default headers values
func NewV1SamlCallbackOK() *V1SamlCallbackOK {
	return &V1SamlCallbackOK{}
}

/*
V1SamlCallbackOK handles this case with default header values.

OK
*/
type V1SamlCallbackOK struct {
	Payload *models.V1UserToken
}

func (o *V1SamlCallbackOK) Error() string {
	return fmt.Sprintf("[POST /v1/auth/org/{org}/saml/callback][%d] v1SamlCallbackOK  %+v", 200, o.Payload)
}

func (o *V1SamlCallbackOK) GetPayload() *models.V1UserToken {
	return o.Payload
}

func (o *V1SamlCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
