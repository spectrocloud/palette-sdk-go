// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1APIKeysUIDUpdateReader is a Reader for the V1APIKeysUIDUpdate structure.
type V1APIKeysUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1APIKeysUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1APIKeysUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1APIKeysUIDUpdateNoContent creates a V1APIKeysUIDUpdateNoContent with default headers values
func NewV1APIKeysUIDUpdateNoContent() *V1APIKeysUIDUpdateNoContent {
	return &V1APIKeysUIDUpdateNoContent{}
}

/*
V1APIKeysUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1APIKeysUIDUpdateNoContent struct {
}

func (o *V1APIKeysUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/apiKeys/{uid}][%d] v1ApiKeysUidUpdateNoContent ", 204)
}

func (o *V1APIKeysUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
