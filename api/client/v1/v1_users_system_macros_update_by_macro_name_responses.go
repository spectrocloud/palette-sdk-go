// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersSystemMacrosUpdateByMacroNameReader is a Reader for the V1UsersSystemMacrosUpdateByMacroName structure.
type V1UsersSystemMacrosUpdateByMacroNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersSystemMacrosUpdateByMacroNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersSystemMacrosUpdateByMacroNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersSystemMacrosUpdateByMacroNameNoContent creates a V1UsersSystemMacrosUpdateByMacroNameNoContent with default headers values
func NewV1UsersSystemMacrosUpdateByMacroNameNoContent() *V1UsersSystemMacrosUpdateByMacroNameNoContent {
	return &V1UsersSystemMacrosUpdateByMacroNameNoContent{}
}

/*
V1UsersSystemMacrosUpdateByMacroNameNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1UsersSystemMacrosUpdateByMacroNameNoContent struct {
}

func (o *V1UsersSystemMacrosUpdateByMacroNameNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/users/system/macros][%d] v1UsersSystemMacrosUpdateByMacroNameNoContent ", 204)
}

func (o *V1UsersSystemMacrosUpdateByMacroNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
