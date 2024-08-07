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

// V1SpectroClustersClusterRbacReader is a Reader for the V1SpectroClustersClusterRbac structure.
type V1SpectroClustersClusterRbacReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersClusterRbacReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersClusterRbacOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersClusterRbacOK creates a V1SpectroClustersClusterRbacOK with default headers values
func NewV1SpectroClustersClusterRbacOK() *V1SpectroClustersClusterRbacOK {
	return &V1SpectroClustersClusterRbacOK{}
}

/*V1SpectroClustersClusterRbacOK handles this case with default header values.

(empty)
*/
type V1SpectroClustersClusterRbacOK struct {
	Payload *models.V1ClusterRbacs
}

func (o *V1SpectroClustersClusterRbacOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/clusterrbac][%d] v1SpectroClustersClusterRbacOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersClusterRbacOK) GetPayload() *models.V1ClusterRbacs {
	return o.Payload
}

func (o *V1SpectroClustersClusterRbacOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterRbacs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
