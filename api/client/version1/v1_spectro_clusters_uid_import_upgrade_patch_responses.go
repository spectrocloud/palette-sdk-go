// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDImportUpgradePatchReader is a Reader for the V1SpectroClustersUIDImportUpgradePatch structure.
type V1SpectroClustersUIDImportUpgradePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDImportUpgradePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDImportUpgradePatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDImportUpgradePatchNoContent creates a V1SpectroClustersUIDImportUpgradePatchNoContent with default headers values
func NewV1SpectroClustersUIDImportUpgradePatchNoContent() *V1SpectroClustersUIDImportUpgradePatchNoContent {
	return &V1SpectroClustersUIDImportUpgradePatchNoContent{}
}

/*
V1SpectroClustersUIDImportUpgradePatchNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDImportUpgradePatchNoContent struct {
}

func (o *V1SpectroClustersUIDImportUpgradePatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/spectroclusters/{uid}/import/upgrade][%d] v1SpectroClustersUidImportUpgradePatchNoContent ", 204)
}

func (o *V1SpectroClustersUIDImportUpgradePatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
