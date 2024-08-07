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

// V1SpectroClustersVsphereRateReader is a Reader for the V1SpectroClustersVsphereRate structure.
type V1SpectroClustersVsphereRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersVsphereRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersVsphereRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersVsphereRateOK creates a V1SpectroClustersVsphereRateOK with default headers values
func NewV1SpectroClustersVsphereRateOK() *V1SpectroClustersVsphereRateOK {
	return &V1SpectroClustersVsphereRateOK{}
}

/*V1SpectroClustersVsphereRateOK handles this case with default header values.

Vsphere Cluster estimated rate response
*/
type V1SpectroClustersVsphereRateOK struct {
	Payload *models.V1SpectroClusterRate
}

func (o *V1SpectroClustersVsphereRateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/vsphere/rate][%d] v1SpectroClustersVsphereRateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersVsphereRateOK) GetPayload() *models.V1SpectroClusterRate {
	return o.Payload
}

func (o *V1SpectroClustersVsphereRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
