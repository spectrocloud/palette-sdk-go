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

// V1PasswordsBlockListUpdateReader is a Reader for the V1PasswordsBlockListUpdate structure.
type V1PasswordsBlockListUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1PasswordsBlockListUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1PasswordsBlockListUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PasswordsBlockListUpdateNoContent creates a V1PasswordsBlockListUpdateNoContent with default headers values
func NewV1PasswordsBlockListUpdateNoContent() *V1PasswordsBlockListUpdateNoContent {
	return &V1PasswordsBlockListUpdateNoContent{}
}

/*
V1PasswordsBlockListUpdateNoContent handles this case with default header values.

(empty)
*/
type V1PasswordsBlockListUpdateNoContent struct {
	Payload models.V1Updated
}

func (o *V1PasswordsBlockListUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/system/passwords/blocklist][%d] v1PasswordsBlockListUpdateNoContent  %+v", 204, o.Payload)
}

func (o *V1PasswordsBlockListUpdateNoContent) GetPayload() models.V1Updated {
	return o.Payload
}

func (o *V1PasswordsBlockListUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
