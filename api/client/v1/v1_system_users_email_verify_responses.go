// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SystemUsersEmailVerifyReader is a Reader for the V1SystemUsersEmailVerify structure.
type V1SystemUsersEmailVerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SystemUsersEmailVerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SystemUsersEmailVerifyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SystemUsersEmailVerifyNoContent creates a V1SystemUsersEmailVerifyNoContent with default headers values
func NewV1SystemUsersEmailVerifyNoContent() *V1SystemUsersEmailVerifyNoContent {
	return &V1SystemUsersEmailVerifyNoContent{}
}

/*
V1SystemUsersEmailVerifyNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SystemUsersEmailVerifyNoContent struct {
}

func (o *V1SystemUsersEmailVerifyNoContent) Error() string {
	return fmt.Sprintf("[GET /v1/system/users/email/{emailToken}/verify][%d] v1SystemUsersEmailVerifyNoContent ", 204)
}

func (o *V1SystemUsersEmailVerifyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
