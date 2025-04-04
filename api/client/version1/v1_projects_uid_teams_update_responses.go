// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ProjectsUIDTeamsUpdateReader is a Reader for the V1ProjectsUIDTeamsUpdate structure.
type V1ProjectsUIDTeamsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ProjectsUIDTeamsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ProjectsUIDTeamsUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ProjectsUIDTeamsUpdateNoContent creates a V1ProjectsUIDTeamsUpdateNoContent with default headers values
func NewV1ProjectsUIDTeamsUpdateNoContent() *V1ProjectsUIDTeamsUpdateNoContent {
	return &V1ProjectsUIDTeamsUpdateNoContent{}
}

/*
V1ProjectsUIDTeamsUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1ProjectsUIDTeamsUpdateNoContent struct {
}

func (o *V1ProjectsUIDTeamsUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/{uid}/teams][%d] v1ProjectsUidTeamsUpdateNoContent ", 204)
}

func (o *V1ProjectsUIDTeamsUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
