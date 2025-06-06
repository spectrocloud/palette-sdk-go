// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1ClusterGroupUIDVirtualClustersSummaryReader is a Reader for the V1ClusterGroupUIDVirtualClustersSummary structure.
type V1ClusterGroupUIDVirtualClustersSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterGroupUIDVirtualClustersSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterGroupUIDVirtualClustersSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterGroupUIDVirtualClustersSummaryOK creates a V1ClusterGroupUIDVirtualClustersSummaryOK with default headers values
func NewV1ClusterGroupUIDVirtualClustersSummaryOK() *V1ClusterGroupUIDVirtualClustersSummaryOK {
	return &V1ClusterGroupUIDVirtualClustersSummaryOK{}
}

/*
V1ClusterGroupUIDVirtualClustersSummaryOK handles this case with default header values.

An array of cluster summary items
*/
type V1ClusterGroupUIDVirtualClustersSummaryOK struct {
	Payload *models.V1SpectroClustersSummary
}

func (o *V1ClusterGroupUIDVirtualClustersSummaryOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/clustergroups/{uid}/virtualClusters][%d] v1ClusterGroupUidVirtualClustersSummaryOK  %+v", 200, o.Payload)
}

func (o *V1ClusterGroupUIDVirtualClustersSummaryOK) GetPayload() *models.V1SpectroClustersSummary {
	return o.Payload
}

func (o *V1ClusterGroupUIDVirtualClustersSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClustersSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
