// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersVMDeleteReader is a Reader for the V1SpectroClustersVMDelete structure.
type V1SpectroClustersVMDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersVMDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersVMDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersVMDeleteNoContent creates a V1SpectroClustersVMDeleteNoContent with default headers values
func NewV1SpectroClustersVMDeleteNoContent() *V1SpectroClustersVMDeleteNoContent {
	return &V1SpectroClustersVMDeleteNoContent{}
}

/*
V1SpectroClustersVMDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1SpectroClustersVMDeleteNoContent struct {
}

func (o *V1SpectroClustersVMDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/spectroclusters/{uid}/vms/{vmName}][%d] v1SpectroClustersVmDeleteNoContent ", 204)
}

func (o *V1SpectroClustersVMDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
