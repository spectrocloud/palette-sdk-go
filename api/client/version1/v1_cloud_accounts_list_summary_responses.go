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

// V1CloudAccountsListSummaryReader is a Reader for the V1CloudAccountsListSummary structure.
type V1CloudAccountsListSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsListSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsListSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsListSummaryOK creates a V1CloudAccountsListSummaryOK with default headers values
func NewV1CloudAccountsListSummaryOK() *V1CloudAccountsListSummaryOK {
	return &V1CloudAccountsListSummaryOK{}
}

/*
V1CloudAccountsListSummaryOK handles this case with default header values.

An array of cloud account summary items
*/
type V1CloudAccountsListSummaryOK struct {
	Payload *models.V1CloudAccountsSummary
}

func (o *V1CloudAccountsListSummaryOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/summary][%d] v1CloudAccountsListSummaryOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsListSummaryOK) GetPayload() *models.V1CloudAccountsSummary {
	return o.Payload
}

func (o *V1CloudAccountsListSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CloudAccountsSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
