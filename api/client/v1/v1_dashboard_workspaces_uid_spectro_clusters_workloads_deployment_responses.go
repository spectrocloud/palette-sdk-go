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

// V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentReader is a Reader for the V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeployment structure.
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK creates a V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK with default headers values
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK {
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK{}
}

/*V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK handles this case with default header values.

An array of clusters workload deployments
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK struct {
	Payload *models.V1WorkspaceClustersWorkloadDeployments
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/workspaces/{uid}/spectroclusters/workloads/deployment][%d] v1DashboardWorkspacesUidSpectroClustersWorkloadsDeploymentOK  %+v", 200, o.Payload)
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK) GetPayload() *models.V1WorkspaceClustersWorkloadDeployments {
	return o.Payload
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1WorkspaceClustersWorkloadDeployments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
