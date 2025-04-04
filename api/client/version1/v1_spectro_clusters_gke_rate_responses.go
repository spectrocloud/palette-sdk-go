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

// V1SpectroClustersGkeRateReader is a Reader for the V1SpectroClustersGkeRate structure.
type V1SpectroClustersGkeRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersGkeRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersGkeRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersGkeRateOK creates a V1SpectroClustersGkeRateOK with default headers values
func NewV1SpectroClustersGkeRateOK() *V1SpectroClustersGkeRateOK {
	return &V1SpectroClustersGkeRateOK{}
}

/*
V1SpectroClustersGkeRateOK handles this case with default header values.

Gke Cluster estimated rate response
*/
type V1SpectroClustersGkeRateOK struct {
	Payload *models.V1SpectroClusterRate
}

func (o *V1SpectroClustersGkeRateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/gke/rate][%d] v1SpectroClustersGkeRateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersGkeRateOK) GetPayload() *models.V1SpectroClusterRate {
	return o.Payload
}

func (o *V1SpectroClustersGkeRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
