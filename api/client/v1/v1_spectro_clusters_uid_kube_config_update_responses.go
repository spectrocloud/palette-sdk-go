// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDKubeConfigUpdateReader is a Reader for the V1SpectroClustersUIDKubeConfigUpdate structure.
type V1SpectroClustersUIDKubeConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDKubeConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDKubeConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDKubeConfigUpdateNoContent creates a V1SpectroClustersUIDKubeConfigUpdateNoContent with default headers values
func NewV1SpectroClustersUIDKubeConfigUpdateNoContent() *V1SpectroClustersUIDKubeConfigUpdateNoContent {
	return &V1SpectroClustersUIDKubeConfigUpdateNoContent{}
}

/*
V1SpectroClustersUIDKubeConfigUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDKubeConfigUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDKubeConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/assets/kubeconfig][%d] v1SpectroClustersUidKubeConfigUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDKubeConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
