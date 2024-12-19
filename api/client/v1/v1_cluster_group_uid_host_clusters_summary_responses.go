// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1ClusterGroupUIDHostClustersSummaryReader is a Reader for the V1ClusterGroupUIDHostClustersSummary structure.
type V1ClusterGroupUIDHostClustersSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterGroupUIDHostClustersSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterGroupUIDHostClustersSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterGroupUIDHostClustersSummaryOK creates a V1ClusterGroupUIDHostClustersSummaryOK with default headers values
func NewV1ClusterGroupUIDHostClustersSummaryOK() *V1ClusterGroupUIDHostClustersSummaryOK {
	return &V1ClusterGroupUIDHostClustersSummaryOK{}
}

/*
V1ClusterGroupUIDHostClustersSummaryOK handles this case with default header values.

An array of cluster summary items
*/
type V1ClusterGroupUIDHostClustersSummaryOK struct {
	Payload *models.V1SpectroClustersSummary
}

func (o *V1ClusterGroupUIDHostClustersSummaryOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/clustergroups/{uid}/hostClusters][%d] v1ClusterGroupUidHostClustersSummaryOK  %+v", 200, o.Payload)
}

func (o *V1ClusterGroupUIDHostClustersSummaryOK) GetPayload() *models.V1SpectroClustersSummary {
	return o.Payload
}

func (o *V1ClusterGroupUIDHostClustersSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClustersSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
