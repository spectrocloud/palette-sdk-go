// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1WorkspacesUIDMetaUpdateReader is a Reader for the V1WorkspacesUIDMetaUpdate structure.
type V1WorkspacesUIDMetaUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1WorkspacesUIDMetaUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1WorkspacesUIDMetaUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1WorkspacesUIDMetaUpdateNoContent creates a V1WorkspacesUIDMetaUpdateNoContent with default headers values
func NewV1WorkspacesUIDMetaUpdateNoContent() *V1WorkspacesUIDMetaUpdateNoContent {
	return &V1WorkspacesUIDMetaUpdateNoContent{}
}

/*
V1WorkspacesUIDMetaUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1WorkspacesUIDMetaUpdateNoContent struct {
}

func (o *V1WorkspacesUIDMetaUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/workspaces/{uid}/meta][%d] v1WorkspacesUidMetaUpdateNoContent ", 204)
}

func (o *V1WorkspacesUIDMetaUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
