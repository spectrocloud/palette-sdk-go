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

// V1SpectroClustersGcpRateReader is a Reader for the V1SpectroClustersGcpRate structure.
type V1SpectroClustersGcpRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersGcpRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersGcpRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersGcpRateOK creates a V1SpectroClustersGcpRateOK with default headers values
func NewV1SpectroClustersGcpRateOK() *V1SpectroClustersGcpRateOK {
	return &V1SpectroClustersGcpRateOK{}
}

/*
V1SpectroClustersGcpRateOK handles this case with default header values.

Gcp Cluster estimated rate response
*/
type V1SpectroClustersGcpRateOK struct {
	Payload *models.V1SpectroClusterRate
}

func (o *V1SpectroClustersGcpRateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/gcp/rate][%d] v1SpectroClustersGcpRateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersGcpRateOK) GetPayload() *models.V1SpectroClusterRate {
	return o.Payload
}

func (o *V1SpectroClustersGcpRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}