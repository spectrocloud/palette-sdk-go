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

// V1AwsCloudCostReader is a Reader for the V1AwsCloudCost structure.
type V1AwsCloudCostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsCloudCostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsCloudCostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsCloudCostOK creates a V1AwsCloudCostOK with default headers values
func NewV1AwsCloudCostOK() *V1AwsCloudCostOK {
	return &V1AwsCloudCostOK{}
}

/*
V1AwsCloudCostOK handles this case with default header values.

(empty)
*/
type V1AwsCloudCostOK struct {
	Payload *models.V1AwsCloudCostSummary
}

func (o *V1AwsCloudCostOK) Error() string {
	return fmt.Sprintf("[POST /v1/clouds/aws/cost][%d] v1AwsCloudCostOK  %+v", 200, o.Payload)
}

func (o *V1AwsCloudCostOK) GetPayload() *models.V1AwsCloudCostSummary {
	return o.Payload
}

func (o *V1AwsCloudCostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsCloudCostSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
