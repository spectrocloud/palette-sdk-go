// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1OverlordsUIDMaasCloudConfigUpdateReader is a Reader for the V1OverlordsUIDMaasCloudConfigUpdate structure.
type V1OverlordsUIDMaasCloudConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsUIDMaasCloudConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1OverlordsUIDMaasCloudConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsUIDMaasCloudConfigUpdateNoContent creates a V1OverlordsUIDMaasCloudConfigUpdateNoContent with default headers values
func NewV1OverlordsUIDMaasCloudConfigUpdateNoContent() *V1OverlordsUIDMaasCloudConfigUpdateNoContent {
	return &V1OverlordsUIDMaasCloudConfigUpdateNoContent{}
}

/*
V1OverlordsUIDMaasCloudConfigUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1OverlordsUIDMaasCloudConfigUpdateNoContent struct {
}

func (o *V1OverlordsUIDMaasCloudConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/overlords/maas/{uid}/cloudconfig][%d] v1OverlordsUidMaasCloudConfigUpdateNoContent ", 204)
}

func (o *V1OverlordsUIDMaasCloudConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
