// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersAssetsLocationDeleteReader is a Reader for the V1UsersAssetsLocationDelete structure.
type V1UsersAssetsLocationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersAssetsLocationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersAssetsLocationDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersAssetsLocationDeleteNoContent creates a V1UsersAssetsLocationDeleteNoContent with default headers values
func NewV1UsersAssetsLocationDeleteNoContent() *V1UsersAssetsLocationDeleteNoContent {
	return &V1UsersAssetsLocationDeleteNoContent{}
}

/*
V1UsersAssetsLocationDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1UsersAssetsLocationDeleteNoContent struct {
}

func (o *V1UsersAssetsLocationDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/assets/locations/{uid}][%d] v1UsersAssetsLocationDeleteNoContent ", 204)
}

func (o *V1UsersAssetsLocationDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
