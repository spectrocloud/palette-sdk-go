// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDVariablesPatchReader is a Reader for the V1SpectroClustersUIDVariablesPatch structure.
type V1SpectroClustersUIDVariablesPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDVariablesPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDVariablesPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDVariablesPatchNoContent creates a V1SpectroClustersUIDVariablesPatchNoContent with default headers values
func NewV1SpectroClustersUIDVariablesPatchNoContent() *V1SpectroClustersUIDVariablesPatchNoContent {
	return &V1SpectroClustersUIDVariablesPatchNoContent{}
}

/*
V1SpectroClustersUIDVariablesPatchNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDVariablesPatchNoContent struct {
}

func (o *V1SpectroClustersUIDVariablesPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/spectroclusters/{uid}/variables][%d] v1SpectroClustersUidVariablesPatchNoContent ", 204)
}

func (o *V1SpectroClustersUIDVariablesPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
