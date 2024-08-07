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

// V1CloudConfigsAksGetReader is a Reader for the V1CloudConfigsAksGet structure.
type V1CloudConfigsAksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudConfigsAksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudConfigsAksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudConfigsAksGetOK creates a V1CloudConfigsAksGetOK with default headers values
func NewV1CloudConfigsAksGetOK() *V1CloudConfigsAksGetOK {
	return &V1CloudConfigsAksGetOK{}
}

/*V1CloudConfigsAksGetOK handles this case with default header values.

OK
*/
type V1CloudConfigsAksGetOK struct {
	Payload *models.V1AzureCloudConfig
}

func (o *V1CloudConfigsAksGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudconfigs/aks/{configUid}][%d] v1CloudConfigsAksGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudConfigsAksGetOK) GetPayload() *models.V1AzureCloudConfig {
	return o.Payload
}

func (o *V1CloudConfigsAksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AzureCloudConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
