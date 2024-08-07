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

// V1AuthOrgSwitchReader is a Reader for the V1AuthOrgSwitch structure.
type V1AuthOrgSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AuthOrgSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AuthOrgSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AuthOrgSwitchOK creates a V1AuthOrgSwitchOK with default headers values
func NewV1AuthOrgSwitchOK() *V1AuthOrgSwitchOK {
	return &V1AuthOrgSwitchOK{}
}

/*
V1AuthOrgSwitchOK handles this case with default header values.

OK
*/
type V1AuthOrgSwitchOK struct {
	Payload *models.V1UserToken
}

func (o *V1AuthOrgSwitchOK) Error() string {
	return fmt.Sprintf("[POST /v1/auth/org/{orgName}/switch][%d] v1AuthOrgSwitchOK  %+v", 200, o.Payload)
}

func (o *V1AuthOrgSwitchOK) GetPayload() *models.V1UserToken {
	return o.Payload
}

func (o *V1AuthOrgSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
