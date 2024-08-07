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

// V1UsersAssetsLocationGcpGetReader is a Reader for the V1UsersAssetsLocationGcpGet structure.
type V1UsersAssetsLocationGcpGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersAssetsLocationGcpGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersAssetsLocationGcpGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersAssetsLocationGcpGetOK creates a V1UsersAssetsLocationGcpGetOK with default headers values
func NewV1UsersAssetsLocationGcpGetOK() *V1UsersAssetsLocationGcpGetOK {
	return &V1UsersAssetsLocationGcpGetOK{}
}

/*V1UsersAssetsLocationGcpGetOK handles this case with default header values.

(empty)
*/
type V1UsersAssetsLocationGcpGetOK struct {
	Payload *models.V1UserAssetsLocationGcp
}

func (o *V1UsersAssetsLocationGcpGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/assets/locations/gcp/{uid}][%d] v1UsersAssetsLocationGcpGetOK  %+v", 200, o.Payload)
}

func (o *V1UsersAssetsLocationGcpGetOK) GetPayload() *models.V1UserAssetsLocationGcp {
	return o.Payload
}

func (o *V1UsersAssetsLocationGcpGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserAssetsLocationGcp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
