// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDPacksStatusPatchReader is a Reader for the V1SpectroClustersUIDPacksStatusPatch structure.
type V1SpectroClustersUIDPacksStatusPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDPacksStatusPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDPacksStatusPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDPacksStatusPatchNoContent creates a V1SpectroClustersUIDPacksStatusPatchNoContent with default headers values
func NewV1SpectroClustersUIDPacksStatusPatchNoContent() *V1SpectroClustersUIDPacksStatusPatchNoContent {
	return &V1SpectroClustersUIDPacksStatusPatchNoContent{}
}

/*
V1SpectroClustersUIDPacksStatusPatchNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDPacksStatusPatchNoContent struct {
}

func (o *V1SpectroClustersUIDPacksStatusPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/spectroclusters/{uid}/packs/status][%d] v1SpectroClustersUidPacksStatusPatchNoContent ", 204)
}

func (o *V1SpectroClustersUIDPacksStatusPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
