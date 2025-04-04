// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudAccountsVsphereUpdateReader is a Reader for the V1CloudAccountsVsphereUpdate structure.
type V1CloudAccountsVsphereUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsVsphereUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudAccountsVsphereUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsVsphereUpdateNoContent creates a V1CloudAccountsVsphereUpdateNoContent with default headers values
func NewV1CloudAccountsVsphereUpdateNoContent() *V1CloudAccountsVsphereUpdateNoContent {
	return &V1CloudAccountsVsphereUpdateNoContent{}
}

/*
V1CloudAccountsVsphereUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudAccountsVsphereUpdateNoContent struct {
}

func (o *V1CloudAccountsVsphereUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudaccounts/vsphere/{uid}][%d] v1CloudAccountsVsphereUpdateNoContent ", 204)
}

func (o *V1CloudAccountsVsphereUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
