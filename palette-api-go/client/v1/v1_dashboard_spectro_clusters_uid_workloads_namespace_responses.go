// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1DashboardSpectroClustersUIDWorkloadsNamespaceReader is a Reader for the V1DashboardSpectroClustersUIDWorkloadsNamespace structure.
type V1DashboardSpectroClustersUIDWorkloadsNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardSpectroClustersUIDWorkloadsNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardSpectroClustersUIDWorkloadsNamespaceOK creates a V1DashboardSpectroClustersUIDWorkloadsNamespaceOK with default headers values
func NewV1DashboardSpectroClustersUIDWorkloadsNamespaceOK() *V1DashboardSpectroClustersUIDWorkloadsNamespaceOK {
	return &V1DashboardSpectroClustersUIDWorkloadsNamespaceOK{}
}

/*V1DashboardSpectroClustersUIDWorkloadsNamespaceOK handles this case with default header values.

An array of cluster workload namespaces
*/
type V1DashboardSpectroClustersUIDWorkloadsNamespaceOK struct {
	Payload *models.V1ClusterWorkloadNamespaces
}

func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/spectroclusters/{uid}/workloads/namespace][%d] v1DashboardSpectroClustersUidWorkloadsNamespaceOK  %+v", 200, o.Payload)
}

func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceOK) GetPayload() *models.V1ClusterWorkloadNamespaces {
	return o.Payload
}

func (o *V1DashboardSpectroClustersUIDWorkloadsNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterWorkloadNamespaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
