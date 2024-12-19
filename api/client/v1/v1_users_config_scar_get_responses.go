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

// V1UsersConfigScarGetReader is a Reader for the V1UsersConfigScarGet structure.
type V1UsersConfigScarGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersConfigScarGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersConfigScarGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersConfigScarGetOK creates a V1UsersConfigScarGetOK with default headers values
func NewV1UsersConfigScarGetOK() *V1UsersConfigScarGetOK {
	return &V1UsersConfigScarGetOK{}
}

/*V1UsersConfigScarGetOK handles this case with default header values.

(empty)
*/
type V1UsersConfigScarGetOK struct {
	Payload *models.V1SystemScarSpec
}

func (o *V1UsersConfigScarGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/config/scar][%d] v1UsersConfigScarGetOK  %+v", 200, o.Payload)
}

func (o *V1UsersConfigScarGetOK) GetPayload() *models.V1SystemScarSpec {
	return o.Payload
}

func (o *V1UsersConfigScarGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SystemScarSpec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
