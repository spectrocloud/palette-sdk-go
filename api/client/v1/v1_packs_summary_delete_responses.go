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

// V1PacksSummaryDeleteReader is a Reader for the V1PacksSummaryDelete structure.
type V1PacksSummaryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1PacksSummaryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PacksSummaryDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PacksSummaryDeleteOK creates a V1PacksSummaryDeleteOK with default headers values
func NewV1PacksSummaryDeleteOK() *V1PacksSummaryDeleteOK {
	return &V1PacksSummaryDeleteOK{}
}

/*
V1PacksSummaryDeleteOK handles this case with default header values.

(empty)
*/
type V1PacksSummaryDeleteOK struct {
	Payload *models.V1DeleteMeta
}

func (o *V1PacksSummaryDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/packs][%d] v1PacksSummaryDeleteOK  %+v", 200, o.Payload)
}

func (o *V1PacksSummaryDeleteOK) GetPayload() *models.V1DeleteMeta {
	return o.Payload
}

func (o *V1PacksSummaryDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1DeleteMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}