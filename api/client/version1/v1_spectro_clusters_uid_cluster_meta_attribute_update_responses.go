// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDClusterMetaAttributeUpdateReader is a Reader for the V1SpectroClustersUIDClusterMetaAttributeUpdate structure.
type V1SpectroClustersUIDClusterMetaAttributeUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDClusterMetaAttributeUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDClusterMetaAttributeUpdateNoContent creates a V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent with default headers values
func NewV1SpectroClustersUIDClusterMetaAttributeUpdateNoContent() *V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent {
	return &V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent{}
}

/*
V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/spectroclusters/{uid}/clusterConfig/clusterMetaAttribute][%d] v1SpectroClustersUidClusterMetaAttributeUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDClusterMetaAttributeUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
