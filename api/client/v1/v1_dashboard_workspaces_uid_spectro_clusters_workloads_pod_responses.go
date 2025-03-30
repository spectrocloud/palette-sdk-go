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

// V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodReader is a Reader for the V1DashboardWorkspacesUIDSpectroClustersWorkloadsPod structure.
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK creates a V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK with default headers values
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK {
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK{}
}

/*
V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK handles this case with default header values.

An array of clusters workload pods
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK struct {
	Payload *models.V1WorkspaceClustersWorkloadPods
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/workspaces/{uid}/spectroclusters/workloads/pod][%d] v1DashboardWorkspacesUidSpectroClustersWorkloadsPodOK  %+v", 200, o.Payload)
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK) GetPayload() *models.V1WorkspaceClustersWorkloadPods {
	return o.Payload
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsPodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1WorkspaceClustersWorkloadPods)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
