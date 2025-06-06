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

// V1SpectroClustersResourcesUsageSummaryReader is a Reader for the V1SpectroClustersResourcesUsageSummary structure.
type V1SpectroClustersResourcesUsageSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersResourcesUsageSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersResourcesUsageSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersResourcesUsageSummaryOK creates a V1SpectroClustersResourcesUsageSummaryOK with default headers values
func NewV1SpectroClustersResourcesUsageSummaryOK() *V1SpectroClustersResourcesUsageSummaryOK {
	return &V1SpectroClustersResourcesUsageSummaryOK{}
}

/*
V1SpectroClustersResourcesUsageSummaryOK handles this case with default header values.

An array of resources usage summary items
*/
type V1SpectroClustersResourcesUsageSummaryOK struct {
	Payload *models.V1ResourcesUsageSummary
}

func (o *V1SpectroClustersResourcesUsageSummaryOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/spectroclusters/resources/usage][%d] v1SpectroClustersResourcesUsageSummaryOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersResourcesUsageSummaryOK) GetPayload() *models.V1ResourcesUsageSummary {
	return o.Payload
}

func (o *V1SpectroClustersResourcesUsageSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ResourcesUsageSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
