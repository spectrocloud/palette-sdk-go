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

// V1MetricsUIDListReader is a Reader for the V1MetricsUIDList structure.
type V1MetricsUIDListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MetricsUIDListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1MetricsUIDListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1MetricsUIDListOK creates a V1MetricsUIDListOK with default headers values
func NewV1MetricsUIDListOK() *V1MetricsUIDListOK {
	return &V1MetricsUIDListOK{}
}

/*V1MetricsUIDListOK handles this case with default header values.

An array of metric items
*/
type V1MetricsUIDListOK struct {
	Payload *models.V1MetricTimeSeries
}

func (o *V1MetricsUIDListOK) Error() string {
	return fmt.Sprintf("[GET /v1/metrics/{resourceKind}/{resourceUid}/values][%d] v1MetricsUidListOK  %+v", 200, o.Payload)
}

func (o *V1MetricsUIDListOK) GetPayload() *models.V1MetricTimeSeries {
	return o.Payload
}

func (o *V1MetricsUIDListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MetricTimeSeries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
