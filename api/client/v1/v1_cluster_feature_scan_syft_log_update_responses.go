// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterFeatureScanSyftLogUpdateReader is a Reader for the V1ClusterFeatureScanSyftLogUpdate structure.
type V1ClusterFeatureScanSyftLogUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterFeatureScanSyftLogUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterFeatureScanSyftLogUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterFeatureScanSyftLogUpdateNoContent creates a V1ClusterFeatureScanSyftLogUpdateNoContent with default headers values
func NewV1ClusterFeatureScanSyftLogUpdateNoContent() *V1ClusterFeatureScanSyftLogUpdateNoContent {
	return &V1ClusterFeatureScanSyftLogUpdateNoContent{}
}

/*
V1ClusterFeatureScanSyftLogUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1ClusterFeatureScanSyftLogUpdateNoContent struct {
}

func (o *V1ClusterFeatureScanSyftLogUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/spectroclusters/{uid}/features/complianceScan/logs/drivers/syft][%d] v1ClusterFeatureScanSyftLogUpdateNoContent ", 204)
}

func (o *V1ClusterFeatureScanSyftLogUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
