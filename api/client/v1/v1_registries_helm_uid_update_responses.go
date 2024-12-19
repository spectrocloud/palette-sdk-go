// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1RegistriesHelmUIDUpdateReader is a Reader for the V1RegistriesHelmUIDUpdate structure.
type V1RegistriesHelmUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RegistriesHelmUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1RegistriesHelmUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RegistriesHelmUIDUpdateNoContent creates a V1RegistriesHelmUIDUpdateNoContent with default headers values
func NewV1RegistriesHelmUIDUpdateNoContent() *V1RegistriesHelmUIDUpdateNoContent {
	return &V1RegistriesHelmUIDUpdateNoContent{}
}

/*V1RegistriesHelmUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1RegistriesHelmUIDUpdateNoContent struct {
}

func (o *V1RegistriesHelmUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/registries/helm/{uid}][%d] v1RegistriesHelmUidUpdateNoContent ", 204)
}

func (o *V1RegistriesHelmUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
