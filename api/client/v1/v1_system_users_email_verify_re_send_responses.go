// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SystemUsersEmailVerifyReSendReader is a Reader for the V1SystemUsersEmailVerifyReSend structure.
type V1SystemUsersEmailVerifyReSendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SystemUsersEmailVerifyReSendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SystemUsersEmailVerifyReSendNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SystemUsersEmailVerifyReSendNoContent creates a V1SystemUsersEmailVerifyReSendNoContent with default headers values
func NewV1SystemUsersEmailVerifyReSendNoContent() *V1SystemUsersEmailVerifyReSendNoContent {
	return &V1SystemUsersEmailVerifyReSendNoContent{}
}

/*
V1SystemUsersEmailVerifyReSendNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SystemUsersEmailVerifyReSendNoContent struct {
}

func (o *V1SystemUsersEmailVerifyReSendNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/system/users/email/verify/resend][%d] v1SystemUsersEmailVerifyReSendNoContent ", 204)
}

func (o *V1SystemUsersEmailVerifyReSendNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
