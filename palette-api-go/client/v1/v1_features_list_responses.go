// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1FeaturesListReader is a Reader for the V1FeaturesList structure.
type V1FeaturesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1FeaturesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1FeaturesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1FeaturesListOK creates a V1FeaturesListOK with default headers values
func NewV1FeaturesListOK() *V1FeaturesListOK {
	return &V1FeaturesListOK{}
}

/*V1FeaturesListOK handles this case with default header values.

OK
*/
type V1FeaturesListOK struct {
	Payload *models.V1Features
}

func (o *V1FeaturesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/features][%d] v1FeaturesListOK  %+v", 200, o.Payload)
}

func (o *V1FeaturesListOK) GetPayload() *models.V1Features {
	return o.Payload
}

func (o *V1FeaturesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Features)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
