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

// V1SpectroClustersSearchFilterSummaryReader is a Reader for the V1SpectroClustersSearchFilterSummary structure.
type V1SpectroClustersSearchFilterSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersSearchFilterSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersSearchFilterSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersSearchFilterSummaryOK creates a V1SpectroClustersSearchFilterSummaryOK with default headers values
func NewV1SpectroClustersSearchFilterSummaryOK() *V1SpectroClustersSearchFilterSummaryOK {
	return &V1SpectroClustersSearchFilterSummaryOK{}
}

/*
V1SpectroClustersSearchFilterSummaryOK handles this case with default header values.

An array of cluster summary items
*/
type V1SpectroClustersSearchFilterSummaryOK struct {
	Payload *models.V1SpectroClustersSummary
}

func (o *V1SpectroClustersSearchFilterSummaryOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/spectroclusters/search][%d] v1SpectroClustersSearchFilterSummaryOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersSearchFilterSummaryOK) GetPayload() *models.V1SpectroClustersSummary {
	return o.Payload
}

func (o *V1SpectroClustersSearchFilterSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClustersSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
