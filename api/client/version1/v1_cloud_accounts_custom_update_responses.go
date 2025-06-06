// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudAccountsCustomUpdateReader is a Reader for the V1CloudAccountsCustomUpdate structure.
type V1CloudAccountsCustomUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsCustomUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudAccountsCustomUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsCustomUpdateNoContent creates a V1CloudAccountsCustomUpdateNoContent with default headers values
func NewV1CloudAccountsCustomUpdateNoContent() *V1CloudAccountsCustomUpdateNoContent {
	return &V1CloudAccountsCustomUpdateNoContent{}
}

/*
V1CloudAccountsCustomUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudAccountsCustomUpdateNoContent struct {
}

func (o *V1CloudAccountsCustomUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/cloudaccounts/cloudTypes/{cloudType}/{uid}][%d] v1CloudAccountsCustomUpdateNoContent ", 204)
}

func (o *V1CloudAccountsCustomUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
