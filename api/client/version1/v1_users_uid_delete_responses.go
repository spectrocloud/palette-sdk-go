// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersUIDDeleteReader is a Reader for the V1UsersUIDDelete structure.
type V1UsersUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersUIDDeleteNoContent creates a V1UsersUIDDeleteNoContent with default headers values
func NewV1UsersUIDDeleteNoContent() *V1UsersUIDDeleteNoContent {
	return &V1UsersUIDDeleteNoContent{}
}

/*
V1UsersUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1UsersUIDDeleteNoContent struct {
}

func (o *V1UsersUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/{uid}][%d] v1UsersUidDeleteNoContent ", 204)
}

func (o *V1UsersUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
