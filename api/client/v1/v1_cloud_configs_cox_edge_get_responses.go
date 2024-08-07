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

// V1CloudConfigsCoxEdgeGetReader is a Reader for the V1CloudConfigsCoxEdgeGet structure.
type V1CloudConfigsCoxEdgeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsCoxEdgeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsCoxEdgeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsCoxEdgeGetOK creates a V1CloudConfigsCoxEdgeGetOK with default headers values
func NewV1CloudConfigsCoxEdgeGetOK() *V1CloudConfigsCoxEdgeGetOK {
	return &V1CloudConfigsCoxEdgeGetOK{}
}

/*V1CloudConfigsCoxEdgeGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsCoxEdgeGetOK struct {
	Payload *models.V1CoxEdgeCloudConfig
}

func (o *V1CloudConfigsCoxEdgeGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/coxedge/{configUid}][%d] v1CloudConfigsCoxEdgeGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsCoxEdgeGetOK) GetPayload() *models.V1CoxEdgeCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsCoxEdgeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CoxEdgeCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
