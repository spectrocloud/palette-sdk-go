// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDKubeConfigClientUpdateReader is a Reader for the V1SpectroClustersUIDKubeConfigClientUpdate structure.
type V1SpectroClustersUIDKubeConfigClientUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDKubeConfigClientUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDKubeConfigClientUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDKubeConfigClientUpdateNoContent creates a V1SpectroClustersUIDKubeConfigClientUpdateNoContent with default headers values
func NewV1SpectroClustersUIDKubeConfigClientUpdateNoContent() *V1SpectroClustersUIDKubeConfigClientUpdateNoContent {
	return &V1SpectroClustersUIDKubeConfigClientUpdateNoContent{}
}

/*
V1SpectroClustersUIDKubeConfigClientUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDKubeConfigClientUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDKubeConfigClientUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/assets/kubeconfigclient][%d] v1SpectroClustersUidKubeConfigClientUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDKubeConfigClientUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
