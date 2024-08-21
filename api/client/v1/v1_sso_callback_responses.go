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

// V1SsoCallbackReader is a Reader for the V1SsoCallback structure.
type V1SsoCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SsoCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SsoCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SsoCallbackOK creates a V1SsoCallbackOK with default headers values
func NewV1SsoCallbackOK() *V1SsoCallbackOK {
	return &V1SsoCallbackOK{}
}

/*
V1SsoCallbackOK handles this case with default header values.

OK
*/
type V1SsoCallbackOK struct {
	Payload *models.V1UserToken
}

func (o *V1SsoCallbackOK) Error() string {
	return fmt.Sprintf("[GET /v1/auth/sso/{ssoApp}/callback][%d] v1SsoCallbackOK  %+v", 200, o.Payload)
}

func (o *V1SsoCallbackOK) GetPayload() *models.V1UserToken {
	return o.Payload
}

func (o *V1SsoCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}