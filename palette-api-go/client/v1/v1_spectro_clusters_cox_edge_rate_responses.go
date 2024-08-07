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

// V1SpectroClustersCoxEdgeRateReader is a Reader for the V1SpectroClustersCoxEdgeRate structure.
type V1SpectroClustersCoxEdgeRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersCoxEdgeRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersCoxEdgeRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersCoxEdgeRateOK creates a V1SpectroClustersCoxEdgeRateOK with default headers values
func NewV1SpectroClustersCoxEdgeRateOK() *V1SpectroClustersCoxEdgeRateOK {
	return &V1SpectroClustersCoxEdgeRateOK{}
}

/*V1SpectroClustersCoxEdgeRateOK handles this case with default header values.

Azure Cluster estimated rate response
*/
type V1SpectroClustersCoxEdgeRateOK struct {
	Payload *models.V1SpectroClusterRate
}

func (o *V1SpectroClustersCoxEdgeRateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/coxedge/rate][%d] v1SpectroClustersCoxEdgeRateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersCoxEdgeRateOK) GetPayload() *models.V1SpectroClusterRate {
	return o.Payload
}

func (o *V1SpectroClustersCoxEdgeRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
