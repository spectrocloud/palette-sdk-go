// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUpdateStatusServicesReader is a Reader for the V1SpectroClustersUpdateStatusServices structure.
type V1SpectroClustersUpdateStatusServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUpdateStatusServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUpdateStatusServicesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUpdateStatusServicesNoContent creates a V1SpectroClustersUpdateStatusServicesNoContent with default headers values
func NewV1SpectroClustersUpdateStatusServicesNoContent() *V1SpectroClustersUpdateStatusServicesNoContent {
	return &V1SpectroClustersUpdateStatusServicesNoContent{}
}

/*
V1SpectroClustersUpdateStatusServicesNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUpdateStatusServicesNoContent struct {
}

func (o *V1SpectroClustersUpdateStatusServicesNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/status/services][%d] v1SpectroClustersUpdateStatusServicesNoContent ", 204)
}

func (o *V1SpectroClustersUpdateStatusServicesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}