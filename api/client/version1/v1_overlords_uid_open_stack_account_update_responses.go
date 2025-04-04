// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1OverlordsUIDOpenStackAccountUpdateReader is a Reader for the V1OverlordsUIDOpenStackAccountUpdate structure.
type V1OverlordsUIDOpenStackAccountUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsUIDOpenStackAccountUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1OverlordsUIDOpenStackAccountUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsUIDOpenStackAccountUpdateNoContent creates a V1OverlordsUIDOpenStackAccountUpdateNoContent with default headers values
func NewV1OverlordsUIDOpenStackAccountUpdateNoContent() *V1OverlordsUIDOpenStackAccountUpdateNoContent {
	return &V1OverlordsUIDOpenStackAccountUpdateNoContent{}
}

/*
V1OverlordsUIDOpenStackAccountUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1OverlordsUIDOpenStackAccountUpdateNoContent struct {
}

func (o *V1OverlordsUIDOpenStackAccountUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/overlords/openstack/{uid}/account][%d] v1OverlordsUidOpenStackAccountUpdateNoContent ", 204)
}

func (o *V1OverlordsUIDOpenStackAccountUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
