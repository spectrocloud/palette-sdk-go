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

// V1MaasSubnetsGetReader is a Reader for the V1MaasSubnetsGet structure.
type V1MaasSubnetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MaasSubnetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1MaasSubnetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1MaasSubnetsGetOK creates a V1MaasSubnetsGetOK with default headers values
func NewV1MaasSubnetsGetOK() *V1MaasSubnetsGetOK {
	return &V1MaasSubnetsGetOK{}
}

/*V1MaasSubnetsGetOK handles this case with default header values.

(empty)
*/
type V1MaasSubnetsGetOK struct {
	Payload *models.V1MaasSubnets
}

func (o *V1MaasSubnetsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/maas/subnets][%d] v1MaasSubnetsGetOK  %+v", 200, o.Payload)
}

func (o *V1MaasSubnetsGetOK) GetPayload() *models.V1MaasSubnets {
	return o.Payload
}

func (o *V1MaasSubnetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MaasSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
