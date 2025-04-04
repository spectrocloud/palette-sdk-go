// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TeamsUIDUpdateReader is a Reader for the V1TeamsUIDUpdate structure.
type V1TeamsUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TeamsUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsUIDUpdateNoContent creates a V1TeamsUIDUpdateNoContent with default headers values
func NewV1TeamsUIDUpdateNoContent() *V1TeamsUIDUpdateNoContent {
	return &V1TeamsUIDUpdateNoContent{}
}

/*
V1TeamsUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TeamsUIDUpdateNoContent struct {
}

func (o *V1TeamsUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/teams/{uid}][%d] v1TeamsUidUpdateNoContent ", 204)
}

func (o *V1TeamsUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
