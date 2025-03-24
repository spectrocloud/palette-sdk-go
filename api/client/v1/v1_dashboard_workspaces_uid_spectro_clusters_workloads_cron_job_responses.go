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

// V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobReader is a Reader for the V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJob structure.
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK creates a V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK with default headers values
func NewV1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK() *V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK {
	return &V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK{}
}

/*V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK handles this case with default header values.

An array of clusters workload cronjobs
*/
type V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK struct {
	Payload *models.V1WorkspaceClustersWorkloadCronJobs
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/workspaces/{uid}/spectroclusters/workloads/cronjob][%d] v1DashboardWorkspacesUidSpectroClustersWorkloadsCronJobOK  %+v", 200, o.Payload)
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK) GetPayload() *models.V1WorkspaceClustersWorkloadCronJobs {
	return o.Payload
}

func (o *V1DashboardWorkspacesUIDSpectroClustersWorkloadsCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1WorkspaceClustersWorkloadCronJobs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
