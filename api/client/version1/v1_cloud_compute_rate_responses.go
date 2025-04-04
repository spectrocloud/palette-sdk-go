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

// V1CloudComputeRateReader is a Reader for the V1CloudComputeRate structure.
type V1CloudComputeRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudComputeRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudComputeRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudComputeRateOK creates a V1CloudComputeRateOK with default headers values
func NewV1CloudComputeRateOK() *V1CloudComputeRateOK {
	return &V1CloudComputeRateOK{}
}

/*
V1CloudComputeRateOK handles this case with default header values.

(empty)
*/
type V1CloudComputeRateOK struct {
	Payload *models.V1CloudCost
}

func (o *V1CloudComputeRateOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/{cloud}/compute/{type}/rate][%d] v1CloudComputeRateOK  %+v", 200, o.Payload)
}

func (o *V1CloudComputeRateOK) GetPayload() *models.V1CloudCost {
	return o.Payload
}

func (o *V1CloudComputeRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CloudCost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
