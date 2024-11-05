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

// V1CloudStorageRateReader is a Reader for the V1CloudStorageRate structure.
type V1CloudStorageRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudStorageRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudStorageRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudStorageRateOK creates a V1CloudStorageRateOK with default headers values
func NewV1CloudStorageRateOK() *V1CloudStorageRateOK {
	return &V1CloudStorageRateOK{}
}

/*V1CloudStorageRateOK handles this case with default header values.

(empty)
*/
type V1CloudStorageRateOK struct {
	Payload *models.V1CloudCost
}

func (o *V1CloudStorageRateOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/{cloud}/storage/{type}/rate][%d] v1CloudStorageRateOK  %+v", 200, o.Payload)
}

func (o *V1CloudStorageRateOK) GetPayload() *models.V1CloudCost {
	return o.Payload
}

func (o *V1CloudStorageRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CloudCost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
