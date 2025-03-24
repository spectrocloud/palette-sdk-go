// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDTokenKubeConfigUpdateReader is a Reader for the V1SpectroClustersUIDTokenKubeConfigUpdate structure.
type V1SpectroClustersUIDTokenKubeConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDTokenKubeConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigUpdateNoContent creates a V1SpectroClustersUIDTokenKubeConfigUpdateNoContent with default headers values
func NewV1SpectroClustersUIDTokenKubeConfigUpdateNoContent() *V1SpectroClustersUIDTokenKubeConfigUpdateNoContent {
	return &V1SpectroClustersUIDTokenKubeConfigUpdateNoContent{}
}

/*V1SpectroClustersUIDTokenKubeConfigUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDTokenKubeConfigUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDTokenKubeConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/assets/adminTokenKubeconfig][%d] v1SpectroClustersUidTokenKubeConfigUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDTokenKubeConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
