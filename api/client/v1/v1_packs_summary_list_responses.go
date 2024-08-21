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

// V1PacksSummaryListReader is a Reader for the V1PacksSummaryList structure.
type V1PacksSummaryListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1PacksSummaryListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PacksSummaryListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PacksSummaryListOK creates a V1PacksSummaryListOK with default headers values
func NewV1PacksSummaryListOK() *V1PacksSummaryListOK {
	return &V1PacksSummaryListOK{}
}

/*
V1PacksSummaryListOK handles this case with default header values.

An array of pack summary items
*/
type V1PacksSummaryListOK struct {
	Payload *models.V1PackSummaries
}

func (o *V1PacksSummaryListOK) Error() string {
	return fmt.Sprintf("[GET /v1/packs][%d] v1PacksSummaryListOK  %+v", 200, o.Payload)
}

func (o *V1PacksSummaryListOK) GetPayload() *models.V1PackSummaries {
	return o.Payload
}

func (o *V1PacksSummaryListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PackSummaries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}