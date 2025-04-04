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

// V1DashboardSpectroClustersUIDWorkloadsCronJobReader is a Reader for the V1DashboardSpectroClustersUIDWorkloadsCronJob structure.
type V1DashboardSpectroClustersUIDWorkloadsCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardSpectroClustersUIDWorkloadsCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardSpectroClustersUIDWorkloadsCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsCronJobOK creates a V1DashboardSpectroClustersUIDWorkloadsCronJobOK with default headers values
func NewV1DashboardSpectroClustersUIDWorkloadsCronJobOK() *V1DashboardSpectroClustersUIDWorkloadsCronJobOK {
	return &V1DashboardSpectroClustersUIDWorkloadsCronJobOK{}
}

/*
V1DashboardSpectroClustersUIDWorkloadsCronJobOK handles this case with default header values.

An array of cluster workload cronjobs
*/
type V1DashboardSpectroClustersUIDWorkloadsCronJobOK struct {
	Payload *models.V1ClusterWorkloadCronJobs
}

func (o *V1DashboardSpectroClustersUIDWorkloadsCronJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/spectroclusters/{uid}/workloads/cronjob][%d] v1DashboardSpectroClustersUidWorkloadsCronJobOK  %+v", 200, o.Payload)
}

func (o *V1DashboardSpectroClustersUIDWorkloadsCronJobOK) GetPayload() *models.V1ClusterWorkloadCronJobs {
	return o.Payload
}

func (o *V1DashboardSpectroClustersUIDWorkloadsCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterWorkloadCronJobs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
