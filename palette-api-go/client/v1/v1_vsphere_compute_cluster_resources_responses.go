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

// V1VsphereComputeClusterResourcesReader is a Reader for the V1VsphereComputeClusterResources structure.
type V1VsphereComputeClusterResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VsphereComputeClusterResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VsphereComputeClusterResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VsphereComputeClusterResourcesOK creates a V1VsphereComputeClusterResourcesOK with default headers values
func NewV1VsphereComputeClusterResourcesOK() *V1VsphereComputeClusterResourcesOK {
	return &V1VsphereComputeClusterResourcesOK{}
}

/*V1VsphereComputeClusterResourcesOK handles this case with default header values.

(empty)
*/
type V1VsphereComputeClusterResourcesOK struct {
	Payload *models.V1VsphereComputeClusterResources
}

func (o *V1VsphereComputeClusterResourcesOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/vsphere/datacenters/{uid}/computeclusters/{computecluster}][%d] v1VsphereComputeClusterResourcesOK  %+v", 200, o.Payload)
}

func (o *V1VsphereComputeClusterResourcesOK) GetPayload() *models.V1VsphereComputeClusterResources {
	return o.Payload
}

func (o *V1VsphereComputeClusterResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VsphereComputeClusterResources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
