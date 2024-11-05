// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1BasicOciRegistriesUIDUpdateReader is a Reader for the V1BasicOciRegistriesUIDUpdate structure.
type V1BasicOciRegistriesUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1BasicOciRegistriesUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1BasicOciRegistriesUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1BasicOciRegistriesUIDUpdateNoContent creates a V1BasicOciRegistriesUIDUpdateNoContent with default headers values
func NewV1BasicOciRegistriesUIDUpdateNoContent() *V1BasicOciRegistriesUIDUpdateNoContent {
	return &V1BasicOciRegistriesUIDUpdateNoContent{}
}

/*
V1BasicOciRegistriesUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1BasicOciRegistriesUIDUpdateNoContent struct {
}

func (o *V1BasicOciRegistriesUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/registries/oci/{uid}/basic][%d] v1BasicOciRegistriesUidUpdateNoContent ", 204)
}

func (o *V1BasicOciRegistriesUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
