// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1UsersAssetsLocationGcpUpdateReader is a Reader for the V1UsersAssetsLocationGcpUpdate structure.
type V1UsersAssetsLocationGcpUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersAssetsLocationGcpUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1UsersAssetsLocationGcpUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersAssetsLocationGcpUpdateNoContent creates a V1UsersAssetsLocationGcpUpdateNoContent with default headers values
func NewV1UsersAssetsLocationGcpUpdateNoContent() *V1UsersAssetsLocationGcpUpdateNoContent {
	return &V1UsersAssetsLocationGcpUpdateNoContent{}
}

/*
V1UsersAssetsLocationGcpUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1UsersAssetsLocationGcpUpdateNoContent struct {
}

func (o *V1UsersAssetsLocationGcpUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/users/assets/locations/gcp/{uid}][%d] v1UsersAssetsLocationGcpUpdateNoContent ", 204)
}

func (o *V1UsersAssetsLocationGcpUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
