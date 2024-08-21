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

// V1APIKeysListReader is a Reader for the V1APIKeysList structure.
type V1APIKeysListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1APIKeysListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1APIKeysListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1APIKeysListOK creates a V1APIKeysListOK with default headers values
func NewV1APIKeysListOK() *V1APIKeysListOK {
	return &V1APIKeysListOK{}
}

/*
V1APIKeysListOK handles this case with default header values.

Retrieves a list of API keys
*/
type V1APIKeysListOK struct {
	Payload *models.V1APIKeys
}

func (o *V1APIKeysListOK) Error() string {
	return fmt.Sprintf("[GET /v1/apiKeys][%d] v1ApiKeysListOK  %+v", 200, o.Payload)
}

func (o *V1APIKeysListOK) GetPayload() *models.V1APIKeys {
	return o.Payload
}

func (o *V1APIKeysListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1APIKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}