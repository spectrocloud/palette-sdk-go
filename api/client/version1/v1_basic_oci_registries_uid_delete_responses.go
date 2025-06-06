// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1BasicOciRegistriesUIDDeleteReader is a Reader for the V1BasicOciRegistriesUIDDelete structure.
type V1BasicOciRegistriesUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1BasicOciRegistriesUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1BasicOciRegistriesUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1BasicOciRegistriesUIDDeleteNoContent creates a V1BasicOciRegistriesUIDDeleteNoContent with default headers values
func NewV1BasicOciRegistriesUIDDeleteNoContent() *V1BasicOciRegistriesUIDDeleteNoContent {
	return &V1BasicOciRegistriesUIDDeleteNoContent{}
}

/*
V1BasicOciRegistriesUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1BasicOciRegistriesUIDDeleteNoContent struct {
}

func (o *V1BasicOciRegistriesUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/registries/oci/{uid}/basic][%d] v1BasicOciRegistriesUidDeleteNoContent ", 204)
}

func (o *V1BasicOciRegistriesUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
