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

// V1SpectroClustersSummaryUIDReader is a Reader for the V1SpectroClustersSummaryUID structure.
type V1SpectroClustersSummaryUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersSummaryUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersSummaryUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersSummaryUIDOK creates a V1SpectroClustersSummaryUIDOK with default headers values
func NewV1SpectroClustersSummaryUIDOK() *V1SpectroClustersSummaryUIDOK {
	return &V1SpectroClustersSummaryUIDOK{}
}

/*V1SpectroClustersSummaryUIDOK handles this case with default header values.

An spectro cluster summary
*/
type V1SpectroClustersSummaryUIDOK struct {
	Payload *models.V1SpectroClusterUIDSummary
}

func (o *V1SpectroClustersSummaryUIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/spectroclusters/{uid}][%d] v1SpectroClustersSummaryUidOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersSummaryUIDOK) GetPayload() *models.V1SpectroClusterUIDSummary {
	return o.Payload
}

func (o *V1SpectroClustersSummaryUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterUIDSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
