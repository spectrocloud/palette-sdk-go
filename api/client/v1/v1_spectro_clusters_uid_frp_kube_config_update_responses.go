// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDFrpKubeConfigUpdateReader is a Reader for the V1SpectroClustersUIDFrpKubeConfigUpdate structure.
type V1SpectroClustersUIDFrpKubeConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDFrpKubeConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDFrpKubeConfigUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDFrpKubeConfigUpdateNoContent creates a V1SpectroClustersUIDFrpKubeConfigUpdateNoContent with default headers values
func NewV1SpectroClustersUIDFrpKubeConfigUpdateNoContent() *V1SpectroClustersUIDFrpKubeConfigUpdateNoContent {
	return &V1SpectroClustersUIDFrpKubeConfigUpdateNoContent{}
}

/*V1SpectroClustersUIDFrpKubeConfigUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersUIDFrpKubeConfigUpdateNoContent struct {
}

func (o *V1SpectroClustersUIDFrpKubeConfigUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/assets/frpKubeconfig][%d] v1SpectroClustersUidFrpKubeConfigUpdateNoContent ", 204)
}

func (o *V1SpectroClustersUIDFrpKubeConfigUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
