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

// V1DashboardPcgsSearchSummaryReader is a Reader for the V1DashboardPcgsSearchSummary structure.
type V1DashboardPcgsSearchSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardPcgsSearchSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardPcgsSearchSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardPcgsSearchSummaryOK creates a V1DashboardPcgsSearchSummaryOK with default headers values
func NewV1DashboardPcgsSearchSummaryOK() *V1DashboardPcgsSearchSummaryOK {
	return &V1DashboardPcgsSearchSummaryOK{}
}

/*
V1DashboardPcgsSearchSummaryOK handles this case with default header values.

An array of cluster summary items
*/
type V1DashboardPcgsSearchSummaryOK struct {
	Payload *models.V1PcgsSummary
}

func (o *V1DashboardPcgsSearchSummaryOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/pcgs/search][%d] v1DashboardPcgsSearchSummaryOK  %+v", 200, o.Payload)
}

func (o *V1DashboardPcgsSearchSummaryOK) GetPayload() *models.V1PcgsSummary {
	return o.Payload
}

func (o *V1DashboardPcgsSearchSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PcgsSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}