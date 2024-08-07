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

// V1SpectroClustersEdgeRateReader is a Reader for the V1SpectroClustersEdgeRate structure.
type V1SpectroClustersEdgeRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersEdgeRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersEdgeRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersEdgeRateOK creates a V1SpectroClustersEdgeRateOK with default headers values
func NewV1SpectroClustersEdgeRateOK() *V1SpectroClustersEdgeRateOK {
	return &V1SpectroClustersEdgeRateOK{}
}

/*
V1SpectroClustersEdgeRateOK handles this case with default header values.

Edge Cluster estimated rate response
*/
type V1SpectroClustersEdgeRateOK struct {
	Payload *models.V1SpectroClusterRate
}

func (o *V1SpectroClustersEdgeRateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/edge/rate][%d] v1SpectroClustersEdgeRateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersEdgeRateOK) GetPayload() *models.V1SpectroClusterRate {
	return o.Payload
}

func (o *V1SpectroClustersEdgeRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
