// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDManifestUpdateReader is a Reader for the V1SpectroClustersUIDManifestUpdate structure.
type V1SpectroClustersUIDManifestUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDManifestUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDManifestUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDManifestUpdateNoContent creates a V1SpectroClustersUIDManifestUpdateNoContent with default headers values
func NewV1SpectroClustersUIDManifestUpdateNoContent() *V1SpectroClustersUIDManifestUpdateNoContent {
	return &V1SpectroClustersUIDManifestUpdateNoContent{}
}

/*
V1SpectroClustersUIDManifestUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDManifestUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDManifestUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/assets/manifest][%d] v1SpectroClustersUidManifestUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDManifestUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
