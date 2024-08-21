// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1EdgeTokensUIDStateReader is a Reader for the V1EdgeTokensUIDState structure.
type V1EdgeTokensUIDStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeTokensUIDStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1EdgeTokensUIDStateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeTokensUIDStateNoContent creates a V1EdgeTokensUIDStateNoContent with default headers values
func NewV1EdgeTokensUIDStateNoContent() *V1EdgeTokensUIDStateNoContent {
	return &V1EdgeTokensUIDStateNoContent{}
}

/*
V1EdgeTokensUIDStateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1EdgeTokensUIDStateNoContent struct {
}

func (o *V1EdgeTokensUIDStateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/edgehosts/tokens/{uid}/state][%d] v1EdgeTokensUidStateNoContent ", 204)
}

func (o *V1EdgeTokensUIDStateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}