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

// V1MaasPoolsGetReader is a Reader for the V1MaasPoolsGet structure.
type V1MaasPoolsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MaasPoolsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1MaasPoolsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1MaasPoolsGetOK creates a V1MaasPoolsGetOK with default headers values
func NewV1MaasPoolsGetOK() *V1MaasPoolsGetOK {
	return &V1MaasPoolsGetOK{}
}

/*
V1MaasPoolsGetOK handles this case with default header values.

(empty)
*/
type V1MaasPoolsGetOK struct {
	Payload *models.V1MaasPools
}

func (o *V1MaasPoolsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/maas/resourcePools][%d] v1MaasPoolsGetOK  %+v", 200, o.Payload)
}

func (o *V1MaasPoolsGetOK) GetPayload() *models.V1MaasPools {
	return o.Payload
}

func (o *V1MaasPoolsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MaasPools)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
