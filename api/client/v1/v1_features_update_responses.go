// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1FeaturesUpdateReader is a Reader for the V1FeaturesUpdate structure.
type V1FeaturesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1FeaturesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1FeaturesUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1FeaturesUpdateNoContent creates a V1FeaturesUpdateNoContent with default headers values
func NewV1FeaturesUpdateNoContent() *V1FeaturesUpdateNoContent {
	return &V1FeaturesUpdateNoContent{}
}

/*
V1FeaturesUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1FeaturesUpdateNoContent struct {
}

func (o *V1FeaturesUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/features/{uid}][%d] v1FeaturesUpdateNoContent ", 204)
}

func (o *V1FeaturesUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}