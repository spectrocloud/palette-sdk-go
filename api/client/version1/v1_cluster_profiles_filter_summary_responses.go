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

// V1ClusterProfilesFilterSummaryReader is a Reader for the V1ClusterProfilesFilterSummary structure.
type V1ClusterProfilesFilterSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesFilterSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterProfilesFilterSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesFilterSummaryOK creates a V1ClusterProfilesFilterSummaryOK with default headers values
func NewV1ClusterProfilesFilterSummaryOK() *V1ClusterProfilesFilterSummaryOK {
	return &V1ClusterProfilesFilterSummaryOK{}
}

/*
V1ClusterProfilesFilterSummaryOK handles this case with default header values.

An array of cluster profiles summary items
*/
type V1ClusterProfilesFilterSummaryOK struct {
	Payload *models.V1ClusterProfilesSummary
}

func (o *V1ClusterProfilesFilterSummaryOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/clusterprofiles][%d] v1ClusterProfilesFilterSummaryOK  %+v", 200, o.Payload)
}

func (o *V1ClusterProfilesFilterSummaryOK) GetPayload() *models.V1ClusterProfilesSummary {
	return o.Payload
}

func (o *V1ClusterProfilesFilterSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterProfilesSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
