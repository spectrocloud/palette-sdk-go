// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDConfigRbacsUIDUpdateReader is a Reader for the V1SpectroClustersUIDConfigRbacsUIDUpdate structure.
type V1SpectroClustersUIDConfigRbacsUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDConfigRbacsUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDConfigRbacsUIDUpdateNoContent creates a V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent with default headers values
func NewV1SpectroClustersUIDConfigRbacsUIDUpdateNoContent() *V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent {
	return &V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent{}
}

/*
V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/config/rbacs/{rbacUid}][%d] v1SpectroClustersUidConfigRbacsUidUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDConfigRbacsUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}