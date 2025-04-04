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

// V1SpectroClustersOpenStackRateReader is a Reader for the V1SpectroClustersOpenStackRate structure.
type V1SpectroClustersOpenStackRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersOpenStackRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersOpenStackRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersOpenStackRateOK creates a V1SpectroClustersOpenStackRateOK with default headers values
func NewV1SpectroClustersOpenStackRateOK() *V1SpectroClustersOpenStackRateOK {
	return &V1SpectroClustersOpenStackRateOK{}
}

/*
V1SpectroClustersOpenStackRateOK handles this case with default header values.

Openstack Cluster estimated rate response
*/
type V1SpectroClustersOpenStackRateOK struct {
	Payload *models.V1SpectroClusterRate
}

func (o *V1SpectroClustersOpenStackRateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/openstack/rate][%d] v1SpectroClustersOpenStackRateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersOpenStackRateOK) GetPayload() *models.V1SpectroClusterRate {
	return o.Payload
}

func (o *V1SpectroClustersOpenStackRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
