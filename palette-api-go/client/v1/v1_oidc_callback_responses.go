// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1OidcCallbackReader is a Reader for the V1OidcCallback structure.
type V1OidcCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OidcCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OidcCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OidcCallbackOK creates a V1OidcCallbackOK with default headers values
func NewV1OidcCallbackOK() *V1OidcCallbackOK {
	return &V1OidcCallbackOK{}
}

/*V1OidcCallbackOK handles this case with default header values.

OK
*/
type V1OidcCallbackOK struct {
	Payload *models.V1UserToken
}

func (o *V1OidcCallbackOK) Error() string {
	return fmt.Sprintf("[GET /v1/auth/org/{org}/oidc/callback][%d] v1OidcCallbackOK  %+v", 200, o.Payload)
}

func (o *V1OidcCallbackOK) GetPayload() *models.V1UserToken {
	return o.Payload
}

func (o *V1OidcCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
