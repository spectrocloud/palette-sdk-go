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

// V1AppProfilesUIDGetReader is a Reader for the V1AppProfilesUIDGet structure.
type V1AppProfilesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppProfilesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDGetOK creates a V1AppProfilesUIDGetOK with default headers values
func NewV1AppProfilesUIDGetOK() *V1AppProfilesUIDGetOK {
	return &V1AppProfilesUIDGetOK{}
}

/*
V1AppProfilesUIDGetOK handles this case with default header values.

OK
*/
type V1AppProfilesUIDGetOK struct {
	Payload *models.V1AppProfile
}

func (o *V1AppProfilesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appProfiles/{uid}][%d] v1AppProfilesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1AppProfilesUIDGetOK) GetPayload() *models.V1AppProfile {
	return o.Payload
}

func (o *V1AppProfilesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
