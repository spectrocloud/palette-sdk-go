// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1RateConfigUpdateReader is a Reader for the V1RateConfigUpdate structure.
type V1RateConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RateConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1RateConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RateConfigUpdateNoContent creates a V1RateConfigUpdateNoContent with default headers values
func NewV1RateConfigUpdateNoContent() *V1RateConfigUpdateNoContent {
	return &V1RateConfigUpdateNoContent{}
}

/*
V1RateConfigUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1RateConfigUpdateNoContent struct {
}

func (o *V1RateConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/tenants/{tenantUid}/rateConfig][%d] v1RateConfigUpdateNoContent ", 204)
}

func (o *V1RateConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
