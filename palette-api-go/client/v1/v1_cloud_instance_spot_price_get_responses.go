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

// V1CloudInstanceSpotPriceGetReader is a Reader for the V1CloudInstanceSpotPriceGet structure.
type V1CloudInstanceSpotPriceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudInstanceSpotPriceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudInstanceSpotPriceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudInstanceSpotPriceGetOK creates a V1CloudInstanceSpotPriceGetOK with default headers values
func NewV1CloudInstanceSpotPriceGetOK() *V1CloudInstanceSpotPriceGetOK {
	return &V1CloudInstanceSpotPriceGetOK{}
}

/*V1CloudInstanceSpotPriceGetOK handles this case with default header values.

(empty)
*/
type V1CloudInstanceSpotPriceGetOK struct {
	Payload *models.V1CloudSpotPrice
}

func (o *V1CloudInstanceSpotPriceGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/{cloudType}/instance/spotprice][%d] v1CloudInstanceSpotPriceGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudInstanceSpotPriceGetOK) GetPayload() *models.V1CloudSpotPrice {
	return o.Payload
}

func (o *V1CloudInstanceSpotPriceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CloudSpotPrice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
