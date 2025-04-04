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

// V1DashboardSpectroClustersUIDWorkloadsJobReader is a Reader for the V1DashboardSpectroClustersUIDWorkloadsJob structure.
type V1DashboardSpectroClustersUIDWorkloadsJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardSpectroClustersUIDWorkloadsJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardSpectroClustersUIDWorkloadsJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsJobOK creates a V1DashboardSpectroClustersUIDWorkloadsJobOK with default headers values
func NewV1DashboardSpectroClustersUIDWorkloadsJobOK() *V1DashboardSpectroClustersUIDWorkloadsJobOK {
	return &V1DashboardSpectroClustersUIDWorkloadsJobOK{}
}

/*
V1DashboardSpectroClustersUIDWorkloadsJobOK handles this case with default header values.

An array of cluster workload jobs
*/
type V1DashboardSpectroClustersUIDWorkloadsJobOK struct {
	Payload *models.V1ClusterWorkloadJobs
}

func (o *V1DashboardSpectroClustersUIDWorkloadsJobOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/spectroclusters/{uid}/workloads/job][%d] v1DashboardSpectroClustersUidWorkloadsJobOK  %+v", 200, o.Payload)
}

func (o *V1DashboardSpectroClustersUIDWorkloadsJobOK) GetPayload() *models.V1ClusterWorkloadJobs {
	return o.Payload
}

func (o *V1DashboardSpectroClustersUIDWorkloadsJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterWorkloadJobs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
