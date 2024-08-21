// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudAccountsMaasPatchReader is a Reader for the V1CloudAccountsMaasPatch structure.
type V1CloudAccountsMaasPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsMaasPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudAccountsMaasPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsMaasPatchNoContent creates a V1CloudAccountsMaasPatchNoContent with default headers values
func NewV1CloudAccountsMaasPatchNoContent() *V1CloudAccountsMaasPatchNoContent {
	return &V1CloudAccountsMaasPatchNoContent{}
}

/*
V1CloudAccountsMaasPatchNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1CloudAccountsMaasPatchNoContent struct {
}

func (o *V1CloudAccountsMaasPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/cloudaccounts/maas/{uid}][%d] v1CloudAccountsMaasPatchNoContent ", 204)
}

func (o *V1CloudAccountsMaasPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}