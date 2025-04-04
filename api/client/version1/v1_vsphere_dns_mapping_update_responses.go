// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1VsphereDNSMappingUpdateReader is a Reader for the V1VsphereDNSMappingUpdate structure.
type V1VsphereDNSMappingUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VsphereDNSMappingUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1VsphereDNSMappingUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VsphereDNSMappingUpdateNoContent creates a V1VsphereDNSMappingUpdateNoContent with default headers values
func NewV1VsphereDNSMappingUpdateNoContent() *V1VsphereDNSMappingUpdateNoContent {
	return &V1VsphereDNSMappingUpdateNoContent{}
}

/*
V1VsphereDNSMappingUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1VsphereDNSMappingUpdateNoContent struct {
}

func (o *V1VsphereDNSMappingUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/users/assets/vsphere/dnsMappings/{uid}][%d] v1VsphereDnsMappingUpdateNoContent ", 204)
}

func (o *V1VsphereDNSMappingUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
