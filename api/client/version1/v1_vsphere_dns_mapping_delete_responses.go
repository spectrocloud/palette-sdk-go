// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1VsphereDNSMappingDeleteReader is a Reader for the V1VsphereDNSMappingDelete structure.
type V1VsphereDNSMappingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VsphereDNSMappingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1VsphereDNSMappingDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VsphereDNSMappingDeleteNoContent creates a V1VsphereDNSMappingDeleteNoContent with default headers values
func NewV1VsphereDNSMappingDeleteNoContent() *V1VsphereDNSMappingDeleteNoContent {
	return &V1VsphereDNSMappingDeleteNoContent{}
}

/*
V1VsphereDNSMappingDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1VsphereDNSMappingDeleteNoContent struct {
}

func (o *V1VsphereDNSMappingDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/assets/vsphere/dnsMappings/{uid}][%d] v1VsphereDnsMappingDeleteNoContent ", 204)
}

func (o *V1VsphereDNSMappingDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
