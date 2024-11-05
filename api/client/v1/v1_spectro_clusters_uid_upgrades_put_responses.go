// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDUpgradesPutReader is a Reader for the V1SpectroClustersUIDUpgradesPut structure.
type V1SpectroClustersUIDUpgradesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDUpgradesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDUpgradesPutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDUpgradesPutNoContent creates a V1SpectroClustersUIDUpgradesPutNoContent with default headers values
func NewV1SpectroClustersUIDUpgradesPutNoContent() *V1SpectroClustersUIDUpgradesPutNoContent {
	return &V1SpectroClustersUIDUpgradesPutNoContent{}
}

/*V1SpectroClustersUIDUpgradesPutNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDUpgradesPutNoContent struct {
}

func (o *V1SpectroClustersUIDUpgradesPutNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/status/upgrades][%d] v1SpectroClustersUidUpgradesPutNoContent ", 204)
}

func (o *V1SpectroClustersUIDUpgradesPutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
