// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1AuthOrgsReader is a Reader for the V1AuthOrgs structure.
type V1AuthOrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AuthOrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AuthOrgsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AuthOrgsOK creates a V1AuthOrgsOK with default headers values
func NewV1AuthOrgsOK() *V1AuthOrgsOK {
	return &V1AuthOrgsOK{}
}

/*
V1AuthOrgsOK handles this case with default header values.

OK
*/
type V1AuthOrgsOK struct {
	Payload *models.V1Organizations
}

func (o *V1AuthOrgsOK) Error() string {
	return fmt.Sprintf("[GET /v1/auth/orgs][%d] v1AuthOrgsOK  %+v", 200, o.Payload)
}

func (o *V1AuthOrgsOK) GetPayload() *models.V1Organizations {
	return o.Payload
}

func (o *V1AuthOrgsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Organizations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
