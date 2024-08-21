// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TeamsUIDDeleteReader is a Reader for the V1TeamsUIDDelete structure.
type V1TeamsUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TeamsUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsUIDDeleteNoContent creates a V1TeamsUIDDeleteNoContent with default headers values
func NewV1TeamsUIDDeleteNoContent() *V1TeamsUIDDeleteNoContent {
	return &V1TeamsUIDDeleteNoContent{}
}

/*
V1TeamsUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1TeamsUIDDeleteNoContent struct {
}

func (o *V1TeamsUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/teams/{uid}][%d] v1TeamsUidDeleteNoContent ", 204)
}

func (o *V1TeamsUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}