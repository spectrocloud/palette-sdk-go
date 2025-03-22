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

// V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobReader is a Reader for the V1DashboardWorkspacesUIDSpectroClustersWorkloadsJob structure.
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK creates a V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK with default headers values
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK {
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK{}
}

/*V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK handles this case with default header values.

An array of clusters workload jobs
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK struct {
	Payload *models.V1WorkspaceClustersWorkloadJobs
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/workspaces/{uid}/spectroclusters/workloads/job][%d] v1DashboardWorkspacesUidSpectroClustersWorkloadsJobOK  %+v", 200, o.Payload)
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK) GetPayload() *models.V1WorkspaceClustersWorkloadJobs {
	return o.Payload
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1WorkspaceClustersWorkloadJobs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
